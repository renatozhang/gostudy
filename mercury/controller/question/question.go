package question

import (
	"strconv"

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

func QustionDetailHandle(ctx *gin.Context) {
	qustionIdStr, ok := ctx.GetQuery("question_id")
	if !ok {
		logger.Error("invalid qustion_id, not found qustion_id")
		util.ResponseError(ctx, util.ErrCodeParmeter)
		return
	}
	questionId, err := strconv.ParseInt(qustionIdStr, 10, 64)
	if err != nil {
		logger.Error("invalid questionId, strconv.PaseInt failed, err:%v, str:%v", err, qustionIdStr)
		util.ResponseError(ctx, util.ErrCodeParmeter)
		return
	}

	question, err := db.GetQustion(questionId)
	if err != nil {
		logger.Error("get question failed, err:%v, str:%v", err, questionId)
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		return
	}

	categoryMap, err := db.MGetCategory([]int64{question.CategoryId})
	if err != nil {
		logger.Error("get category failed, err:%v, question:%v", err, question)
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		return
	}
	category, ok := categoryMap[question.CategoryId]
	if !ok {
		logger.Error("get category failed, err:%v, question:%v", err, question)
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		return
	}

	userInfoList, err := db.GetUserInfoList([]int64{question.AuthorId})
	if err != nil || len(userInfoList) == 0 {
		logger.Error("get user info list failed, user_ids:%#v, err:%v", question.AuthorId, err)
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		return
	}
	apiQuestionDetail := &common.ApiQuestionDetail{}
	apiQuestionDetail.Question = *question
	apiQuestionDetail.AuthorName = userInfoList[0].Username
	apiQuestionDetail.CategoryName = category.CategoryName
	util.ResponseSuccess(ctx, apiQuestionDetail)

}
