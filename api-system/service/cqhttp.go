package service

import (
	"errors"
	"fmt"
	"github.com/axiangcoding/ax-web/cache"
	"github.com/axiangcoding/ax-web/logging"
	"github.com/axiangcoding/ax-web/service/cqhttp"
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
				var event cqhttp.MetaTypeHeartBeatEventMessage
				if err := mapstructure.Decode(data, &event); err != nil {
					return err
				}
				handleCqHttpMetaEventHeartBeat(c, &event)
				break
			default:
				fmt.Println(data["meta_event_type"])
			}

			break
		case cqhttp.PostTypeMessage:
			break
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

func handleCqHttpMetaEventHeartBeat(c *gin.Context, event *cqhttp.MetaTypeHeartBeatEventMessage) {
	key := generateCacheKey(event.PostType, event.MetaEventType, event.SelfId)
	if err := cache.GetClient().Set(c, key, event, time.Minute).Err(); err != nil {
		logging.Error(err)
	}
}

func GetCqHttpStatus(c *gin.Context, selfId int64) (cqhttp.MetaTypeHeartBeatEventMessage, error) {
	var message cqhttp.MetaTypeHeartBeatEventMessage
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
	return fmt.Sprintf("CQHTTP#%s;%s;%d", postType, eventType, selfId)
}
