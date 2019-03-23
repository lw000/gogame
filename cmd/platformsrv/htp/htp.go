package htp

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Htp struct {
}

func StartHtpService(port int) {
	engine := gin.Default()
	engine.GET("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"c": 0, "m": "", "d": gin.H{}})
	})
	log.Panic(engine.Run(fmt.Sprintf(":%d", port)))
}
