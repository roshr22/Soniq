package main

import (
	"github.com/gin-gonic/gin"
	"soniq/internal/server"
	"soniq/internal/server/handlers"
	"net/http"
)

func main() {
	r := gin.Default()

	server.StartBroadcastLoop()

	r.Static("/uploads", "./public/uploads")

	r.GET("/ws", server.HandleWebSocket)
	r.POST("/upload", handlers.UploadAudio)

	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.Run("0.0.0.0:8080")
}
