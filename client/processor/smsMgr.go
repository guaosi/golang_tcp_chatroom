package processor

import (
	"encoding/json"
	"errors"
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
func OutSmsSimpleMessage(msg message.Message) {
	//转成结构体smsMes
	var smsMesSimpleMes = message.SmsMesSimpleMes{}
	err := json.Unmarshal([]byte(msg.Data), &smsMesSimpleMes)
	if err != nil {
		fmt.Println("接收他人消息失败...,原因:", err)
		return
	}
	info := fmt.Sprintf("用户 %d 跟你说:%s", smsMesSimpleMes.SmsMes.User.UserId, smsMesSimpleMes.SmsMes.Content)
	fmt.Println(info)
	return
}
func ResSmsSimpleMessage(messRes message.Message) (err error) {
	if messRes.Type != message.SmsMesSimpleResMesType {
		fmt.Println("回执包不正确")
		err = errors.New("回执包不正确")
		return
	}
	var smsMesSimpleResMes = message.SmsMesSimpleResMes{}
	err = json.Unmarshal([]byte(messRes.Data), &smsMesSimpleResMes)
	if err != nil {
		fmt.Println("json.Unmarshal error = ", err)
		return
	}
	if smsMesSimpleResMes.ResMes.Code != 200 {
		fmt.Println("发送失败,error = ", smsMesSimpleResMes.ResMes.Error)
		return
	}
	fmt.Println("发送成功")
	return
}
