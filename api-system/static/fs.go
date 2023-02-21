package static

import (
	"embed"
	"github.com/axiangcoding/antonstar-bot/pkg/logging"
)

//go:embed message
var fs embed.FS

func MustReadMessageFileAsBytes(filename string) []byte {
	bytes, err := fs.ReadFile("message/" + filename)
	if err != nil {
		logging.L().Warn("read message file error.",
			logging.Any("filename", filename),
			logging.Error(err))
		return nil
	}
	return bytes
}
