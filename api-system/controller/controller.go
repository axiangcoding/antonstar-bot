package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BaseRedirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/web")
}
