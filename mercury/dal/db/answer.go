package db

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/renatozhang/gostudy/mercury/common"
	"github.com/renatozhang/gostudy/mercury/logger"
)

func GetAnswerIdList(questionId int64, offset, limit int64) (answerIdList []int64, err error) {
	sqlstr := `select
					answer_id
				from
					question_answer_rel
				where question_id=?
				limit ?,?`
	err = DB.Select(&answerIdList, sqlstr, questionId, offset, limit)
	if err != nil {
		logger.Error("get answer list failed, err:%v", err)
		return
	}
	return
}

func MGetAnswer(answerIdList []int64) (answerList []*common.Answer, err error) {
	sqlstr := `select 
						answer_id,content,comment_count,
						voteup_count,author_id,status,
						can_comment,create_time,update_time
					from 
						answer
					where answer_id in (?)`
	insqlStr, params, err := sqlx.In(sqlstr, answerIdList)
	if err != nil {
		logger.Error("sqlx.in faile, sqlstr:%v, err:%v", sqlstr, err)
		return
	}
	err = DB.Select(&answerList, insqlStr, params...)
	if err == sql.ErrNoRows {
		err = nil
		return
	}
	if err != nil {
		logger.Error("MGetAnswer failed, sqlstr:%v, answer_ids:%v, err:%v", sqlstr, answerIdList, err)
		return
	}
	return
}

func GetAnswerCount(questionId int64) (answerCount int32, err error) {
	sqlstr := `select
					count(answer_id)
				from
					question_answer_rel
				where question_id=?`
	err = DB.Get(&answerCount, sqlstr, questionId)
	if err != nil {
		logger.Error("GetAnswerCount failed, err:%v", err)
		return
	}
	return
}

func UpdateAnswerLikeCount(answerId int64) (err error) {
	sqlstr := `update
					answer
				set
				voteup_count=voteup_count+1
				where answer_id=?`
	_, err = DB.Exec(sqlstr, answerId)
	if err != nil {
		logger.Error("UpdateAnswerCommentCount failed, err:%v", err)
		return
	}
	return
}
