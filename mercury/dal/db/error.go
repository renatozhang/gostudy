package db

import "errors"

var (
	ErrUserExists        = errors.New("username is exist")
	ErrUserNotExists     = errors.New("username not is exist")
	ErrUserPasswordWrong = errors.New("username or password not right")
)
