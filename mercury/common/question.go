package common

type Question struct {
	QuestionId int64  `json:"question_id" db:"question_id"`
	Caption    string `json:"caption" db:"caption"`
}
