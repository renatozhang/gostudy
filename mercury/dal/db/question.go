package db

import (
	"github.com/renatozhang/gostudy/mercury/common"
	"github.com/renatozhang/gostudy/mercury/logger"
)

func CreateQuestion(question *common.Question) (err error) {
	sqlstr := `insert into question(
			question_id, caption, content, author_id, category_id) 
			values(?,?,?,?,?)`
	_, err = DB.Exec(sqlstr, question.QuestionId, question.Caption, question.Content, question.AuthorId, question.CategoryId)
	if err != nil {
		logger.Error("creeate question failed,question:%v, err:%v", question, err)
		return
	}
	return
}

/*
	QuestionId int64  `json:"question_id" db:"question_id"`
	Caption    string `json:"caption" db:"caption"`
	Content    string `json:"content" db:"content"`
	AuthorId   int64  `json:"author_id" db:"author_id"`
	CategoryId int64  `json:"category_id" db:"category_id"`
	Status     int32  `json:"status" db:"status"`
*/
func GetQuestionList(categoryId int64) (questionList []*common.Question, err error) {
	sqlstr := `select
			   		question_id,caption,content,author_id,category_id,create_time
			   from question
			   where category_id=?`
	err = DB.Select(&questionList, sqlstr, categoryId)
	if err != nil {
		logger.Error("get question list failed, err:%v", err)
		return
	}
	return
}
