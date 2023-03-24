package service

import (
	"errors"
	"fmt"
	"github.com/axiangcoding/antonstar-bot/internal/cache"
	"github.com/axiangcoding/antonstar-bot/internal/data/table"
	"github.com/axiangcoding/antonstar-bot/pkg/bot"
	"github.com/axiangcoding/antonstar-bot/pkg/cqhttp"
	"github.com/axiangcoding/antonstar-bot/pkg/logging"
	"github.com/axiangcoding/antonstar-bot/setting"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
	"time"
)

var allowPostType = []string{cqhttp.PostTypeMessage, cqhttp.PostTypeRequest, cqhttp.PostTypeNotice, cqhttp.PostTypeMetaEvent}

func HandleCqHttpEvent(c *gin.Context, data map[string]any) error {
	postType := data["post_type"]
	if slices.Contains(allowPostType, fmt.Sprintf("%v", postType)) {
		switch postType {
		case cqhttp.PostTypeMetaEvent:
			var event cqhttp.MetaTypeHeartBeatEvent
			if err := mapstructure.Decode(data, &event); err != nil {
				return err
			}
			switch data["meta_event_type"] {
			case cqhttp.EventTypeHeartBeat:
				handleCqHttpMetaEventHeartBeat(c, &event)
				break
			default:
				logging.L().Warn("meta_event_type not supported yet",
					logging.Any("meta_event_type", data["meta_event_type"]))
			}
			break
		case cqhttp.PostTypeMessage:
			var event cqhttp.CommonEvent
			if err := mapstructure.Decode(data, &event); err != nil {
				return err
			}
			switch data["message_type"] {
			case cqhttp.MessageTypeGroup:
				handleCqHttpMessageEventGroup(&event)
				break
			default:
				logging.L().Warn("message_type not supported yet",
					logging.Any("message_type", data["message_type"]))
			}
		case cqhttp.PostTypeRequest:
			var event cqhttp.CommonEvent
			if err := mapstructure.Decode(data, &event); err != nil {
				return err
			}
			switch data["request_type"] {
			case cqhttp.RequestTypeGroup:
				if data["sub_type"] == cqhttp.SubTypeInvite {
					handleAddGroup(&event)
				}
				break
			case cqhttp.RequestTypeFriend:
				if data["sub_type"] == cqhttp.SubTypeInvite {
					handleAddFriend(c, &event)
				}
				break
			}
		case cqhttp.PostTypeNotice:
			break
		}
	} else {
		return errors.New("no such event_type")
	}
	return nil
}

func GetCqHttpStatus(c *gin.Context, selfId int64) (cqhttp.MetaTypeHeartBeatEvent, error) {
	var message cqhttp.MetaTypeHeartBeatEvent
	key := cache.GenerateCQHTTPCacheKey(cqhttp.PostTypeMetaEvent, cqhttp.EventTypeHeartBeat, selfId)
	result, err := cache.Client().Get(c, key).Result()
	if err != nil {
		return message, err
	}
	if err := message.UnmarshalBinary([]byte(result)); err != nil {
		return message, err
	}
	return message, nil
}

func handleCqHttpMetaEventHeartBeat(c *gin.Context, event *cqhttp.MetaTypeHeartBeatEvent) {
	key := cache.GenerateCQHTTPCacheKey(event.PostType, event.MetaEventType, event.SelfId)
	if err := cache.Client().Set(c, key, event, time.Minute).Err(); err != nil {
		logging.L().Error("set cache error", logging.Error(err))
	}
}

