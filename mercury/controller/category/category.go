package category

import (
	"github.com/gin-gonic/gin"
	"github.com/renatozhang/gostudy/mercury/dal/db"
	"github.com/renatozhang/gostudy/mercury/util"
)

func GetCategoryListHandle(ctx *gin.Context) {
	categoryList, err := db.GetCategoryList()
	if err != nil {
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		return
	}
	util.ResponseSuccess(ctx, categoryList)
}
