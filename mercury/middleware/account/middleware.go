package account

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func StatCost() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()
		//可以设置一些公共参数
		ctx.Set("example", "12345")
		// 等其他中间件执行
		ctx.Next()
		latency := time.Since(t)
		log.Printf("total cost time:%d us", latency/1000)
	}
}
