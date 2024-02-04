package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	server := gin.Default() // Create server and upload index.html
	server.LoadHTMLFiles("index.html")

	server.GET("/", func(ctx *gin.Context) { // Return index.html
		ctx.HTML(
			http.StatusOK,
			"index.html",
			gin.H{},
		)
	})

	server.GET("/time", func(ctx *gin.Context) { // Return current Moscow time in format HH:MM:SS
		location, _ := time.LoadLocation("Europe/Moscow")
		time := time.Now().In(location)
		ctx.JSON(200, gin.H{
			"time": fmt.Sprintf("%.2d:%.2d:%.2d", time.Hour(), time.Minute(), time.Second()),
		})
	})
	server.Run() // Run server
}
