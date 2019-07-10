package net

import (
	"fmt"
	"github.com/ilovesusu/Supreme/suinterface"
	"github.com/ilovesusu/Supreme/utils"
	"net"
	"time"
)

type Serve struct {
	Name      string              //服务器名称
	IPVersion string              //服务器IP版本
	IP        string              //服务器IP
	Port      int                 //服务器端口
	Router    suinterface.IRouter //路由
}

//启动服务器
func (s *Serve) Start() {

	fmt.Printf("[START] Server name: %s,listenner at IP: %s, Port %d is starting\n", s.Name, s.IP, s.Port)
	fmt.Printf("[Supreme] Version: %s, MaxConn: %d,  MaxPacketSize: %d\n",
		utils.GlobalObject.Version,
		utils.GlobalObject.MaxConn,
		utils.GlobalObject.MaxPacketSize)

	//1 获取一个TCP的Addr
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println("Resolve TCP Addr error:")
		return
	}
	//2 监听服务器地址
	listenner, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		fmt.Println("listen :", s.IP, "err:", err)
		return
	}
	//监听成功
	fmt.Println("start server  ", s.Name, " success, now listenning...")
	//3 启动server网络连接业务
	for {
		//TODO server.go 应该有一个自动生成ID的方法
		var cid uint32
		cid = 0

		//3.1 阻塞等待客户端建立连接请求
		conn, err := listenner.AcceptTCP()
		if err != nil {
			fmt.Println("Accept err ", err)
			continue
		}
		//3.2 TODO Server.Start() 设置服务器最大连接控制,如果超过最大连接，那么则关闭此新的连接

		//3.3 处理该新连接请求的 业务 方法， 此时 handler 和 conn是绑定的
		connection := NewConnection(conn, cid, s.Router)
		cid++

		//启动当前的链接业务处理
		go connection.Start()
	}
}

//停止服务器
func (s *Serve) Stop() {
	fmt.Println("[STOP] Supreme server , name ", s.Name)
	//TODO  Server.Stop() 将其他需要清理的连接信息或者其他信息 也要一并停止或者清理
}

//运行服务器
func (s *Serve) Serve() {
	s.Start()
	//TODO Server.Serve() 是否在启动服务的时候 还要处理其他的事情呢 可以在这里添加
	//阻塞,否则主Go退出， listenner的go将会退出
	for {
		time.Sleep(10 * time.Second)
	}
}

func (s *Serve) AddRouter(router suinterface.IRouter) {
	s.Router = router
}

/*
初始化服务器
*/

func NewServer(name string) (serve suinterface.IServer) {
	//先初始化全局配置文件
	utils.GlobalObject.Reload()

	s := &Serve{
		Name:      utils.GlobalObject.Name, //从全局参数获取
		IPVersion: "tcp4",
		IP:        utils.GlobalObject.Host,    //从全局参数获取
		Port:      utils.GlobalObject.TcpPort, //从全局参数获取
		Router:    nil,
	}
	return s
}
