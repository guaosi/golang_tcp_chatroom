package processor

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom1/common/message"
)

func OutSmsMessage(msg message.Message) {
	//转成结构体smsMes
	var smsMes = message.SmsMes{}
	err := json.Unmarshal([]byte(msg.Data), &smsMes)
	if err != nil {
		fmt.Println("接收他人消息失败...,原因:", err)
		return
	}
	info := fmt.Sprintf("用户 %d 群发消息:%s", smsMes.User.UserId, smsMes.Content)
	fmt.Println(info)
	return
}
