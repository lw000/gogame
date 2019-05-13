package ggmiddleware

import (
	log "github.com/alecthomas/log4go"
	"github.com/gin-contrib/cors"
	"github.com/unrolled/secure"
	"time"
	"tuyue/tuyue_common/web/whitelist"

	"github.com/gin-gonic/gin"
	"net/http"
	// "github.com/gorilla/securecookie"
	// "github.com/gorilla/csrf"
)

func CorsHandler(originUrls map[string]bool) gin.HandlerFunc {
	cfg := cors.DefaultConfig()
	cfg.AllowOriginFunc = func(origin string) bool {
		// allowed, ok := originUrls[origin]
		// return ok && allowed
		return true
	}
	cfg.AllowOrigins = []string{"*"}
	cfg.AllowMethods = []string{"POST", "GET"}
	cfg.AllowCredentials = true
	return cors.New(cfg)
}

func WhiteListHanlder(wl *tywhitelist.WhiteList) gin.HandlerFunc {
	return func(c *gin.Context) {
		if er := wl.CheckWhiteList(c.Writer, c.Request); er != nil {
			log.Error("[%s] %s", c.ClientIP(), er.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"c": http.StatusUnauthorized, "error": er.Error()})
			c.Abort()
		}
		c.Next()
	}
}

func LoggerHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Stop timer
		end := time.Now()
		latency := end.Sub(start)

		method := c.Request.Method
		statusCode := c.Writer.Status()

		if raw != "" {
			path = path + "?" + raw
		}

		log.Info("|%3d| %13v |%-7s| %s",
			statusCode,
			latency,
			method,
			path,
		)
	}
}

func TlsHandler(host string) gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     host,
		})

		er := secureMiddleware.Process(c.Writer, c.Request)
		if er != nil {
			log.Error(er)
			return
		}

		c.Next()
	}
}
