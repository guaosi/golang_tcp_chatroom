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
func (this *SmsProcess) SendMesToSimple(mess message.Message, send_conn net.Conn) (err error) {
	//先获取到userId
	var smsMesSimpleMes = message.SmsMesSimpleMes{}
	err = json.Unmarshal([]byte(mess.Data), &smsMesSimpleMes)
	if err != nil {
		fmt.Println("json.Unmarshal error = ", err)
		return
	}

	var smsMesSimpleResMes = message.SmsMesSimpleResMes{}

	conn, err := MyUserMgr.GetSimpleUserById(smsMesSimpleMes.ToUserId)
	if err != nil {
		smsMesSimpleResMes.ResMes.Code = 404
		smsMesSimpleResMes.ResMes.Error = err.Error()
	} else {
		err = this.SendToSimple(mess, conn)
		if err != nil {
			fmt.Printf("用户 %d 发送给 %d 的消息发送失败", smsMesSimpleMes.SmsMes.User.UserId, smsMesSimpleMes.ToUserId)
			smsMesSimpleResMes.ResMes.Code = 500
			smsMesSimpleResMes.ResMes.Error = err.Error()
		} else {
			smsMesSimpleResMes.ResMes.Code = 200
		}
		//发送回执消息
	}
	//开始封装
	smsMesSimpleResMesByte, err := json.Marshal(smsMesSimpleResMes)
	if err != nil {
		fmt.Println("json.Marshal fail error =", err)
		return
	}
	var resMess = message.Message{
		Type: message.SmsMesSimpleResMesType,
		Data: string(smsMesSimpleResMesByte),
	}
	var tf = utils.Transfer{
		Conn: send_conn,
	}
	tf.WritePkg(&resMess)
	if err != nil {
		fmt.Println("发送回执数据失败:", err)
	} else {
		fmt.Println("发送回执数据成功")
	}
	return
}
