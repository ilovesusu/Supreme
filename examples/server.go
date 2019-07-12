package main

import (
	"fmt"
	"github.com/ilovesusu/Supreme/suinterface"
	"github.com/ilovesusu/Supreme/sunet"
	"strconv"
)

func main() {
	serve := sunet.NewServer("Susu TCP [V0.3]")
	// 添加自定义路由在server中
	serve.AddRouter(0, &PingRouter{})
	serve.AddRouter(1, &HelloRouter{})
	//开启服务
	serve.Serve()
}

/*
hello test 自定义路由
*/

type HelloRouter struct {
}

func (*HelloRouter) PreHandle(request suinterface.IRequest) {
	err := request.GetConnection().SendMsg(1, []byte("调用路由的 PreHandle\n"))
	if err != nil {
		fmt.Println("调用路由的 PreHandle 错误")
	}
}

func (*HelloRouter) Handle(request suinterface.IRequest) {
	connID := request.GetConnection().GetConnID()
	binaryMsg := []byte("userid:" + strconv.FormatUint(uint64(connID), 10) + ",hello Supreme!!!\n")
	err := request.GetConnection().SendMsg(1, binaryMsg)
	if err != nil {
		fmt.Println("调用路由的 Handle 错误")
	}
}

func (*HelloRouter) PostHandle(request suinterface.IRequest) {
	err := request.GetConnection().SendMsg(1, []byte("调用路由的 PostHandle\n"))
	if err != nil {
		fmt.Println("调用路由的 PostHandle 错误")
	}
}

/*
ping test 自定义路由
*/

type PingRouter struct {
	sunet.BaseRouter
}

//处理conn之前的方法
func (p *PingRouter) PreHandle(request suinterface.IRequest) {
	err := request.GetConnection().SendMsg(0, []byte("调用路由的 PreHandle\n"))
	if err != nil {
		fmt.Println("调用路由的 PreHandle 错误")
	}
}

//处理conn时的主方法
func (p *PingRouter) Handle(request suinterface.IRequest) {
	connID := request.GetConnection().GetConnID()
	binaryMsg := []byte(strconv.FormatUint(uint64(connID), 10) + "pong,pong,pong ... 调用路由的 Handle\n")
	err := request.GetConnection().SendMsg(0, binaryMsg)
	if err != nil {
		fmt.Println("调用路由的 Handle 错误")
	}
}

//处理conn之后的方法
func (p *PingRouter) PostHandle(request suinterface.IRequest) {
	err := request.GetConnection().SendMsg(0, []byte("调用路由的 PostHandle\n"))
	if err != nil {
		fmt.Println("调用路由的 PostHandle 错误")
	}
}
