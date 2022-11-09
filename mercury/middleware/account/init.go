package account

import "github.com/renatozhang/gostudy/mercury/session"

// 初始化一个sessionMgr
func InitSession(provider string, addr string, options ...string) (err error) {
	return session.Init(provider, addr, options...)
}
