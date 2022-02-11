package service

import (
	"github.com/dchest/captcha"
)

func GenerateCaptcha() map[string]string {
	length := captcha.DefaultLen
	captchaId := captcha.NewLen(length)
	return map[string]string{
		"file_name": captchaId + ".png",
	}
}
