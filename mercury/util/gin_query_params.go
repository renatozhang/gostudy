package util

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/renatozhang/gostudy/mercury/logger"
)

func GetQueryInt64(ctx *gin.Context, key string) (value int64, err error) {
	idstr, ok := ctx.GetQuery(key)
	if !ok {
		logger.Error("invalid %d, not found!")
		err = fmt.Errorf("invalid params, not found key:%s", key)
		return
	}
	value, err = strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		logger.Error("invalid params, strconv.PaseInt failed, err:%v, str:%v", err, idstr)
		return
	}
	return
}
