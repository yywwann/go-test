package main

import (
	"net/http"
	"time"

	"example/midware_test/midware"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	// 新建一个限速器，允许突发 10 个并发，限速 3rps，超过 500ms 就不再等待
	e.Use(midware.NewLimiter(3, 10, 500*time.Millisecond))
	e.GET("ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	e.Run(":8080")
}
