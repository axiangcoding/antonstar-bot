package e

const (
	Success               = 00000
	Error                 = 10000
	RequestParamsNotValid = 10001

	TokenNotExist   = 11000
	TokenNotValid   = 11001
	TokenExpired    = 11002
	NoPermission    = 11003
	CaptchaNotValid = 11004

	UserLoginFailed        = 12000
	UserRegisterFailed     = 12001
	UserExist              = 12002
	UserPasswordNotMatched = 12003
)

var errCodeText = map[int]string{
	Success:               "OK",
	Error:                 "System error",
	RequestParamsNotValid: "Request params not valid",

	TokenNotExist:   "Authorization not exist",
	TokenNotValid:   "Authorization not valid",
	TokenExpired:    "Authorization expired",
	NoPermission:    "No permission",
	CaptchaNotValid: "Captcha not valid",

	UserLoginFailed:        "Login failed",
	UserRegisterFailed:     "Register failed",
	UserExist:              "User Exist",
	UserPasswordNotMatched: "User username or password not matched",
}

func CodeText(code int) string {
	return errCodeText[code]
}
