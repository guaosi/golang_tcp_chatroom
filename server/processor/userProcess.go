package processor

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom1/common/message"
	"go_code/chatroom1/common/utils"
	"go_code/chatroom1/server/model"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

func (this *UserProcess) ServeLoginProcess(mess message.Message) (err error) {
	//反序列化获得账号密码
	var loginMes = message.LoginMes{}
	err = json.Unmarshal([]byte(mess.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal error = ", err)
		return
	}
	fmt.Printf("用户id:%d,用户密码:%s 请求登陆\n", loginMes.User.UserId, loginMes.User.UserPwd)
	user, err := model.MyUserDao.Login(loginMes.User.UserId, loginMes.User.UserPwd)
	//开始组装回执信息

	var loginResMes = message.LoginResMes{}
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.ResMes.Code = 404
			loginResMes.ResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.ResMes.Code = 400
			loginResMes.ResMes.Error = err.Error()
		} else {
			loginResMes.ResMes.Code = 500
			loginResMes.ResMes.Error = "服务器内部错误"
		}
		fmt.Println("登陆失败，失败原因:", loginResMes.ResMes.Error)
	} else {
		loginResMes.ResMes.Code = 200
		loginResMes.User = user
		loginResMes.User.UserPwd = ""
		//查看该用户是否已经登陆了
		// old_conn, err := MyUserMgr.GetSimpleUserById(user.UserId)
		// if err == nil {
		// 	//如果没有报错，证明找到了，已经登陆
		// 	old_conn.Close()
		// }
		//将用户信息保存
		MyUserMgr.AddOnlineUser(user.UserId, this.Conn)
		for id, _ := range MyUserMgr.GetAllOnlineUser() {
			loginResMes.Users = append(loginResMes.Users, id)
		}
		this.NotifyOtherOnlineUser(user.UserId, message.NotifyUserMesUp)
		fmt.Printf("用户id: %d 登陆成功\n", loginMes.User.UserId)
	}
	//开始封装
	loginResMesByte, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail error =", err)
		return
	}
	var resMess = message.Message{
		Type: message.LoginResMesType,
		Data: string(loginResMesByte),
	}

	var tf = utils.Transfer{
		Conn: this.Conn,
	}
	tf.WritePkg(&resMess)
	return
}
func (this *UserProcess) ServerRegisterProcess(mess message.Message) (err error) {
	//反序列化获得账号密码
	var registerMes = message.RegisterMes{}
	err = json.Unmarshal([]byte(mess.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal error = ", err)
		return
	}
	fmt.Printf("用户id:%d,用户密码:%s,用户昵称:%s 请求注册\n", registerMes.User.UserId, registerMes.User.UserPwd, registerMes.User.UserName)

	err = model.MyUserDao.Register(registerMes.User.UserId, registerMes.User.UserPwd, registerMes.User.UserName)
	//开始组装回执信息

	var registerResMes = message.RegisterResMes{}
	if err != nil {
		fmt.Println(err)
		if err == model.ERROR_USER_EXISTS {
			registerResMes.ResMes.Code = 403
			registerResMes.ResMes.Error = err.Error()
		} else {
			registerResMes.ResMes.Code = 500
			registerResMes.ResMes.Error = "服务器内部错误"
		}
		fmt.Println("注册失败，失败原因:", registerResMes.ResMes.Error)
	} else {
		registerResMes.ResMes.Code = 200

		fmt.Printf("用户id: %d 注册成功\n", registerMes.User.UserId)
	}
	//开始封装
	registerResMesByte, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal fail error =", err)
		return
	}
	var resMess = message.Message{
		Type: message.RegisterResMesType,
		Data: string(registerResMesByte),
	}

	var tf = utils.Transfer{
		Conn: this.Conn,
	}
	tf.WritePkg(&resMess)
	return
}
func (this *UserProcess) NotifyOtherOnlineUser(userId int, mesType string) {
	for id, conn := range MyUserMgr.GetAllOnlineUser() {
		if id == userId {
			continue
		}
		this.notifyUserMes(userId, conn, mesType)
	}
}

//制作通知上线消息
func (this *UserProcess) notifyUserMes(userId int, conn net.Conn, mesType string) {
	var notifyUserMes = message.NotifyUserMes{
		UserId:  userId,
		MesType: mesType,
	}
	notifyUserMesByte, err := json.Marshal(notifyUserMes)
	if err != nil {
		fmt.Println("json.Marshal error = ", err)
		return
	}
	var mess = message.Message{
		Type: message.NotifyUserMesType,
		Data: string(notifyUserMesByte),
	}
	var tf = utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(&mess)
	if err != nil {
		fmt.Println("通知消息发送失败...error = ", err)
	}
	return
}
