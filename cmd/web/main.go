package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSONP(http.StatusOK, gin.H{
			"foo": "bar",
		})
	})

	router.Run("127.0.0.1:8080")
}
