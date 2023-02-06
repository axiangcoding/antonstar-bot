package static

import (
	"embed"
	"github.com/axiangcoding/antonstar-bot/logging"
)

//go:embed message
var fs embed.FS

func MustReadMessageFileAsBytes(filename string) []byte {
	bytes, err := fs.ReadFile("message/" + filename)
	if err != nil {
		logging.Warnf("read message file {%s} error.", filename)
		return nil
	}
	return bytes
}
