package service

import (
	"errors"
	"fmt"
	"github.com/axiangcoding/ax-web/cache"
	"github.com/axiangcoding/ax-web/logging"
	"github.com/axiangcoding/ax-web/service/bot"
	"github.com/axiangcoding/ax-web/service/cqhttp"
	"github.com/axiangcoding/ax-web/settings"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/exp/slices"
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
	key := generateCacheKey(cqhttp.PostTypeMetaEvent, cqhttp.EventTypeHeartBeat, selfId)
	result, err := cache.GetClient().Get(c, key).Result()
	if err != nil {
		return message, err
	}
	if err := message.UnmarshalBinary([]byte(result)); err != nil {
		return message, err
	}
	return message, nil
}

func generateCacheKey(postType string, eventType string, selfId int64) string {
	return fmt.Sprintf("CQHTTP:%s;%s;%d", postType, eventType, selfId)
}

func handleCqHttpMetaEventHeartBeat(c *gin.Context, event *cqhttp.MetaTypeHeartBeatEvent) {
	key := generateCacheKey(event.PostType, event.MetaEventType, event.SelfId)
	if err := cache.GetClient().Set(c, key, event, time.Minute).Err(); err != nil {
		logging.Error(err)
	}
}

func handleCqHttpMessageEventGroup(c *gin.Context, event *cqhttp.MessageGroupEvent) {
	messageType := event.MessageType
	msg := event.Message
	// 处理群组at消息，只有at我的群组消息才处理，其他的一律抛弃
	settingSelfQQ := settings.Config.Service.CqHttp.SelfQQ
	if messageType != "group" || !cqhttp.MustContainsCqCode(msg) || settingSelfQQ != cqhttp.MustGetCqCodeAtQQ(msg) {
		return
	}
	action := bot.ParseMessageCommand(msg)
	var retMsgForm cqhttp.SendGroupMsgForm
	retMsgForm.GroupId = event.GroupId
	retMsgPrefix := fmt.Sprintf("[CQ:at,qq=%d] ", event.Sender.UserId)
	if action == nil {
		retMsgForm.Message = retMsgPrefix + "我不道你在说什么，笨笨，呜呜"
	} else {
		switch action.Key {
		case bot.ActionQuery:
			async, user, err := QueryWTGamerProfile(action.Value)
			if err != nil {
				logging.Warnf("query WT gamer profile error. %s", err)
				retMsgForm.Message = retMsgPrefix + "啊哦，目前无法查询，请稍后重试"
			}
			if async {
				retMsgForm.Message = retMsgPrefix + "正在查询中，请稍后..."
			} else {
				retMsgForm.Message = retMsgPrefix + user.ToFriendlyString()
			}
			break
		case bot.ActionRefresh:
			if err := RefreshWTGamerProfile(action.Value); err != nil {
				logging.Warn("refresh WT gamer profile error. ", err)
				retMsgForm.Message = retMsgPrefix + "啊哦，目前无法刷新，请稍后重试"
			}
			retMsgForm.Message = retMsgPrefix + "正在刷新已有数据，请稍后..."
			break
		case bot.ActionReport:
			retMsgForm.Message = retMsgPrefix + "举办他是吧，记住你了，晚上别锁房门"
			break
		case bot.ActionDrawCard:
			id := event.Sender.UserId
			number := DrawNumber(id, time.Now().In(time.FixedZone("CST", 8*3600)))
			retMsgForm.Message = retMsgPrefix + fmt.Sprintf("你今天的气运值是%d", number)
			break
		default:
			retMsgForm.Message = retMsgPrefix + "我不道你想干啥，笨笨，呜呜"
			break
		}
	}

	result, err := cqhttp.SendGroupMsg(retMsgForm)
	if err != nil {
		logging.Warnf("handle group message error %s", err)
		return
	}
	if result != nil && result.Status == "failed" {
		logging.Warnf("send message failed. response json: %#v", result)
	}
}