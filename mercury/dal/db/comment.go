package db

import (
	"github.com/jmoiron/sqlx"
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
	sqlstr = "update answer set comment_count=comment_count+1 where answer_id=?"
	_, err = tx.Exec(sqlstr, comment.QuestionId)
	if err != nil {
		logger.Error("update answer failed, comment:%#v, err:%v", comment, err)
		return
	}
	err = tx.Commit()
	return
}

func GetCommentList(answerId int64, offset, limit int64) (commentList []*common.Comment, count int64, err error) {
	var commentIdList []int64
	sqlstr := "select comment_id from comment_rel where question_id=? and level=1 limit ?,?"
	logger.Debug("%v, %v, %v", answerId, offset, limit)
	err = DB.Select(&commentIdList, sqlstr, answerId, offset, limit)
	if err != nil {
		logger.Error("query comment list failed, answer_id:%#v, offset:%v, limit:%v err:%v", answerId, err, offset, limit)
		return
	}
	sqlstr = `select
					comment_id,content,author_id,like_count,comment_count,create_time
				from
					comment
				where comment_id in (?)`
	var tempList []interface{}
	for _, val := range commentIdList {
		tempList = append(tempList, val)
	}
	sqlstr, paramList, err := sqlx.In(sqlstr, tempList)
	if err != nil {
		logger.Error("sqlx in failed, answer_id:%#v, err:%v", answerId, err)
		return
	}
	err = DB.Select(&commentList, sqlstr, paramList...)
	if err != nil {
		logger.Error("sql.select failed, answer_id:%v, err:%v", answerId, err)
		return
	}
	// 查询总的记录条数
	sqlstr = "select count(comment_id) from comment_rel where question_id=? and level=1"
	err = DB.Get(&count, sqlstr, answerId)
	if err != nil {
		logger.Error("DB.Get failed, answer_id:%v, err:%v", answerId, err)
		return
	}

	return
}

func GetReplyList(commentId int64, offset, limit int64) (commentList []*common.Comment, count int64, err error) {
	var commentIdList []int64
	sqlstr := "select comment_id from comment_rel where parent_id=? and level=2 limit ?,?"
	logger.Debug("%v, %v, %v", commentId, offset, limit)
	err = DB.Select(&commentIdList, sqlstr, commentId, offset, limit)
	if err != nil {
		logger.Error("query comment list failed, answer_id:%#v, offset:%v, limit:%v err:%v", commentId, err, offset, limit)
		return
	}
	sqlstr = `select
					comment_id,content,author_id,like_count,comment_count,create_time
				from
					comment
				where comment_id in (?)`
	var tempList []interface{}
	for _, val := range commentIdList {
		tempList = append(tempList, val)
	}
	sqlstr, paramList, err := sqlx.In(sqlstr, tempList)
	if err != nil {
		logger.Error("sqlx in failed, comment_id:%#v, err:%v", commentId, err)
		return
	}
	err = DB.Select(&commentList, sqlstr, paramList...)
	if err != nil {
		logger.Error("sql.select failed, answer_id:%v, err:%v", commentId, err)
		return
	}
	// 查询总的记录条数
	sqlstr = "select count(comment_id) from comment_rel where parent_id=? and level=2"
	err = DB.Get(&count, sqlstr, commentId)
	if err != nil {
		logger.Error("DB.Get failed, answer_id:%v, err:%v", commentId, err)
		return
	}

	return
}

func UpdateCommentLikeCount(commentId int64) (err error) {
	sqlstr := `update
					comment
				set
				like_count=like_count+1
				where comment_id=?`
	_, err = DB.Exec(sqlstr, commentId)
	if err != nil {
		logger.Error("UpdateCommentLikeCount failed, err:%v", err)
		return
	}
	return
}
