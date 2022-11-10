package db

import (
	"database/sql"

	"github.com/renatozhang/gostudy/mercury/common"
	"github.com/renatozhang/gostudy/mercury/util"
)

const (
	PasswordSalt = "tLcZlok88secbtcU5tJoJ&KQdUu9$&vR"
)

func Register(user *common.UserInfo) (err error) {
	var count int
	sqlstr := "select count(user_id) from user where username=?"
	err = DB.Get(&count, sqlstr, user.Username)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	if count > 0 {
		err = ErrUserExists
		return
	}
	passwd := user.Password + PasswordSalt
	dbPassword := util.Md5([]byte(passwd))
	sqlstr = "insert into user(username,password,email,user_id,nickname,sex)values(?,?,?,?,?,?)"
	_, err = DB.Exec(sqlstr, user.Username, dbPassword, user.Email, user.UserId, user.NickName, user.Sex)
	return
}

func Login(user *common.UserInfo) (err error) {
	originPassword := user.Password
	sqlstr := "select username,password,user_id from user where username=?"
	err = DB.Get(user, sqlstr, user.Username)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	if err == sql.ErrNoRows {
		err = ErrUserNotExists
		return
	}
	passwd := originPassword + PasswordSalt
	originPasswordSalt := util.Md5([]byte(passwd))
	if originPasswordSalt != user.Password {
		err = ErrUserPasswordWrong
		return
	}
	return
}
