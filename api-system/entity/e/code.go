package e

const (
	Success = 00000
	Error   = 10000
)

var errCodeText = map[int]string{
	Success: "OK",
	Error:   "System error",
}

func CodeText(code int) string {
	return errCodeText[code]
}
