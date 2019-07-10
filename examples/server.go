package main

import (
	"fmt"
	"github.com/ilovesusu/Supreme/net"
	"github.com/ilovesusu/Supreme/suinterface"
)

func main() {
	serve := net.NewServer("Susu TCP [V0.3]")
	// 添加自定义路由在server中
	serve.AddRouter(&PingRouter{})
	serve.Serve()
}

/*
ping test 自定义路由
*/

type PingRouter struct {
	net.BaseRouter
}

//处理conn之前的方法
func (p *PingRouter) PreHandle(request suinterface.IRequest) {
	fmt.Println("调用路由的 PreHandle")
	_, err := request.GetConnection().GetConnection().Write([]byte("调用路由的 PreHandle\n"))
	if err != nil {
		fmt.Println("调用路由的 PreHandle 错误")
	}
}

//处理conn时的主方法
func (p *PingRouter) Handle(request suinterface.IRequest) {
	fmt.Println("调用路由的 Handle")
	_, err := request.GetConnection().GetConnection().Write([]byte("pong,pong,pong ... 调用路由的 Handle\n"))
	if err != nil {
		fmt.Println("调用路由的 Handle 错误")
	}
}

//处理conn之后的方法
func (p *PingRouter) PostHandle(request suinterface.IRequest) {
	fmt.Println("调用路由的 PostHandle")
	_, err := request.GetConnection().GetConnection().Write([]byte("调用路由的 PostHandle\n"))
	if err != nil {
		fmt.Println("调用路由的 PostHandle 错误")
	}
}
