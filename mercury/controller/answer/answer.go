package answer

import (
	"github.com/gin-gonic/gin"
	"github.com/renatozhang/gostudy/mercury/common"
	"github.com/renatozhang/gostudy/mercury/dal/db"
	"github.com/renatozhang/gostudy/mercury/logger"
	"github.com/renatozhang/gostudy/mercury/util"
)

func AnswerListHandle(ctx *gin.Context) {
	questionId, err := util.GetQueryInt64(ctx, "question_id")
	if err != nil {
		logger.Error("get question_id params failed, err:%v", err)
		util.ResponseError(ctx, util.ErrCodeParmeter)
		return
	}
	offset, err := util.GetQueryInt64(ctx, "offset")
	if err != nil {
		logger.Error("get offest params failed, err:%v", err)
		util.ResponseError(ctx, util.ErrCodeParmeter)
		return
	}
	limit, err := util.GetQueryInt64(ctx, "limit")
	if err != nil {
		logger.Error("get limit params failed, err:%v", err)
		util.ResponseError(ctx, util.ErrCodeParmeter)
		return
	}
	logger.Debug("get answer id list parameter succ, qid:%v, offset:%v, limit:%v", questionId, offset, limit)
	answerIdList, err := db.GetAnswerIdList(questionId, offset, limit)
	if err != nil {
		logger.Error("get answerIdList failed, question_id:%v, offset:%v, limit:%v, err:%v", questionId, offset, limit, err)
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		return
	}

	answerList, err := db.MGetAnswer(answerIdList)
	if err != nil {
		logger.Error(" db.MGetAnswer failed, answer_ids:%v err:%v", answerIdList, err)
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		return
	}
	logger.Debug("get answerlist succ, answer_ids:%v, answer_list:%#v", answerIdList, answerList)

	var userIdList []int64
	for _, v := range answerList {
		userIdList = append(userIdList, v.AuthorId)
	}

	userInfoList, err := db.GetUserInfoList(userIdList)
	if err != nil {
		logger.Error(" db.GetUserInfoList failed, userIdList:%v err:%v", userIdList, err)
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		return
	}

	apiAnswerList := common.ApiAnswerList{}
	for _, v := range answerList {
		apiAnswer := &common.ApiAnswer{}
		apiAnswer.Answer = *v
		for _, user := range userInfoList {
			if int64(user.UserId) == v.AuthorId {
				apiAnswer.AuthorName = user.Username
				break
			}
		}
		apiAnswerList.AnswerList = append(apiAnswerList.AnswerList, apiAnswer)
	}
	count, err := db.GetAnswerCount(questionId)
	if err != nil {
		logger.Error(" db.GetAnswerCount failed, question_id:%v err:%v", questionId, err)
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		return
	}
	apiAnswerList.TotalCount = count

	util.ResponseSuccess(ctx, apiAnswerList)
}
