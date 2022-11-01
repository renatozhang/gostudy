package main

import "github.com/gin-gonic/gin"

func main() {
	//Default返回一个默认的路由引擎
	r := gin.Default()
	r.GET("/user/info", func(ctx *gin.Context) {
		//输出Json给调用方
		ctx.JSON(200, gin.H{
			"message": "get user info",
		})
	})
	r.POST("/user/info", func(ctx *gin.Context) {
		//输出Json给调用方
		ctx.JSON(200, gin.H{
			"message": "create user info",
		})
	})
	r.PUT("/user/info", func(ctx *gin.Context) {
		//输出Json给调用方
		ctx.JSON(200, gin.H{
			"message": "update user info",
		})
	})
	r.DELETE("/user/info", func(ctx *gin.Context) {
		//输出Json给调用方
		ctx.JSON(200, gin.H{
			"message": "delete user info",
		})
	})

	r.Run()
}
