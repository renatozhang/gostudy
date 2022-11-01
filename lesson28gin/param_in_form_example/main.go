package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/user/search", func(ctx *gin.Context) {
		username := ctx.DefaultPostForm("username", "unknown")
		// username := ctx.PostForm("username")
		address := ctx.PostForm("address")
		ctx.JSON(200, gin.H{
			"message":  "pong",
			"username": username,
			"address":  address,
		})
	})

	r.Run() //listen and serve on 0.0.0.0:8080
}
