package v1

import (
	"axiangcoding/antonstar/api-system/pkg/logging"

	"github.com/gin-gonic/gin"
)

// DemoGet
// @Summary   demo，测试get
// @Produce   json
// @Tags      Demo/Test
// @Param     param1  query     string  false  "some params named param1"
// @Param     param2  query     string  false  "some params named param2"
// @Success   200     {string}  json    ""
// @Router    /v1/demo/get [get]
// @Security  ApiKeyAuth
func DemoGet(c *gin.Context) {
	param1 := c.Query("param1")
	param2 := c.Query("param2")
	c.JSON(200, gin.H{
		"method": "get",
		"param1": param1,
		"param2": param2,
	})
}

type Params struct {
	Param1 string `json:"param1"`
	Param2 string `json:"param2"`
}

// DemoPost
// @Summary   demo，测试post
// @Produce   json
// @Tags      demo
// @Param     params  body      Params  false  "some params json"
// @Success   200     {string}  json    ""
// @Router    /v1/demo/post [post]
// @Security  ApiKeyAuth
func DemoPost(c *gin.Context) {
	params := Params{}
	c.ShouldBindJSON(&params)
	c.JSON(200, gin.H{
		"method":    "post",
		"post_body": params,
	})
}

// TestLog
// @Summary   demo，测试post
// @Produce   json
// @Tags      demo
// @Success   200  {string}  json  ""
// @Router    /v1/test/test-log [get]
// @Security  ApiKeyAuth
func TestLog(c *gin.Context) {
	logging.Debug("this is a debug log")
	logging.Info("this is a info log")
	logging.Info("this is a info log with params", "value1")
	logging.Warn("this is a warn log")
}
