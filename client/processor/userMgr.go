package processor

import (
	"fmt"
	"go_code/chatroom1/client/model"
	"go_code/chatroom1/common/message"
)

var OnlineUser map[int]*message.User = make(map[int]*message.User, 1024)

func initOnlineUser(loginResMes message.LoginResMes) {
	for _, id := range loginResMes.Users {
		if id == model.CurrentUser.User.UserId {
			continue
		}
		addOnlineUser(id)
	}
}

// func addOnlineUser()
func outOnlineUser() {
	fmt.Println("当前在线用户列表:")
	for _, v := range OnlineUser {
		fmt.Println("用户id: ", v.UserId)
	}
	return
}
func addOnlineUser(userId int) {
	user := &message.User{
		UserId: userId,
	}
	OnlineUser[userId] = user
}
func deleteOnlineUser(userId int) {
	delete(OnlineUser, userId)
}
