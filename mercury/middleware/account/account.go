package account

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatozhang/gostudy/mercury/session"
)

// 处理请求函数
func ProcessRequest(ctx *gin.Context) {
	var userSession session.Session
	var err error
	// 没有获取到用户session，新建一个session
	defer func() {
		if userSession == nil {
			userSession, _ = session.CreateSession()
		}
		ctx.Set(MercurySessionName, userSession)
	}()
	// //从cookie中获取session_id
	cookie, err := ctx.Request.Cookie(CookieSessionId)
	if err != nil {
		//不存在的话，设置user_id=0, login_status为0。表示未登录状态
		ctx.Set(MercuryUserId, int64(0))
		ctx.Set(MercuryUserLoginStatus, int64(0))
		return
	}

	// 从cookie的值中获取sessionId
	sessionId := cookie.Value
	if len(sessionId) == 0 {
		ctx.Set(MercuryUserId, int64(0))
		ctx.Set(MercuryUserLoginStatus, int64(0))
		return
	}
	// 通过sessionMgr 获取session信息
	userSession, err = session.Get(sessionId)
	if err != nil {
		// 获取session失败
		ctx.Set(MercuryUserId, int64(0))
		ctx.Set(MercuryUserLoginStatus, int64(0))

		return
	}
	// 获取用户id user_id
	tmpUserId, err := userSession.Get(MercuryUserId)
	if err != nil {
		ctx.Set(MercuryUserId, int64(0))
		ctx.Set(MercuryUserLoginStatus, int64(0))
		return
	}
	userId, ok := tmpUserId.(int64)
	if !ok || userId == 0 {
		// 没有转成功设置用户id与登录状态到gin框架中
		ctx.Set(MercuryUserId, int64(0))
		ctx.Set(MercuryUserLoginStatus, int64(0))
		return
	}

	// 设置用户id与登录状态到gin框架中
	ctx.Set(MercuryUserId, int64(userId))
	ctx.Set(MercuryUserLoginStatus, int64(1))
}

// 获取user_id （从gin框架的通用参数中取）
func GetUserId(ctx *gin.Context) (userId int64, err error) {
	tmpUserId, exists := ctx.Get(MercuryUserId)
	// 不存在返回错误
	if !exists {
		err = errors.New("user id not exists")
		return
	}
	// 如果存在转为int64
	userId, ok := tmpUserId.(int64)
	if !ok {
		err = errors.New("user id convert to int64 failed")
		return
	}
	return
}

// 获取用户是否登录login_status（从gin框架的通用参数中取）
func IsLogin(ctx *gin.Context) (login bool) {
	tmpIsLoginStatus, exists := ctx.Get(MercuryUserLoginStatus)
	if !exists {
		return
	}
	LoginStatus, ok := tmpIsLoginStatus.(int64)
	if !ok {
		return
	}
	if LoginStatus == 0 {
		return
	}
	login = true
	return
}

func SetUserId(userId int64, ctx *gin.Context) {
	var userSession session.Session
	tempSession, exists := ctx.Get(MercurySessionName)
	if !exists {
		return
	}
	userSession, ok := tempSession.(session.Session)
	if !ok {
		return
	}
	if userSession == nil {
		return
	}

	userSession.Set(MercuryUserId, userId)
}

// 处理请求返回（种cookie）
func ProcessResponse(ctx *gin.Context) {
	var userSession session.Session
	// 获取用户session mercury_session
	tempSession, exists := ctx.Get(MercurySessionName)
	if !exists {
		return
	}
	userSession, ok := tempSession.(session.Session)
	if !ok {
		return
	}
	// 是否有修改，有修改保存
	if !userSession.IsModify() {
		return
	}
	err := userSession.Save()
	if err != nil {
		return
	}
	sessionId := userSession.Id()
	// 初始化cookie
	cookie := &http.Cookie{
		Name:     CookieSessionId,
		Value:    sessionId,
		MaxAge:   CookieMaxAge,
		HttpOnly: true,
		Path:     "/",
	}
	// 设置cookie返回到浏览器
	http.SetCookie(ctx.Writer, cookie)
}
