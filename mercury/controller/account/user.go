package account

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/renatozhang/gostudy/mercury/common"
	"github.com/renatozhang/gostudy/mercury/dal/db"
	"github.com/renatozhang/gostudy/mercury/id_gen"
	"github.com/renatozhang/gostudy/mercury/middleware/account"
	"github.com/renatozhang/gostudy/mercury/util"
)

func LoginHandle(ctx *gin.Context) {
	account.ProcessRequest(ctx)
	var userInfo common.UserInfo
	var err error
	defer func() {
		if err != nil {
			return
		}
		// 用户登录成功之后，需要把user_id设置到session中
		account.SetUserId(int64(userInfo.UserId), ctx)
		//当调用responseSuccess的时候，gin框架已经把数据发送给浏览器了
		//所以在responseSuccess之后，SetCookie就不会生效。因此，account.ProcessResponse
		//必须在util.ResponseSuccess之前调用
		account.ProcessResponse(ctx)
		util.ResponseSuccess(ctx, err)
	}()
	err = ctx.BindJSON(&userInfo)
	if err != nil {
		util.ResponseError(ctx, util.ErrCodeParmeter)
		return
	}

	if len(userInfo.Username) == 0 && len(userInfo.Password) == 0 {
		util.ResponseError(ctx, util.ErrCodeParmeter)
		return
	}

	err = db.Login(&userInfo)
	if err == db.ErrUserNotExists {
		util.ResponseError(ctx, util.ErrCodeUserNotExists)
		return
	}
	if err == db.ErrUserPasswordWrong {
		util.ResponseError(ctx, util.ErrCodeUserPasswordWrong)
		return
	}
	if err != nil {
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		return
	}
}

func RegisterHandle(ctx *gin.Context) {
	var userInfo common.UserInfo
	err := ctx.BindJSON(&userInfo)
	if err != nil {
		util.ResponseError(ctx, util.ErrCodeParmeter)
		return
	}
	if len(userInfo.Email) == 0 || len(userInfo.Password) == 0 || len(userInfo.Username) == 0 {
		util.ResponseError(ctx, util.ErrCodeParmeter)
		return
	}
	// sex=1表示男生， sex=2表示女生
	if userInfo.Sex != common.UserSexMan && userInfo.Sex != common.UserSexWomen {
		util.ResponseError(ctx, util.ErrCodeParmeter)
		return
	}

	userInfo.UserId, err = id_gen.GetID()
	if err != nil {
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		return
	}
	err = db.Register(&userInfo)
	if err == db.ErrUserExists {
		util.ResponseError(ctx, util.ErrCodeUserExist)
		return
	}
	fmt.Println("db err:", err)
	if err != nil {
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		return
	}
	util.ResponseSuccess(ctx, nil)
}
