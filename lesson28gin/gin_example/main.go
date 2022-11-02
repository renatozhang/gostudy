package main

import "github.com/gin-gonic/gin"

func testHandle(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	//Default返回一个默认的路由引擎
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		//输出json结果给调用方
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/test", testHandle)
	// r.Run() //listen and serve on 0.0.0.0:8080
	r.Run(":9090")
}
