package processor

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom1/client/model"
	"go_code/chatroom1/common/message"
	"go_code/chatroom1/common/utils"
)

type SmsProcess struct {
}

func (this *SmsProcess) sendToAll(content string) (err error) {
	var smsMes = message.SmsMes{
		User:    model.CurrentUser.User,
		Content: content,
	}
	smsMesByte, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("json.Marshal error=", err)
		return
	}
	var mess = message.Message{
		Type: message.SmsMesType,
		Data: string(smsMesByte),
	}
	var tf = utils.Transfer{
		Conn: model.CurrentUser.Conn,
	}
	err = tf.WritePkg(&mess)
	if err != nil {
		fmt.Println("发送数据失败:", err)
	}
	return
}
func (this *SmsProcess) sendToSimple(content string, userId int) (err error) {
	if model.CurrentUser.User.UserId == userId {
		fmt.Println("不允许给自己发送聊天消息")
		return
	}
	var smsMesSimpleMes = message.SmsMesSimpleMes{
		SmsMes:   message.SmsMes{User: model.CurrentUser.User, Content: content},
		ToUserId: userId,
	}
	smsMesSimpleMesByte, err := json.Marshal(smsMesSimpleMes)
	if err != nil {
		fmt.Println("json.Marshal error=", err)
		return
	}
	var mess = message.Message{
		Type: message.SmsMesSimpleMesType,
		Data: string(smsMesSimpleMesByte),
	}
	var tf = utils.Transfer{
		Conn: model.CurrentUser.Conn,
	}
	err = tf.WritePkg(&mess)
	if err != nil {
		fmt.Println("发送数据失败:", err)
	}
	return

}
