package session

import "errors"

var (
	ErrSessionNotExist = errors.New("session not exists")
	ErrKeyNotExist     = errors.New("key not exists in session")
)
