package validation

import (
	"axiangcoding/antonstar/api-system/pkg/logging"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"regexp"
)

func Setup() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("username", username)
		v.RegisterValidation("password", password)
	}
}

// 任意字符，长度在8-16之间
// TODO: 应该做到能够检测不同强弱规则的密码，参考 https://blog.csdn.net/TCatTime/article/details/106306793
var password validator.Func = func(fl validator.FieldLevel) bool {
	str := fl.Field().Interface().(string)
	matched, err := regexp.MatchString("^.{8,16}$", str)
	if err != nil {
		logging.Error(err)
		return false
	}
	return matched
}

// 字母，数字，下划线组成，长度在5-16之间
var username validator.Func = func(fl validator.FieldLevel) bool {
	str := fl.Field().Interface().(string)
	matched, err := regexp.MatchString("^[a-zA-Z0-9_]{5,16}$", str)
	if err != nil {
		logging.Error(err)
		return false
	}
	return matched
}
