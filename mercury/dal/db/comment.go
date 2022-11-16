package db

import (
	"github.com/renatozhang/gostudy/mercury/common"
	"github.com/renatozhang/gostudy/mercury/logger"
)

func CreateReplyComment(comment *common.Comment) (err error) {
	tx, err := DB.Begin()
	if err != nil {
		logger.Error("create reply comment failed, comment:%#v, err:%v", comment, err)
		return
	}

	defer func() {
		if err != nil {
			logger.Debug("tx roolback, err:%v", err)
			tx.Rollback()
			return
		}
	}()

	// 根据ReplayCommentId查询对应的authorid
	var replyAuthorId int64
	sqlstr := "select author_id from comment where comment_id=?"
	err = DB.Get(&replyAuthorId, sqlstr, comment.ReplyCommentId)
	if err != nil {
		logger.Error("select author id failed, err:%v, cid:%v", err, comment.ReplyCommentId)
		return
	}
	sqlstr = `insert into comment(
					comment_id,content,author_id,reply_comment_id)
			values(?,?,?,?)`
	_, err = tx.Exec(sqlstr, comment.CommentId, comment.Content, comment.AuthorId, comment.ReplyCommentId)
	if err != nil {
		logger.Error("create reply comment failed, comment:%#v, err:%v", comment, err)
		return
	}

	sqlstr = `insert into comment_rel(
					comment_id,parent_id,level,question_id,reply_author_id) 
				values(?,?,?,?,?)`
	_, err = tx.Exec(sqlstr, comment.CommentId, comment.ParentId, 2, comment.QuestionId, replyAuthorId)
	if err != nil {
		logger.Error("create reply comment_rel failed, comment:%#v, err:%v", comment, err)
		return
	}

	sqlstr = "update comment set comment_count=comment_count+1 where comment_id=?"
	_, err = tx.Exec(sqlstr, comment.ParentId)
	if err != nil {
		logger.Error("update comment failed, comment:%#v, err:%v", comment, err)
		return
	}
	err = tx.Commit()
	return
}

func CreatePostComment(comment *common.Comment) (err error) {
	tx, err := DB.Begin()
	if err != nil {
		logger.Error("create post comment failed, comment:%#v, err:%v", comment, err)
		return
	}
	defer func() {
		if err != nil {
			logger.Debug("tx roolback, err:%v", err)
			tx.Rollback()
			return
		}
	}()
	sqlstr := `insert into comment(
					comment_id,content,author_id) 
				values(?,?,?)`
	_, err = tx.Exec(sqlstr, comment.CommentId, comment.Content, comment.AuthorId)
	if err != nil {
		logger.Error("insert comment failed, comment:%#v, err:%v", comment, err)
		return
	}

	sqlstr = `insert into comment_rel(
					comment_id,parent_id,level,question_id,reply_author_id) 
				values(?,?,?,?,?)`
	_, err = tx.Exec(sqlstr, comment.CommentId, comment.ParentId, 1, comment.QuestionId, comment.ReplyAuthorId)
	if err != nil {
		logger.Error("insert comment_rel failed, comment:%#v, err:%v", comment, err)
		return
	}
	err = tx.Commit()
	return
}
