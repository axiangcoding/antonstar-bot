package service

import (
	"errors"
	"fmt"
	"github.com/axiangcoding/ax-web/cache"
	"github.com/axiangcoding/ax-web/data/table"
	"github.com/axiangcoding/ax-web/logging"
	"github.com/axiangcoding/ax-web/service/bot"
	"github.com/axiangcoding/ax-web/service/cqhttp"
	"github.com/axiangcoding/ax-web/settings"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
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
				logging.Warnf("meta_event_type %s not supported yet.", data["meta_event_type"])
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
				logging.Warnf("message_type %s not supported yet.", data["message_type"])
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
	result, err := cache.GetClient().Get(c, key).Result()
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
	if err := cache.GetClient().Set(c, key, event, time.Minute).Err(); err != nil {
		logging.Error(err)
	}
}

func handleCqHttpMessageEventGroup(event *cqhttp.CommonEvent) {
	messageType := event.MessageType
	msg := event.Message
	if messageType != "group" || !cqhttp.MustContainsTrigger(msg) {
		return
	}
	groupId := event.GroupId
	gc, err := FindGroupConfig(groupId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			defaultGC := table.DefaultGroupConfig(groupId)
			gc = &defaultGC
			if err := SaveGroupConfig(defaultGC); err != nil {
				logging.Warn(err)
			}
		} else {
			logging.Warn(err)
		}
	}
	var retMsgForm cqhttp.SendGroupMsgForm
	retMsgForm.GroupId = groupId
	retMsgForm.MessageTemplate = gc.MessageTemplate
	if *gc.Banned {
		retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.GroupGetBanned
	} else {
		retMsgForm.MessagePrefix = fmt.Sprintf("[CQ:at,qq=%d] ", event.Sender.UserId)
		action := bot.ParseMessageCommand(msg)
		if action == nil {
			retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.Common
		} else {
			value := action.Value
			switch action.Key {
			case bot.ActionQuery:
				DoActionQuery(&retMsgForm, value, false)
				break
			case bot.ActionFullQuery:
				DoActionQuery(&retMsgForm, value, true)
				break
			case bot.ActionRefresh:
				DoActionRefresh(&retMsgForm, value)
				break
			case bot.ActionReport:
				retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.Report
				break
			case bot.ActionDrawCard:
				DoActionDrawCard(&retMsgForm, value, event.Sender.UserId)
				break
			case bot.ActionLuck:
				DoActionLuck(&retMsgForm, value, event.Sender.UserId)
				break
			case bot.ActionVersion:
				retMsgForm.Message = fmt.Sprintf(bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.Version, settings.Config.Version)
				break
			case bot.ActionGetHelp:
				retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.GetHelp
				break
			case bot.ActionGroupStatus:
				DoActionGroupStatus(&retMsgForm)
				break
			default:
				retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.GetHelp
				break
			}
		}
	}
	MustSendGroupMsg(retMsgForm)
}

func handleAddGroup(event *cqhttp.CommonEvent) {
	if event.SubType == cqhttp.SubTypeInvite {
		config := MustFindGroupConfig(event.GroupId)
		if config == nil || !(*config.Banned) {
			MustAcceptInviteToGroup(event.Flag, event.SubType, true, "")
		}
	}
}

func handleAddFriend(c *gin.Context, event *cqhttp.CommonEvent) {

}

// MustSendGroupMsg
// https://docs.go-cqhttp.org/api/#%E5%A4%84%E7%90%86%E5%8A%A0%E7%BE%A4%E8%AF%B7%E6%B1%82-%E9%82%80%E8%AF%B7
func MustSendGroupMsg(form cqhttp.SendGroupMsgForm) {
	url := settings.Config.Service.CqHttp.Url + "/send_group_msg"
	client := resty.New().SetTimeout(time.Second * 20)
	var commonResp cqhttp.CommonResponse
	resp, err := client.R().SetHeader("Content-Type", "application/json").
		SetBody(map[string]any{
			"message":  form.MessagePrefix + form.Message,
			"group_id": form.GroupId,
		}).SetResult(&commonResp).Post(url)
	if err != nil {
		logging.Warnf("send group message error. %s", err)
	}
	if resp.IsError() {
		logging.Warnf("post %s error. code=%d, message=%s", url, resp.StatusCode(), resp.String())
	}
	if commonResp.Status == "failed" {
		logging.Warnf("send message failed. response json: %#v", commonResp)
	}
}

// MustAcceptInviteToGroup
// https://docs.go-cqhttp.org/api/#%E5%A4%84%E7%90%86%E5%8A%A0%E7%BE%A4%E8%AF%B7%E6%B1%82-%E9%82%80%E8%AF%B7
func MustAcceptInviteToGroup(flag string, subType string, approve bool, reason string) {
	url := settings.Config.Service.CqHttp.Url + "/set_group_add_request"
	client := resty.New().SetTimeout(time.Second * 20)
	resp, err := client.R().SetHeader("Content-Type", "application/json").
		SetBody(map[string]any{
			"flag":     flag,
			"sub_type": subType,
			"approve":  approve,
			"reason":   reason,
		}).Post(url)
	if err != nil {
		logging.Warnf("send group add request error. %s", err)
	}
	if resp.IsError() {
		logging.Warnf("post %s error. code=%d, message=%s", url, resp.StatusCode(), resp.String())
	}
}
