package processor

import (
	"errors"
	"fmt"
	"net"
)

var MyUserMgr *userMgr

type userMgr struct {
	OnliuneUser map[int]net.Conn
}

func NewUserMgr() *userMgr {
	return &userMgr{
		OnliuneUser: make(map[int]net.Conn, 1024),
	}
}
func InitUserMgr() {
	MyUserMgr = NewUserMgr()
}
func (this *userMgr) AddOnlineUser(userId int, conn net.Conn) {
	this.OnliuneUser[userId] = conn
}
func (this *userMgr) DeleteOnlineUser(userId int) {
	delete(this.OnliuneUser, userId)
}
func (this *userMgr) GetAllOnlineUser() map[int]net.Conn {
	return this.OnliuneUser
}
func (this *userMgr) GetSimpleUserById(userId int) (conn net.Conn, err error) {
	conn, ok := this.OnliuneUser[userId]
	if !ok {
		fmt.Printf("用户 %d 不存在\n", userId)
		err = fmt.Errorf("用户 %d 不存在\n", userId)
	}
	return
}
func (this *userMgr) GetIdByConn(conn net.Conn) (userId int, err error) {
	for id, userConn := range this.OnliuneUser {
		if userConn == conn {
			userId = id
			return
		}
	}
	err = errors.New("没有找到该链接")
	return
}