func handleCqHttpMessageEventGroup(event *cqhttp.CommonEvent) {
	messageType := event.MessageType
	msg := event.Message
	if messageType != "group" || !cqhttp.MustContainsTrigger(msg) {
		return
	}
	action := bot.ParseMessageCommand(msg)
	stopAllResponse := IsStopAllResponse()
	if stopAllResponse && (action == nil || action.Key != bot.ActionManager) {
		return
	}
	groupId := event.GroupId
	userId := event.UserId
	gc, err := FindGroupConfig(groupId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			defaultGC := table.DefaultGroupConfig(groupId)
			gc = &defaultGC
			MustSaveGroupConfig(gc)
		} else {
			logging.L().Warn("find group config failed", logging.Error(err))
		}
	}
	uc, err := FindUserConfig(userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			defaultUC := table.DefaultUserConfig(userId)
			uc = &defaultUC
			MustSaveUserConfig(uc)
		} else {
			logging.L().Warn("find user config failed", logging.Error(err))
		}
	}
	var retMsgForm cqhttp.SendGroupMsgForm
	retMsgForm.GroupId = groupId
	retMsgForm.MessageTemplate = gc.MessageTemplate
	retMsgForm.UserId = userId
	// 检查qq群请求限制
	if limit, usage, total := CheckGroupTodayUsageLimit(groupId); limit {
		if !ExistGroupUsageLimitFlag(groupId) {
			retMsgForm.Message = fmt.Sprintf(bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.TodayGroupUsageLimit, usage, total)
			cqhttp.MustSendGroupMsg(retMsgForm)
			MustPutGroupUsageLimitFlag(groupId)
			return
		} else {
			return
		}
	}
	// 检查qq请求限制
	if limit, usage, total := CheckUserTodayUsageLimit(userId); limit {
		if !ExistUserUsageLimitFlag(userId) {
			retMsgForm.MessagePrefix = fmt.Sprintf("[CQ:at,qq=%d] ", event.Sender.UserId)
			retMsgForm.Message = fmt.Sprintf(bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.TodayUserUsageLimit, usage, total)
			cqhttp.MustSendGroupMsg(retMsgForm)
			MustPutUserUsageLimitFlag(userId)
			return
		} else {
			return
		}
	}
	MustAddGroupConfigTodayUsageCount(groupId, 1)
	MustAddGroupConfigTotalUsageCount(groupId, 1)
	MustAddUserConfigTodayUsageCount(userId, 1)
	MustAddUserConfigTotalUsageCount(userId, 1)

	if *gc.Banned {
		retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.GroupGetBanned
		cqhttp.MustSendGroupMsg(retMsgForm)
		return
	}
	if *uc.Banned {
		retMsgForm.MessagePrefix = fmt.Sprintf("[CQ:at,qq=%d] ", event.Sender.UserId)
		retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.UserGetBanned
		cqhttp.MustSendGroupMsg(retMsgForm)
		return
	}
	retMsgForm.MessagePrefix = fmt.Sprintf("[CQ:at,qq=%d] ", event.Sender.UserId)
	if action == nil {
		retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.Common
	} else {
		value := action.Value
		switch action.Key {
		case bot.ActionQuery:
			DoActionQuery(&retMsgForm, value, false)
		case bot.ActionFullQuery:
			DoActionQuery(&retMsgForm, value, true)
		case bot.ActionRefresh:
			DoActionRefresh(&retMsgForm, value)
		case bot.ActionReport:
			retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.Report
		case bot.ActionDrawCard:
			DoActionDrawCard(&retMsgForm, value, event.Sender.UserId)
		case bot.ActionLuck:
			DoActionLuck(&retMsgForm, value, event.Sender.UserId)
		case bot.ActionVersion:
			retMsgForm.Message = fmt.Sprintf(bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.Version, setting.C().App.Version)
		case bot.ActionGetHelp:
			retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.GetHelp
		case bot.ActionGroupStatus:
			DoActionGroupStatus(&retMsgForm)
		case bot.ActionGroupManager:
			DoActionGroupManager(&retMsgForm, uc, value)
		case bot.ActionData:
			DoActionData(&retMsgForm, value)
		case bot.ActionBinding:
			DoActionBinding(&retMsgForm, value)
		case bot.ActionUnbinding:
			DoActionUnbinding(&retMsgForm)
		case bot.ActionManager:
			DoActionManager(&retMsgForm, uc, value)
		default:
			retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.GetHelp
		}
	}

	cqhttp.MustSendGroupMsg(retMsgForm)
}

func handleAddGroup(event *cqhttp.CommonEvent) {
	if event.SubType == cqhttp.SubTypeInvite {
		groupId := event.GroupId
		userId := event.UserId
		groupConfig := MustFindGroupConfig(groupId)
		userConfig := MustFindUserConfig(userId)
		if (groupConfig == nil || !(*groupConfig.Banned)) && (userConfig == nil || !(*userConfig.Banned)) && (userConfig.SuperAdmin != nil && *userConfig.SuperAdmin) {
			cqhttp.MustAcceptInviteToGroup(event.Flag, event.SubType, true, "")
		} else {
			logging.L().Warn("application for joining the group was rejected",
				logging.Any("userId", userId),
				logging.Any("groupId", groupId))
		}
	}
}

func handleAddFriend(c *gin.Context, event *cqhttp.CommonEvent) {

}
