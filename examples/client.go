package main

import (
	"fmt"
	"github.com/ilovesusu/Supreme/sunet"
	"io"
	"net"
	"time"
)

func main() {
	//1 连接服务器
	conn, err := net.Dial("tcp4", "127.0.0.1:12727")
	if err != nil {
		fmt.Println("ss")
	}
	//2 写数据
	for {
		buf := []byte("woaini")
		msgPackage := sunet.NewMsgPackage(1, buf)
		dataPack := sunet.NewDataPack()
		binaryMsg, err := dataPack.Pack(msgPackage)
		if err != nil {
			fmt.Println("message pack error:", err)
		}
		_, err = conn.Write(binaryMsg)
		if err != nil {
			fmt.Println("write error err ", err)
			return
		}

		time.Sleep(2 * time.Second)

		//先读出流中的head部分
		headData := make([]byte, dataPack.GetHeadLen())
		_, err = io.ReadFull(conn, headData) //ReadFull 会把msg填充满为止
		if err != nil {
			fmt.Println("read head error")
			break
		}
		//将headData字节流 拆包到msg中,返回值是接口类型
		imsg, err := dataPack.Unpack(headData)
		if err != nil {
			fmt.Println("server unpack err:", err)
			return
		}

		if imsg.GetDataLen() > 0 {
			//接口对象转换message对象
			message := imsg.(*sunet.Message)
			message.Data = make([]byte, imsg.GetDataLen())

			//根据dataLen从io中读取字节流
			_, err := io.ReadFull(conn, message.Data)
			if err != nil {
				fmt.Println("server unpack data err:", err)
				return
			}
			fmt.Println("==> Recv Msg: ID=", message.Id, ", len=", message.DataLen, ", data=", string(message.Data))
		}
	}
}
