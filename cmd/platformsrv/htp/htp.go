package htp

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type PlatformHtp struct {
	port   int64
	engine *gin.Engine
}

func (h *PlatformHtp) Start(port int64) error {
	h.port = port
	h.engine = gin.New()
	h.engine.Use(gin.Logger())
	h.engine.Use(gin.Recovery())

	h.engine.GET("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"c": 0, "m": "", "d": gin.H{}})
	})

	h.engine.GET("/gamelist", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"c": 0, "m": "", "d": gin.H{}})
	})

	go h.run()

	return nil
}

func (h *PlatformHtp) Stop() error {

	return nil
}

func (h *PlatformHtp) run() {
	log.Panic(h.engine.Run(fmt.Sprintf(":%d", h.port)))
}
