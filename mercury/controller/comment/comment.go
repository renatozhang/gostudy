package comment

import (
	"html"

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
