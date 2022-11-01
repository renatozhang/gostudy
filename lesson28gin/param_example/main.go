package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/user/search", func(ctx *gin.Context) {
		username := ctx.DefaultQuery("username", "unknown")
		// username := ctx.Query("username")
		fmt.Println(username)
		address := ctx.Query("address")
		ctx.JSON(200, gin.H{
			"message":  "pong",
			"username": username,
			"address":  address,
		})
	})

	r.Run() //listen and serve 0.0.0.0:8080
}
