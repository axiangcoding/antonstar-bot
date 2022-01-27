package e

const (
	Success               = 00000
	Error                 = 10000
	RequestParamsNotValid = 10001

	TokenNotExist = 11000
	TokenNotValid = 11001
	TokenExpired  = 11002
	NoPermission  = 11003

	LoginFailed    = 12000
	RegisterFailed = 12001

	ReachRefreshLimit = 13000
)

var errCodeText = map[int]string{
	Success:               "OK",
	Error:                 "System error",
	RequestParamsNotValid: "Request params not valid",

	TokenNotExist: "Authorization not exist",
	TokenNotValid: "Authorization not valid",
	TokenExpired:  "Authorization expired",
	NoPermission:  "No permission",

	LoginFailed:       "Login failed",
	RegisterFailed:    "Register failed",
	ReachRefreshLimit: "Reach Refresh Limit",
}

func CodeText(code int) string {
	return errCodeText[code]
}
