package comment

import (
	"html"
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

const (
	MinCommentContentSize = 10
)

func PostCommentHandle(ctx *gin.Context) {
	var comment common.Comment
	err := ctx.BindJSON(&comment)
	if err != nil {
		util.ResponseError(ctx, util.ErrCodeParmeter)
		return
	}
	logger.Debug("bind json succ,comment:%#v", comment)
	if len(comment.Content) <= MinCommentContentSize || comment.QuestionId == 0 {
		util.ResponseError(ctx, util.ErrCodeParmeter)
		logger.Error("len(comment.Content):%v, qid:%v. invalid param.", len(comment.Content), comment.QuestionId)
		return
	}

	userId, err := account.GetUserId(ctx)
	if err != nil || userId == 0 {
		util.ResponseError(ctx, util.ErrCodeNotLogin)
		return
	}
	comment.AuthorId = userId
	// 1. 针对content做一个转义，防止xss攻击
	comment.Content = html.EscapeString(comment.Content)
	// 2.生成评论的id
	commentId, err := id_gen.GetID()
	if err != nil {
		util.ResponseError(ctx, util.ErrCodeParmeter)
		logger.Error("id_gen.GetId failed, comment:%#v, err:%v", comment, err)
		return
	}
	comment.CommentId = int64(commentId)

	err = db.CreatePostComment(&comment)
	if err != nil {
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		logger.Error("CreatePostComment failed, comment:%#v, err:%v", comment, err)
		return
	}
	util.ResponseSuccess(ctx, nil)
}

func PostReplyHandle(ctx *gin.Context) {
	var comment common.Comment
	err := ctx.BindJSON(&comment)
	if err != nil {
		util.ResponseError(ctx, util.ErrCodeParmeter)
		return
	}
	logger.Debug("bind json succ,comment:%#v", comment)
	if len(comment.Content) <= MinCommentContentSize || comment.QuestionId == 0 || comment.ParentId == 0 || comment.ReplyCommentId == 0 {
		util.ResponseError(ctx, util.ErrCodeParmeter)
		logger.Error("len(comment.Content):%v, qid:%v. invalid param.", len(comment.Content), comment.QuestionId)
		return
	}

	userId, err := account.GetUserId(ctx)
	if err != nil || userId == 0 {
		util.ResponseError(ctx, util.ErrCodeNotLogin)
		return
	}

	comment.AuthorId = userId
	// 1. 针对content做一个转义，防止xss攻击
	comment.Content = html.EscapeString(comment.Content)
	// 2.生成评论的id
	commentId, err := id_gen.GetID()
	if err != nil {
		util.ResponseError(ctx, util.ErrCodeParmeter)
		logger.Error("id_gen.GetId failed, comment:%#v, err:%v", comment, err)
		return
	}
	comment.CommentId = int64(commentId)
	// 3. 根据replyCommentId,查询这个ReplyCommentId的author_id,也就是ReplyAuthorId
	err = db.CreateReplyComment(&comment)
	if err != nil {
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		logger.Error("CreateReplyComment failed, comment:%#v, err:%v", comment, err)
		return
	}
	util.ResponseSuccess(ctx, nil)
}

func CommentListHandle(ctx *gin.Context) {
	answerIdStr, ok := ctx.GetQuery("answer_id")
	answerIdStr = strings.TrimSpace(answerIdStr)
	if !ok || len(answerIdStr) == 0 {
		util.ResponseError(ctx, util.ErrCodeParmeter)
		logger.Error("valid answer id, val:%v, ok:%v", answerIdStr, ok)
		return
	}
	logger.Debug("get query answer_id succ, val:%v", answerIdStr)
	answerId, err := strconv.ParseInt(answerIdStr, 10, 64)
	if err != nil || answerId == 0 {
		util.ResponseError(ctx, util.ErrCodeParmeter)
		logger.Error("valid answer id, val:%v, err:%v", answerId, err)
		return
	}

	//解析offset
	var offset int64
	offsetStr, ok := ctx.GetQuery("offset")
	offsetStr = strings.TrimSpace(offsetStr)
	if !ok || len(offsetStr) == 0 {
		offset = 0
		logger.Error("valid offset, val:%v, err:%v", offsetStr, err)
	}
	offset, err = strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		offset = 0
		logger.Error("invalid offset, val:%v, err:%v", offsetStr, err)
	}
	logger.Debug("get query offset succ, val:%v", offset)
	//解析limit
	var limit int64
	limitStr, ok := ctx.GetQuery("limit")
	limitStr = strings.TrimSpace(limitStr)
	if !ok || len(limitStr) == 0 {
		limit = 10
		logger.Error("invalid limit, val:%v", limitStr)
	}
	limit, err = strconv.ParseInt(limitStr, 10, 64)
	if err != nil || limit == 0 {
		limit = 10
		logger.Error("invalid limit, val:%v, err:%v", limitStr, err)
	}

	logger.Debug("get query limit succ, val:%v", limit)
	// 获取1级评论列表
	commentList, count, err := db.GetCommentList(answerId, offset, limit)
	if err != nil {
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		logger.Error("GetCommentList failed, answer_id:%v err:%v", answerId, err)
		return
	}

	var userIdList []int64
	for _, v := range commentList {
		userIdList = append(userIdList[:], v.AuthorId, v.ReplyAuthorId)
	}
	userList, err := db.GetUserInfoList(userIdList)
	if err != nil {
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		logger.Error("GetUserInfoList failed, userIdList:%#v,err%v", userIdList, err)
		return
	}

	userInfoMap := make(map[int64]*common.UserInfo, len(userIdList))
	for _, user := range userList {
		userInfoMap[int64(user.UserId)] = user
	}
	for _, v := range commentList {
		user, ok := userInfoMap[v.AuthorId]
		if ok {
			v.AuthorName = user.Username
		}
		user, ok = userInfoMap[v.ReplyAuthorId]
		if ok {
			v.ReplyAuthorName = user.Username
		}
	}

	apiCommentList := &common.ApiCommentList{}
	apiCommentList.CommentList = commentList
	apiCommentList.Count = count
	util.ResponseSuccess(ctx, apiCommentList)
}

