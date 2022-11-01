package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 可以设置一些公共参数
		c.Set("example", "12345")
		// 等其他中间件先执行
		c.Next()
		// 获取耗时
		latency := time.Since(t)
		log.Printf("total cost time:%v", latency)

		// 获取发送的 status
		status := c.Writer.Status()
		log.Print(status)
	}
}

func main() {
	// r := gin.New()
	r := gin.Default()
	r.Use(StatCost())

	r.GET("/test", func(ctx *gin.Context) {
		example := ctx.Copy().MustGet("example").(string)
		//打印“12345”
		log.Println(example)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})

	r.Run()
}
