package common

import "time"

type Comment struct {
	CommentId       int64     `json:"comment_id" db:"comment_id"`
	Content         string    `json:"content" db:"content"`
	AuthorId        int64     `json:"author_id" db:"author_id"`
	LikeCount       int       `json:"like_count" db:"like_count"`
	CommentCount    int       `json:"comment_count" db:"comment_count"`
	CreateTime      time.Time `json:"create_time" db:"create_time"`
	ParentId        int64     `json:"parent_id" db:"parent_id"`
	QuestionId      int64     `json:"question_id" db:"question_id"`
	ReplyAuthorId   int64     `json:"reply_author_id" db:"reply_author_id"`
	ReplyCommentId  int64     `json:"reply_comment_id" db:"reply_comment_id"`
	AuthorName      string    `json:"author_name"`
	ReplyAuthorName string    `json:"reply_author_name"`
}

type ApiCommentList struct {
	CommentList []*Comment `json:"comment_list"`
	Count       int64      `json:"count"`
}