func ReplyListHandle(ctx *gin.Context) {
	commentIdStr, ok := ctx.GetQuery("comment_id")
	commentIdStr = strings.TrimSpace(commentIdStr)
	if !ok || len(commentIdStr) == 0 {
		util.ResponseError(ctx, util.ErrCodeParmeter)
		logger.Error("valid comment id, val:%v, ok:%v", commentIdStr, ok)
		return
	}
	logger.Debug("get query answer_id succ, val:%v", commentIdStr)
	commentId, err := strconv.ParseInt(commentIdStr, 10, 64)
	if err != nil || commentId == 0 {
		util.ResponseError(ctx, util.ErrCodeParmeter)
		logger.Error("valid comment id, val:%v, err:%v", commentId, err)
		return
	}

	//解析offset
	var offset int64
	offsetStr, ok := ctx.GetQuery("offset")
	offsetStr = strings.TrimSpace(offsetStr)
	if !ok || len(offsetStr) == 0 {
		offset = 0
		logger.Error("valid offset, val:%v, err:%v", offsetStr, err)
	}
	offset, err = strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		offset = 0
		logger.Error("valid offset, val:%v, err:%v", offsetStr, err)
	}
	logger.Debug("get query offset succ, val:%v", offset)
	//解析limit
	var limit int64
	limitStr, ok := ctx.GetQuery("limit")
	limitStr = strings.TrimSpace(limitStr)
	if !ok || len(limitStr) == 0 {
		limit = 10
		logger.Error("valid limit, val:%v", limitStr)
	}
	limit, err = strconv.ParseInt(limitStr, 10, 64)
	if err != nil || limit == 0 {
		limit = 10
		logger.Error("valid offset, val:%v, err:%v", limitStr, err)
	}

	logger.Debug("get query limit succ, val:%v", limit)
	// 已经拿到answer_id
	commentList, count, err := db.GetReplyList(commentId, offset, limit)
	if err != nil {
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		logger.Error("GetReplyList failed, answer_id:%v err:%v", commentId, err)
		return
	}

	var userIdList []int64
	for _, v := range commentList {
		userIdList = append(userIdList[:], v.AuthorId, v.ReplyAuthorId)
	}
	userList, err := db.GetUserInfoList(userIdList)
	if err != nil {
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		logger.Error("GetUserInfoList failed, userIdList:%#v,err%v", userIdList, err)
		return
	}

	userInfoMap := make(map[int64]*common.UserInfo, len(userIdList))
	for _, user := range userList {
		userInfoMap[int64(user.UserId)] = user
	}
	for _, v := range commentList {
		user, ok := userInfoMap[v.AuthorId]
		if ok {
			v.AuthorName = user.Username
		}
		user, ok = userInfoMap[v.ReplyAuthorId]
		if ok {
			v.ReplyAuthorName = user.Username
		}
	}

	apiCommentList := &common.ApiCommentList{}
	apiCommentList.CommentList = commentList
	apiCommentList.Count = count
	util.ResponseSuccess(ctx, apiCommentList)
}

func LikeHandle(ctx *gin.Context) {
	var like common.Like
	err := ctx.BindJSON(&like)
	if err != nil {
		util.ResponseError(ctx, util.ErrCodeParmeter)
		logger.Error("like handler failed, err:%v", err)
		return
	}
	if like.Id == 0 || (like.LikeType != common.LikeTypeAnswer && like.LikeType != common.LikeTypeComment) {
		util.ResponseError(ctx, util.ErrCodeParmeter)
		logger.Error("invalid like paramter, data:%#v", like)
		return
	}

	if like.LikeType == common.LikeTypeAnswer {
		err = db.UpdateAnswerLikeCount(like.Id)
	} else {
		err = db.UpdateCommentLikeCount(like.Id)
	}

	if err != nil {
		util.ResponseError(ctx, util.ErrCodeServerBusy)
		logger.Error("update like count failed, err:%v", err)
		return
	}
	util.ResponseSuccess(ctx, nil)
}
