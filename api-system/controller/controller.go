package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RootRedirect(c *gin.Context) {
	c.Redirect(http.StatusSeeOther, "/web")
}
