package model

import (
	"go_code/chatroom1/common/message"
	"net"
)

var CurrentUser = CurUser{}

type CurUser struct {
	User message.User
	Conn net.Conn
}
