package processor

import (
	"encoding/json"
	"errors"
	"fmt"
	"go_code/chatroom1/client/model"
	"go_code/chatroom1/common/message"
	"go_code/chatroom1/common/utils"
	"net"
)

type UserProcess struct {
}

func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("连接服务器失败,error = ", err)
		return
	}
	defer conn.Close()
	//将接收到的数据放到结构体中去
	var loginmes = message.LoginMes{
		User: message.User{
			UserId:  userId,
			UserPwd: userPwd,
		},
	}
	loginmesByte, err := json.Marshal(loginmes)
	if err != nil {
		fmt.Println("json.Marshal error = ", err)
		return
	}
	var mess = message.Message{
		Type: message.LoginMesType,
		Data: string(loginmesByte),
	}
	var tf = utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(&mess)

	//开始接收服务端发送过来的消息
	messRes, err := tf.ReadPkg()
	if err != nil {
		return
	}
	if messRes.Type != message.LoginResMesType {
		fmt.Println("回执包不正确")
		errors.New("回执包不正确")
		return
	}
	var loginResMes = message.LoginResMes{}
	err = json.Unmarshal([]byte(messRes.Data), &loginResMes)
	if err != nil {
		fmt.Println("json.Unmarshal error = ", err)
		return
	}
	if loginResMes.ResMes.Code != 200 {
		fmt.Println("登陆失败")
		return
	}
	model.CurrentUser.User = loginResMes.User
	model.CurrentUser.Conn = conn
	initOnlineUser(loginResMes)
	go serverProcess(conn)
	for {
		showMenu()
	}
	return
}
func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("连接服务器失败,error = ", err)
		return
	}
	defer conn.Close()
	//将接收到的数据放到结构体中去
	var registermes = message.RegisterMes{
		User: message.User{
			UserId:   userId,
			UserPwd:  userPwd,
			UserName: userName,
		},
	}
	registermesByte, err := json.Marshal(registermes)
	if err != nil {
		fmt.Println("json.Marshal error = ", err)
		return
	}
	var mess = message.Message{
		Type: message.RegisterMesType,
		Data: string(registermesByte),
	}
	var tf = utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(&mess)

	//开始接收服务端发送过来的消息
	messRes, err := tf.ReadPkg()
	if err != nil {
		return
	}
	if messRes.Type != message.RegisterResMesType {
		fmt.Println("回执包不正确")
		errors.New("回执包不正确")
		return
	}
	var registerResMes = message.RegisterResMes{}
	err = json.Unmarshal([]byte(messRes.Data), &registerResMes)
	if err != nil {
		fmt.Println("json.Unmarshal error = ", err)
		return
	}
	if registerResMes.ResMes.Code != 200 {
		fmt.Println("注册失败,error = ", registerResMes.Error)
		return
	}
	fmt.Println("注册成功")
	return
}
