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
			switch data["meta_event_type"] {
			case cqhttp.EventTypeHeartBeat:
				var event cqhttp.MetaTypeHeartBeatEvent
				if err := mapstructure.Decode(data, &event); err != nil {
					return err
				}
				handleCqHttpMetaEventHeartBeat(c, &event)
				break
			default:
				logging.Warnf("meta_event_type %s not supported yet.", data["meta_event_type"])
			}
			break
		case cqhttp.PostTypeMessage:
			switch data["message_type"] {
			case cqhttp.MessageTypeGroup:
				var event cqhttp.MessageGroupEvent
				if err := mapstructure.Decode(data, &event); err != nil {
					return err
				}
				handleCqHttpMessageEventGroup(c, &event)
				break
			default:
				logging.Warnf("message_type %s not supported yet.", data["message_type"])
			}
		case cqhttp.PostTypeRequest:
			break
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

func handleCqHttpMessageEventGroup(c *gin.Context, event *cqhttp.MessageGroupEvent) {
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
	if *gc.Banned {
		retMsgForm.Message = bot.RespGroupGetBanned
	} else {
		retMsgForm.MessagePrefix = fmt.Sprintf("[CQ:at,qq=%d] ", event.Sender.UserId)
		action := bot.ParseMessageCommand(msg)
		if action == nil {
			retMsgForm.Message = bot.RespCommon
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
				retMsgForm.Message = bot.RespReport
				break
			case bot.ActionDrawCard:
				DoActionDrawCard(&retMsgForm, value, event.Sender.UserId)
				break
			case bot.ActionLuck:
				DoActionLuck(&retMsgForm, value, event.Sender.UserId)
				break
			case bot.ActionVersion:
				retMsgForm.Message = fmt.Sprintf(bot.RespVersion, settings.Config.Version)
				break
			case bot.ActionGetHelp:
				retMsgForm.Message = bot.RespGetHelp
				break
			default:
				retMsgForm.Message = bot.RespGetHelp
				break
			}
		}
	}
	MustSendGroupMsg(retMsgForm)
}

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
