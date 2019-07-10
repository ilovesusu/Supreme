package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	//1 连接服务器
	conn, err := net.Dial("tcp4", "127.0.0.1:2727")
	if err != nil {
		fmt.Println("ss")
	}
	//2 写数据
	for {
		buf := []byte("woaini")
		_, err := conn.Write(buf)
		if err != nil {
			fmt.Println("wirte conn err", err)
		}

		time.Sleep(5 * time.Second)

		buf = make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read conn err", err)
		}
		fmt.Printf("server call back : %s\n", buf[0:cnt])
	}
}
