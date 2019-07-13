package sunet

import (
	"errors"
	"fmt"
	"github.com/ilovesusu/Supreme/suinterface"
	"io"
	"net"
)

/*
	链接模块
*/
type Connection struct {
	Conn         *net.TCPConn           // Conn 的 socket
	ConnID       uint32                 // Conn 的 ID
	IsClose      bool                   // Conn的状态
	ExitBuffChan chan bool              // 告知conn已经关闭退出
	MsgHandle    suinterface.IMsgHandle // Conn 的处理方法Router
	msgChan      chan []byte            // 无缓冲管道，用于读、写两个goroutine之间的消息通信
}

//初始化连接模块的函数

func NewConnection(conn *net.TCPConn, connID uint32, msgHandle suinterface.IMsgHandle) *Connection {
	c := &Connection{
		Conn:         conn,
		ConnID:       connID,
		IsClose:      false,
		MsgHandle:    msgHandle,
		ExitBuffChan: make(chan bool),
		msgChan:      make(chan []byte),
	}
	return c
}

// 启动连接
func (c *Connection) Start() {
	//1 开启用于写回客户端数据流程的 Goroutine
	go c.StartWriter()
	//2 开启用户从客户端读取数据流程的 Goroutine
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
func (c *Connection) GetTCPConnection() *net.TCPConn {
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

//直接将Message数据发送数据给远程的TCP客户端
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.IsClose == true {
		return errors.New("Connection closed when send msg")
	}
	//将data封包，并且发送
	dp := NewDataPack()
	msg, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		fmt.Println("Pack error msg id = ", msgId)
		return errors.New("Pack error msg ")
	}

	//写入消息管道,由写 Goroutine 发送到客户端
	c.msgChan <- msg

	return nil
}

func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is  running")
	defer fmt.Println(c.GetRemoteAddr().String(), " conn reader exit!")
	defer c.Stop()

	for {
		// 创建拆包解包的对象
		dataPack := NewDataPack()
		//读取客户端的Msg head
		headData := make([]byte, dataPack.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("read msg head error ", err)
			break
		}
		//拆包，得到msgid 和 datalen 放在msg中
		msg, err := dataPack.Unpack(headData)
		if err != nil {
			fmt.Println("unpack error ", err)
			break
		}
		//根据 dataLen 读取 data，放在msg.Data中
		var data []byte
		if msg.GetDataLen() > 0 {
			data = make([]byte, msg.GetDataLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read msg data error ", err)
				break
			}
		}
		msg.SetData(data)
		//得到当前客户端请求的Request数据
		req := Request{
			conn: c,
			data: msg,
		}

		//调用conn中绑定的router调用
		go c.MsgHandle.DoMsgHandle(&req)

	}

}

func (c *Connection) StartWriter() {
	fmt.Println("Writer Goroutine is  running")
	defer fmt.Println(c.GetRemoteAddr().String(), " conn Writer exit!")

	for {
		select {
		case data := <-c.msgChan:
			//有数据要写给客户端
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("Send Data error:, ", err, " Conn Writer exit")
				c.ExitBuffChan <- true
				return
			}
		case <-c.ExitBuffChan:
			//conn已经关闭
			return
		}
	}
}
