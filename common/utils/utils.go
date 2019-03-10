package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/chatroom1/common/message"
	"io"
	"net"
)

type Transfer struct {
	Conn net.Conn
	buf  [8096]byte
}

//发送传入的message.Message
func (this *Transfer) WritePkg(mess *message.Message) (err error) {
	//1.先计算传入的长度,并且发送长度包

	//1.1 json化
	messByte, err := json.Marshal(*mess)
	if err != nil {
		fmt.Println("json.Marshal fail error =", err)
		return
	}
	//1.2 计算长度,转化为byte
	//因为是Uint32，表示4个字节。4个字节需要4个byte，一个字节8位，4个就是2的32次方
	//最大可以表示4G的数据长度
	binary.BigEndian.PutUint32(this.buf[:4], uint32(len(messByte)))
	//1.4 发送长度
	_, err = this.Conn.Write(this.buf[:4])
	if err != nil {
		fmt.Println("发送数据长度失败,error = ", err)
		return
	}
	//2.再发送真正的内容
	_, err = this.Conn.Write(messByte)
	if err != nil {
		fmt.Println("发送内容失败,error = ", err)
	}
	return
}

//用于将接收包进行反json化
func (this *Transfer) ReadPkg() (mess message.Message, err error) {
	n, err := this.Conn.Read(this.buf[:4])
	if err != nil || n != 4 {
		if err == io.EOF {
			fmt.Println("客户端异常退出，关闭连接,error = ", err)
		} else {
			fmt.Println("读取数据错误,error = ", err)
		}
		return
	}
	//将接收到的byte再转回数字
	pkg_len := binary.BigEndian.Uint32(this.buf[:4])
	n, err = this.Conn.Read(this.buf[:pkg_len])
	if err != nil || n != int(pkg_len) {
		if err == io.EOF {
			fmt.Println("客户端异常退出，关闭连接,error = ", err)
		} else {
			fmt.Println("读取数据错误,error = ", err)
		}
		return
	}
	err = json.Unmarshal(this.buf[:pkg_len], &mess)
	if err != nil {
		fmt.Println("json.Unmarshal error = ", err)
	}
	return
}
