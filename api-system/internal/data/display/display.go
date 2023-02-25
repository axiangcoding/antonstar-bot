package display

import (
	"bytes"
	"github.com/axiangcoding/antonstar-bot/pkg/logging"
	"text/template"
)

func parseTemplate(templateStr string, u any) string {
	var buf bytes.Buffer
	t, err := template.New("").Parse(templateStr)
	if err != nil {
		logging.L().Error("parse template failed", logging.Error(err))
		return "Error"
	}
	if err := t.Execute(&buf, u); err != nil {
		logging.L().Error("exec template failed", logging.Error(err))
		return "Error"
	}
	return buf.String()
}
