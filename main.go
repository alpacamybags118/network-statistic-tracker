package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func configureRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	router.GET("/roundtriptime", func(context *gin.Context) {
		fmt.Println("start")
		GetHttpRequestTime()
		context.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return router
}

func main() {
	r := configureRouter()
	r.Run(":8080")
}
