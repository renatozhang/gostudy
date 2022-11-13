package account

import (
	"github.com/gin-gonic/gin"
	"github.com/renatozhang/gostudy/mercury/util"
)

func AuthMiddleware(ctx *gin.Context) {
	ProcessRequest(ctx)
	isLogin := IsLogin(ctx)
	if !isLogin {
		util.ResponseError(ctx, util.ErrCodeNotLogin)
		// 中止当前请求
		ctx.Abort()
		return
	}
	ctx.Next()
}
