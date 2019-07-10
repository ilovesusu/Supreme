package net

import (
	"fmt"
	"github.com/ilovesusu/Supreme/suinterface"
	"net"
)

/*
	链接模块
*/
type Connection struct {
	Conn         *net.TCPConn        // Conn 的 socket
	ConnID       uint32              // Conn 的 ID
	IsClose      bool                // Conn的状态
	ExitBuffChan chan bool           //告知conn已经关闭退出
	Router       suinterface.IRouter //Conn 的处理方法Router
}

//初始化连接模块的函数

func NewConnection(conn *net.TCPConn, connID uint32, router suinterface.IRouter) *Connection {
	c := &Connection{
		Conn:         conn,
		ConnID:       connID,
		IsClose:      false,
		Router:       router,
		ExitBuffChan: make(chan bool),
	}
	return c
}

// 启动连接
func (c *Connection) Start() {
	//1 开启用于写回客户端数据流程的Goroutine
	//go c.StartWriter()
	//2 开启用户从客户端读取数据流程的Goroutine
	go c.StartReader()
}

// 关闭连接
func (c *Connection) Stop() {
	fmt.Println("Conn Close Connid=", c.ConnID)

	if c.IsClose == true {
		return
	}
	c.IsClose = true
	//关闭Conn
	c.Conn.Close()
	//关闭管道,回收资源
	close(c.ExitBuffChan)
}

// 获取conn
func (c *Connection) GetConnection() *net.TCPConn {
	return c.Conn
}

// 获取conn 对应的ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// 获取远程客户端的状态 ip port
func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// 发送数据
func (c *Connection) Send(data []byte) error {
	return nil
}

func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is  running")
	defer fmt.Println(c.GetRemoteAddr().String(), " conn reader exit!")
	defer c.Stop()

	for {
		//读取我们最大的数据到buf中
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err ", err)
			c.ExitBuffChan <- true
			continue
		}

		//得到当前conn的request请求

		req := Request{
			conn: c,
			data: buf,
		}

		//调用conn中绑定的router调用

		go func(request suinterface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)

	}

}

func (c *Connection) StartWriter() {

}
