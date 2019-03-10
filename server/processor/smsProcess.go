package processor

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom1/common/message"
	"go_code/chatroom1/common/utils"
	"net"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendToAll(mess message.Message) (err error) {
	var smsMes = message.SmsMes{}
	err = json.Unmarshal([]byte(mess.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal error = ", err)
		return
	}
	for id, conn := range MyUserMgr.GetAllOnlineUser() {
		if id == smsMes.User.UserId {
			continue
		}
		err = this.SendToSimple(mess, conn)
		if err != nil {
			fmt.Printf("发送消息给:%d 失败，原因:%v", id, err)
		}
	}
	return
}
func (this *SmsProcess) SendToSimple(mess message.Message, conn net.Conn) (err error) {
	//再把信息转回json，进行发送
	var tf = &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(&mess)
	return
}
