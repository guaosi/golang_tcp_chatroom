package main

import (
	"fmt"
	"go_code/chatroom1/client/processor"
	"os"
)

func main() {
	var key int
	for {
		fmt.Println("--------------------欢迎登陆多人聊天系统--------------")
		fmt.Println("\t\t\t 1 登陆聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		//scanf一定后面要加上\n,否则它会认为回车也是一个输入
		fmt.Print("请选择(1-3): ")
		fmt.Scanf("%d\n", &key)
		var userName string
		var userPwd string
		var userId int
		var userProcess = processor.UserProcess{}
		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			fmt.Println("请输入登陆的用户Id:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入密码:")
			fmt.Scanf("%s\n", &userPwd)
			userProcess.Login(userId, userPwd)
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入注册的用户Id:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入密码:")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入要注册的昵称:")
			fmt.Scanf("%s\n", &userName)
			userProcess.Register(userId, userPwd, userName)
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("输入错误，请重新输入")
		}
	}

}
