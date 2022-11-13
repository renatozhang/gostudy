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
