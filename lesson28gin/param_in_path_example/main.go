package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/user/search/:username/:address", func(ctx *gin.Context) {
		username := ctx.Param("username")
		address := ctx.Param("address")
		ctx.JSON(200, gin.H{
			"message":  "pong",
			"username": username,
			"address":  address,
		})
	})

	r.Run()
}
