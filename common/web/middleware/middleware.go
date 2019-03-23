package ggmiddleware

import (
	log "github.com/alecthomas/log4go"
	"tuyue/tuyue_common/web/whitelist"

	"github.com/gin-gonic/gin"
	"net/http"
	// "github.com/gorilla/securecookie"
	// "github.com/gorilla/csrf"
)

func WhiteList(wl *tywhitelist.WhiteList) gin.HandlerFunc {
	return func(c *gin.Context) {
		if er := wl.CheckWhiteList(c.Writer, c.Request); er != nil {
			log.Error("[%s] %s", c.ClientIP(), er.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"c": http.StatusUnauthorized, "error": er.Error()})
			c.Abort()
		}
		c.Next()
	}
}
