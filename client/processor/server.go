package processor

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom1/client/model"
	"go_code/chatroom1/common/message"
	"go_code/chatroom1/common/utils"
	"net"
	"os"
)

func serverProcess(conn net.Conn) {
	var tf = utils.Transfer{
		Conn: conn,
	}
	for {
		mess, err := tf.ReadPkg()
		if err != nil {
			return
		}
		switch mess.Type {

		case message.NotifyUserMesType:
			var notifyUserMes = message.NotifyUserMes{}
			err := json.Unmarshal([]byte(mess.Data), &notifyUserMes)
			if err != nil {
				fmt.Println("json.Unmarshal error = ", err)
				continue
			}
			if notifyUserMes.MesType == message.NotifyUserMesUp {
				fmt.Printf("用户 %d 已上线 \n", notifyUserMes.UserId)
				addOnlineUser(notifyUserMes.UserId)
			} else if notifyUserMes.MesType == message.NotifyUserMesDown {
				fmt.Printf("用户 %d 已下线 \n", notifyUserMes.UserId)
				deleteOnlineUser(notifyUserMes.UserId)
			}
			continue
		case message.SmsMesType:
			OutSmsMessage(mess)
		case message.SmsMesSimpleMesType:
			OutSmsSimpleMessage(mess)
		case message.SmsMesSimpleResMesType:
			ResSmsSimpleMessage(mess)
		default:
			fmt.Println("类型不正确")
			continue
		}
	}

}

func showMenu() {
	var key int
	var content string
	fmt.Printf("--------------------恭喜%s登陆成功--------------\n", model.CurrentUser.User.UserName)
	fmt.Println("\t\t\t 1 显示在线用户列表")
	fmt.Println("\t\t\t 2 群发消息")
	fmt.Println("\t\t\t 3 私聊消息")
	fmt.Println("\t\t\t 4 信息列表")
	fmt.Println("\t\t\t 5 退出系统")
	fmt.Print("请选择(1-4): ")
	fmt.Scanf("%d\n", &key)
	var userId int
	var smsprocess = SmsProcess{}
	switch key {
	case 1:
		outOnlineUser()
	case 2:
		fmt.Println("请输入你要发送给全体的消息:")
		fmt.Scanf("%s\n", &content)
		smsprocess.sendToAll(content)
	case 3:
		fmt.Println("请输入你要发送的用户ID:")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("请输入你要发送的内容:")
		fmt.Scanf("%s\n", &content)
		smsprocess.sendToSimple(content, userId)
	case 4:
		fmt.Println("信息列表")
	case 5:
		fmt.Println("你选择了退出系统...")
		os.Exit(0)
	default:
		fmt.Println("输入有误，请重新输入")
	}
}
