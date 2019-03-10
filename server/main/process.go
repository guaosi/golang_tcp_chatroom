package main

import (
	"fmt"
	"go_code/chatroom1/common/message"
	"go_code/chatroom1/common/utils"
	"go_code/chatroom1/server/processor"
	"io"
	"net"
)

type MainProcess struct {
	Conn net.Conn
}

func (this *MainProcess) ServeProcess() (err error) {
	var tr = utils.Transfer{
		Conn: this.Conn,
	}
	for {
		mess, err1 := tr.ReadPkg()
		err = err1
		if err != nil {

			if err == io.EOF {
				//这里编写下线通知

				userId, err := processor.MyUserMgr.GetIdByConn(this.Conn)
				if err != nil {
					return err
				}
				processor.MyUserMgr.DeleteOnlineUser(userId)
				var userprocess = processor.UserProcess{
					Conn: this.Conn,
				}

				userprocess.NotifyOtherOnlineUser(userId, message.NotifyUserMesDown)

			}
			return
		}
		err = this.CategoryProcess(mess)
		if err != nil {
			fmt.Println("回执包发送失败")
		}
	}
	return
}
func (this *MainProcess) CategoryProcess(mess message.Message) (err error) {
	var userprocess = processor.UserProcess{
		Conn: this.Conn,
	}
	var smsProcess = processor.SmsProcess{}
	switch mess.Type {
	case message.LoginMesType:
		userprocess.ServeLoginProcess(mess)
		return
	case message.RegisterMesType:
		userprocess.ServerRegisterProcess(mess)
		return
	case message.SmsMesType:
		smsProcess.SendToAll(mess)
	case message.SmsMesSimpleMesType:
		smsProcess.SendMesToSimple(mess, this.Conn)
	default:
		fmt.Println("数据包类型不正确")
		return
	}
	return
}
