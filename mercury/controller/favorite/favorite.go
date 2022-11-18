package favorite

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/renatozhang/gostudy/mercury/common"
	"github.com/renatozhang/gostudy/mercury/dal/db"
	"github.com/renatozhang/gostudy/mercury/id_gen"
	"github.com/renatozhang/gostudy/mercury/logger"
	"github.com/renatozhang/gostudy/mercury/middleware/account"
	"github.com/renatozhang/gostudy/mercury/util"
)

func AddDirHandle(ctx *gin.Context) {
	var favoriteDir common.FavoriteDir
	err := ctx.BindJSON(&favoriteDir)
	if err != nil {
		util.ResponseError(ctx, util.ErrCodeParmeter)
		return
	}
	favoriteDir.DirName = strings.TrimSpace(favoriteDir.DirName)
	if len(favoriteDir.DirName) == 0 {
		util.ResponseError(ctx, util.ErrCodeParmeter)
		logger.Error("invalid dir name:%v", favoriteDir.DirName)
		return
	}
	logger.Debug("bind json succ,favoriteDir:%#v", favoriteDir)

	dir_id, err := id_gen.GetID()
	if err != nil {
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		logger.Error("id_gen.GetID failed, favoriteDir:%#v, err:%v", favoriteDir, err)
		return
	}
	favoriteDir.DirId = int64(dir_id)

	userId, err := account.GetUserId(ctx)
	if err != nil || userId == 0 {
		util.ResponseError(ctx, util.ErrCodeNotLogin)
		return
	}
	favoriteDir.UserId = userId
	err = db.CreateFavoriteDir(&favoriteDir)
	if err != nil {
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		logger.Error("CreateFavoriteDir failed, favoriteDir:%#v, err:%v", favoriteDir, err)
		return
	}
	util.ResponseSuccess(ctx, nil)
}

func AddFavoriteHandle(ctx *gin.Context) {
	var favorite common.Favorite
	err := ctx.BindJSON(&favorite)
	if err != nil {
		util.ResponseError(ctx, util.ErrCodeParmeter)
		return
	}
	if favorite.AnswerId == 0 || favorite.DirId == 0 {
		util.ResponseError(ctx, util.ErrCodeParmeter)
		logger.Error("invalid dir id favorite:%#v, err:%v", favorite, err)
		return
	}

	logger.Debug("bind json succ,favorite:%#v", favorite)

	userId, err := account.GetUserId(ctx)
	if err != nil || userId == 0 {
		util.ResponseError(ctx, util.ErrCodeNotLogin)
		return
	}
	favorite.UserId = userId
	err = db.CreateFavorite(&favorite)
	if err != nil {
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		logger.Error("CreateFavoriteDir failed, favorite:%#v, err:%v", favorite, err)
		return
	}
	util.ResponseSuccess(ctx, nil)

}

func DirListHandle(ctx *gin.Context) {
	userId, err := account.GetUserId(ctx)
	if err != nil || userId == 0 {
		util.ResponseError(ctx, util.ErrCodeNotLogin)
		return
	}

	favoriteDirList, err := db.GetFavoriteDirList(userId)
	if err != nil {
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		logger.Error("db.GetFavoriteDirList failed, err:%v", err)
		return
	}
	util.ResponseSuccess(ctx, favoriteDirList)
}

func FavoriteListHandle(ctx *gin.Context) {
	userId, err := account.GetUserId(ctx)
	if err != nil || userId == 0 {
		util.ResponseError(ctx, util.ErrCodeNotLogin)
		return
	}
	dirIdStr, ok := ctx.GetQuery("dir_id")
	dirIdStr = strings.TrimSpace(dirIdStr)
	if !ok || len(dirIdStr) == 0 {
		util.ResponseError(ctx, util.ErrCodeParmeter)
		logger.Error("valid comment id, val:%v, ok:%v", dirIdStr, ok)
		return
	}
	logger.Debug("get query dir_id succ, val:%v", dirIdStr)
	dirId, err := strconv.ParseInt(dirIdStr, 10, 64)
	if err != nil || dirId == 0 {
		util.ResponseError(ctx, util.ErrCodeParmeter)
		logger.Error("valid dir id, val:%v, err:%v", dirId, err)
		return
	}

	var offset int64
	offsetStr, ok := ctx.GetQuery("offset")
	offsetStr = strings.TrimSpace(offsetStr)
	if !ok || len(offsetStr) == 0 {
		offset = 0
		logger.Error("get offset failed, err:%v", err)
	}
	offset, err = strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		offset = 0
		logger.Error("valid dir id, val:%v, err:%v", offset, err)
	}
	logger.Debug("get offset succ, val:%v", offset)
	var limit int64
	limitStr, ok := ctx.GetQuery("offset")
	limitStr = strings.TrimSpace(limitStr)
	if !ok || len(limitStr) == 0 {
		limit = 10
		logger.Error("get limit failed, err:%v", err)
	}
	limit, err = strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		limit = offset + 10
		logger.Error("get limit failed, err:%v", err)
	}
	logger.Debug("get limit succ, val:%v", limit)
	favoriteList, err := db.GetFavoriteList(dirId, userId, offset, limit)
	if err != nil {
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		logger.Error("db.GetFavoriteList failed, err:%v", err)
		return
	}
	var answerIdList []int64
	for _, v := range favoriteList {
		answerIdList = append(answerIdList, v.AnswerId)
	}
	answerList, err := db.MGetAnswer(answerIdList)
	if err != nil {
		logger.Error("db.MGetAnswer failed, answer_ids:%v err:%v", answerIdList, err)
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		return
	}
	var userIdList []int64
	for _, v := range favoriteList {
		userIdList = append(userIdList, v.UserId)
	}
	userinfoList, err := db.GetUserInfoList(userIdList)
	if err != nil {
		logger.Error("db.GetUserInfoList failed, answer_ids:%v err:%v", answerIdList, err)
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		return
	}
	apianswerList := &common.ApiAnswerList{}
	for _, v := range answerList {
		apiAnswer := &common.ApiAnswer{}
		apiAnswer.Answer = *v
		for _, user := range userinfoList {
			if int64(user.UserId) == v.AuthorId {
				apiAnswer.AuthorName = user.Username
				break
			}
		}
		apianswerList.AnswerList = append(apianswerList.AnswerList, apiAnswer)
	}
	apianswerList.TotalCount = int32(len(favoriteList))
	util.ResponseSuccess(ctx, apianswerList)

}
