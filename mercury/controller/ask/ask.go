package ask

import (
	"github.com/gin-gonic/gin"
	"github.com/renatozhang/gostudy/mercury/common"
	"github.com/renatozhang/gostudy/mercury/dal/db"
	"github.com/renatozhang/gostudy/mercury/filter"
	"github.com/renatozhang/gostudy/mercury/id_gen"
	"github.com/renatozhang/gostudy/mercury/logger"
	"github.com/renatozhang/gostudy/mercury/middleware/account"
	"github.com/renatozhang/gostudy/mercury/util"
)

func QuestionSubmitHandle(ctx *gin.Context) {
	var question common.Question
	err := ctx.BindJSON(&question)
	if err != nil {
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		return
	}
	logger.Debug("bind json succ,question:%#v", question)
	result, hit := filter.Replace(question.Caption, "***")
	if hit {
		logger.Error("caption is filter, result:%v", result)
		util.ResponseError(ctx, util.ErrCodeCaptionHit)
		return
	}

	result, hit = filter.Replace(question.Content, "***")
	if hit {
		logger.Debug("content is filter, result:%v", result)
		util.ResponseError(ctx, util.ErrCodeContentHit)
		return
	}
	logger.Debug("filter succ, result:%#v", result)
	qid, err := id_gen.GetID()
	if err != nil {
		logger.Error("generate question id failed, err:%v", err)
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		return
	}
	question.QuestionId = int64(qid)
	userId, err := account.GetUserId(ctx)
	if err != nil || userId <= 0 {
		logger.Error("user is not login, err:%v", err)
		util.ResponseError(ctx, util.ErrCodeNotLogin)
		return
	}
	question.AuthorId = userId
	err = db.CreateQuestion(&question)
	if err != nil {
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		return
	}
	logger.Debug("create question succ, question:%#v", question)
	util.ResponseSuccess(ctx, nil)
}
