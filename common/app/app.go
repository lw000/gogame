package ggapp

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http/pprof"
)

type WebApplication struct {
	engine      *gin.Engine
	debug       int64
	port        int64
	enableTLS   bool
	tlsCertFile string // TLS 证书文件
	tlsKeyFile  string // TLS key文件
}

func NewApplication(debug int64) *WebApplication {
	if debug == 1 {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	return &WebApplication{engine: gin.New(), debug: debug}
}

func (app *WebApplication) Port() int64 {
	return app.port
}

func (app *WebApplication) EnableTLS() bool {
	return app.enableTLS
}

func (app *WebApplication) SetEnableTLS(enableTLS bool) {
	app.enableTLS = enableTLS
}

func (app *WebApplication) Debug() int64 {
	return app.debug
}

func (app *WebApplication) SetDebug(debug int64) {
	app.debug = debug
}

func (app *WebApplication) Engine() *gin.Engine {
	return app.engine
}

func (app *WebApplication) SetTlsFile(tlsCertFile, tlsKeyFile string) {
	app.tlsCertFile = tlsCertFile
	app.tlsKeyFile = tlsKeyFile
}

func (app *WebApplication) Run(port int64, f func(app *WebApplication)) error {
	app.port = port

	if app.debug == 1 {
		app.engine.GET("/debug/pprof/", func(c *gin.Context) {
			pprof.Index(c.Writer, c.Request)
		})

		app.engine.GET("/debug/pprof/cmdline", func(c *gin.Context) {
			pprof.Cmdline(c.Writer, c.Request)
		})

		app.engine.GET("/debug/pprof/profile", func(c *gin.Context) {
			pprof.Profile(c.Writer, c.Request)
		})

		app.engine.GET("/debug/pprof/symbol", func(c *gin.Context) {
			pprof.Symbol(c.Writer, c.Request)
		})

		app.engine.GET("/debug/pprof/trace", func(c *gin.Context) {
			pprof.Trace(c.Writer, c.Request)
		})
	}

	f(app)

	if app.enableTLS {
		app.runTLS(port)
	} else {
		app.runHttp(port)
	}

	return nil
}

func (app *WebApplication) runHttp(port int64) {
	var er error
	er = app.engine.Run(fmt.Sprintf(":%d", port))
	if er != nil {
		log.Panic(er)
	}
}

func (app *WebApplication) runTLS(port int64) {
	var er error
	er = app.engine.RunTLS(fmt.Sprintf(":%d", port), app.tlsCertFile, app.tlsKeyFile)
	if er != nil {
		log.Panic(er)
	}
}
