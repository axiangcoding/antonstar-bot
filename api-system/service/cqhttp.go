package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/axiangcoding/ax-web/logging"
	"github.com/axiangcoding/ax-web/service/cqhttp"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

var allowPostType = []string{cqhttp.PostTypeMessage, cqhttp.PostTypeRequest, cqhttp.PostTypeNotice, cqhttp.PostTypeMetaEvent}

func HandleCqHttpEvent(c *gin.Context, data map[string]any) error {
	postType := data["post_type"]
	if slices.Contains(allowPostType, fmt.Sprintf("%v", postType)) {
		marshal, err := json.MarshalIndent(data, "", "\t")
		if err != nil {
			return err
		}
		logging.Info(string(marshal))
	} else {
		return errors.New("no such event_type")
	}
	return nil
}
