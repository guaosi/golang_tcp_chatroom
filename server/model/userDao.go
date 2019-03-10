package model

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom1/common/message"

	"github.com/garyburd/redigo/redis"
)

var MyUserDao *userDao

type userDao struct {
	pool *redis.Pool
}

func NewUserDao(pool *redis.Pool) *userDao {
	return &userDao{
		pool: pool,
	}
}
func (this *userDao) getUserById(conn redis.Conn, userId int) (user message.User, err error) {
	res, err := redis.String(conn.Do("hget", "users", userId))
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	//反序列化
	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		fmt.Println("json.Unmarshal error = ", err)
	}
	return
}

func (this *userDao) Login(userId int, userPwd string) (user message.User, err error) {
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
	}
	return
}

func (this *userDao) Register(userId int, userPwd string, userName string) (err error) {
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.getUserById(conn, userId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	} else {
		if err != ERROR_USER_NOTEXISTS {
			fmt.Println("内部错误")
		} else {
			err = nil
		}
	}
	//注册新用户
	var user = message.User{
		UserId:   userId,
		UserPwd:  userPwd,
		UserName: userName,
	}
	userByte, err := json.Marshal(user)
	if err != nil {
		fmt.Println("json.Marshal error =", err)
	}
	_, err = conn.Do("hset", "users", userId, string(userByte))
	return
}
