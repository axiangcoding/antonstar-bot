package v1

import (
	"axiangcoding/antonstar/api-system/entity/app"
	"axiangcoding/antonstar/api-system/entity/e"
	"axiangcoding/antonstar/api-system/service"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

// allowPictureContentType 运行的类型
var allowPictureContentType = []string{"image/png", "image/jpeg", "image/jpg"}

// UploadPicture
// @Summary   上传图片
// @Tags      Upload API
// @Param     file  formData  file  true  "image file"
// @Accept    mpfd
// @Success   200  {object}  app.ApiJson  ""
// @Router    /v1/upload/picture [post]
// @Security  ApiKeyAuth
func UploadPicture(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		app.BadRequest(c, e.FileFormNotValid, err)
		return
	}
	src, err := file.Open()
	if err != nil {
		app.BizFailed(c, e.FileCantBeOpened, err)
		return
	}
	defer src.Close()
	fBytes, err := ioutil.ReadAll(src)
	contentType := http.DetectContentType(fBytes)

	allow := false
	for _, ct := range allowPictureContentType {
		if contentType == ct {
			allow = true
		}
	}
	if !allow {
		app.BizFailed(c, e.FileTypeNotAllowed)
		return
	}

	url, err := service.UploadToSuperBed(fBytes)
	if err != nil {
		app.BizFailed(c, e.FileCantBeSaved, err)
		return
	}
	app.Success(c, url)

}
