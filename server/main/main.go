package main

import (
	"fmt"
	"go_code/chatroom1/server/model"
	"go_code/chatroom1/server/processor"
	"net"
	"time"
)

func init() {
	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()
	processor.InitUserMgr()
}
func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}
func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net connect error = ", err)
	}
	defer listen.Close()
	fmt.Println("服务器开始监听...")
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("客户端连接上服务器出错,error = ", err)
			continue
		}
		go process(conn)
	}
}
func process(conn net.Conn) {
	fmt.Println("有客户端连接上了,IP:", conn.RemoteAddr().String())
	var mainProcess = MainProcess{
		Conn: conn,
	}
	mainProcess.ServeProcess()

}
