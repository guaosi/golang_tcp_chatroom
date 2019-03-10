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
