package common

const (
	UserSexMan   = 1
	UserSexWomen = 2
)

type UserInfo struct {
	UserId   uint64 `json:"user_id" db:"user_id"`
	NickName string `json:"nickname" db:"nickname"`
	Username string `json:"user" db:"username"`
	Password string `json:"password" db:"password"`
	Email    string `json:"email" db:"email"`
	Sex      int    `json:"sex" db:"sex"`
}
