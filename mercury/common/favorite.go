package common

type FavoriteDir struct {
	DirId   int64  `json:"dir_id" db:"dir_id"`
	DirName string `json:"dir_name" db:"dir_name"`
	Count   int32  `json:"count" db:"count"`
	UserId  int64  `json:"user_id" db:"user_id"`
}

type Favorite struct {
	AnswerId int64 `json:"answer_id" db:"answer_id"`
	UserId   int64 `json:"user_id" db:"user_id"`
	DirId    int64 `json:"dir_id" db:"dir_id"`
}
