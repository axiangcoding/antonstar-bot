package service

import (
	"axiangcoding/antonstar/api-system/logging"
	"axiangcoding/antonstar/api-system/settings"
	"encoding/base64"
	"errors"
	"github.com/go-resty/resty/v2"
)

var superBedUploadUrl = "https://api.superbed.cn/upload"

func UploadToSuperBed(bytes []byte) (map[string]interface{}, error) {
	params := map[string]string{
		"token":      settings.Config.App.Upload.SuperBed.Token,
		"categories": settings.Config.App.Upload.SuperBed.Categories,
		"b64_data":   base64.StdEncoding.EncodeToString(bytes),
		"v":          "2",
	}
	client := resty.New()
	resBody := map[string]interface{}{}
	_, err := client.R().SetBody(params).SetResult(&resBody).Post(superBedUploadUrl)
	if err != nil {
		logging.Warn("Post file to Superbed failed", err)
		return nil, err
	}
	if resBody["err"] != 0.0 {
		logging.Warnf("Post file to Superbed return err %d, %s", resBody["err"], resBody["msg"])
		return nil, errors.New("post file to Superbed failed")
	}
	return map[string]interface{}{
		"url": resBody["url"],
	}, nil
}
