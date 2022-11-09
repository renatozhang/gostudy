package session

import (
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
	uuid "github.com/satori/go.uuid"
)

type RedisSessionMgr struct {
	addr       string
	password   string
	pool       *redis.Pool
	rwlock     sync.RWMutex
	sessionMap map[string]Session
}

func NewRedisSessionMgr() SessionMgr {
	sr := &RedisSessionMgr{
		sessionMap: make(map[string]Session, 1024),
	}
	return sr
}

//初始化一个pool
func newPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     64,
		MaxActive:   1000,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			/*
				if _,err:=c.Do("AUTH",password){
					c.Close()
					return nil,err
				}
			*/
			return c, err

		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func (r *RedisSessionMgr) Init(addr string, options ...string) (err error) {
	if len(options) > 0 {
		r.password = options[0]
	}
	r.pool = newPool(addr, r.password)
	r.addr = addr
	return
}

func (r *RedisSessionMgr) CreateSession() (session Session, err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	id, err := uuid.NewV4()
	if err != nil {
		return
	}
	sessionId := id.String()
	session = NewRedisSession(sessionId, r.pool)
	r.sessionMap[sessionId] = session

	return
}

func (r *RedisSessionMgr) Get(sessionId string) (session Session, err error) {
	r.rwlock.RLock()
	defer r.rwlock.RUnlock()
	session, ok := r.sessionMap[sessionId]
	if !ok {
		err = ErrSessionNotExist
		return
	}

	return
}
