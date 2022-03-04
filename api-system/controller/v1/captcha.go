package v1

import (
	"axiangcoding/antonstar/api-system/entity/app"
	"axiangcoding/antonstar/api-system/entity/e"
	"axiangcoding/antonstar/api-system/service"
	"bytes"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"time"
)

// GenerateCaptcha
// @Summary  请求生成验证码
// @Tags     Captcha API
// @Success  200  {object}  app.ApiJson  ""
// @Router   /v1/captcha [get]
func GenerateCaptcha(c *gin.Context) {
	captchaMap := service.GenerateCaptcha()
	app.Success(c, captchaMap)
}

type GetCaptchaForm struct {
	Reload bool   `binding:"omitempty" json:"reload" form:"reload"`
	Lang   string `binding:"" json:"lang" form:"lang"`
}

// GetCaptcha
// @Summary  获取验证码图片
// @Tags     Captcha API
// @Param    file  path   string          true   "image file name"
// @Param    form  query  GetCaptchaForm  false  "more options"
// @Accept   png
// @Success  200
// @Router   /v1/captcha/{file} [get]
func GetCaptcha(c *gin.Context) {
	file := c.Param("file")
	var form GetCaptchaForm
	err := c.ShouldBindQuery(&form)
	if err != nil {
		app.BadRequest(c, e.RequestParamsNotValid, err)
		return
	}
	ServeHTTP(c.Writer, c.Request, file, form.Reload, form.Lang)
}

func ServeHTTP(w http.ResponseWriter, r *http.Request, file string, reload bool, lang string) {
	ext := path.Ext(file)
	id := file[:len(file)-len(ext)]
	if ext == "" || id == "" {
		http.NotFound(w, r)
		return
	}
	if reload {
		captcha.Reload(id)
	}
	if Serve(w, r, id, ext, lang, false, captcha.StdWidth, captcha.StdHeight) == captcha.ErrNotFound {
		http.NotFound(w, r)
	}
}

func Serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	var content bytes.Buffer
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		err := captcha.WriteImage(&content, id, width, height)
		if err != nil {
			return err
		}
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		err := captcha.WriteAudio(&content, id, lang)
		if err != nil {
			return err
		}
	default:
		return captcha.ErrNotFound
	}

	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}
