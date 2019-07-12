package suinterface

import "net"

type IConnection interface {
	Start()                                  // 启动连接
	Stop()                                   // 关闭连接
	GetTCPConnection() *net.TCPConn          // 获取conn
	GetConnID() uint32                       // 获取conn 对应的ID
	GetRemoteAddr() net.Addr                 // 获取远程客户端的状态 ip port
	SendMsg(msgId uint32, data []byte) error // 发送数据
}

// 处理conn的方法
type HandleFunc func(*net.TCPConn, []byte, int) error
