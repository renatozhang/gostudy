package category

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/renatozhang/gostudy/mercury/common"
	"github.com/renatozhang/gostudy/mercury/dal/db"
	"github.com/renatozhang/gostudy/mercury/logger"
	"github.com/renatozhang/gostudy/mercury/util"
)

func GetCategoryListHandle(ctx *gin.Context) {
	categoryList, err := db.GetCategoryList()
	if err != nil {
		logger.Error("db.GetCategoryList failed. err:%v", err)
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		return
	}
	util.ResponseSuccess(ctx, categoryList)
}

func GetQustionListHandle(ctx *gin.Context) {
	categoryIdStr, ok := ctx.GetQuery("category_id")
	if !ok {
		logger.Error("invalid category_id, not found category_id")
		util.ResponseError(ctx, util.ErrCodeParmeter)
		return
	}
	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
	if err != nil {
		logger.Error("invalid category_id, strconv.PaseInt failed, err:%v, str:%v", err, categoryIdStr)
		util.ResponseError(ctx, util.ErrCodeParmeter)
		return
	}

	questionList, err := db.GetQuestionList(categoryId)
	if err != nil {
		logger.Error("get question list failed, category_id:%v, err:%v", categoryId, err)
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		return
	}

	if len(questionList) == 0 {
		logger.Warn("get question list succ, empty list,category_id:%v", categoryId)
		util.ResponseSuccess(ctx, questionList)
		return
	}

	var userIdList []int64
	userIdMap := make(map[int64]bool, 16)
	for _, question := range questionList {
		_, ok := userIdMap[question.AuthorId]
		if ok {
			continue
		}
		userIdMap[question.AuthorId] = true
		userIdList = append(userIdList[:], question.AuthorId)
	}

	userInfoList, err := db.GetUserInfoList(userIdList)
	if err != nil {
		logger.Error("get user info list failed, user_ids:%#v, err:%v", userIdList, err)
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		return
	}
	var apiQuestionList []*common.ApiQuestion
	for _, question := range questionList {
		var apiQuestion = &common.ApiQuestion{}
		apiQuestion.Question = *question
		apiQuestion.CreateTimeStr = apiQuestion.CreeateTime.Format("2006/1/2 15:04:05")
		for _, userInfo := range userInfoList {
			if question.AuthorId == int64(userInfo.UserId) {
				apiQuestion.AuthorName = userInfo.NickName
				break
			}
		}
		apiQuestionList = append(apiQuestionList, apiQuestion)
	}

	util.ResponseSuccess(ctx, apiQuestionList)

}
