package sunet

import "github.com/ilovesusu/Supreme/suinterface"

type BaseRouter struct {
}

//处理conn之前的方法
func (b *BaseRouter) PreHandle(request suinterface.IRequest) {}

//处理conn时的主方法
func (b *BaseRouter) Handle(request suinterface.IRequest) {}

//处理conn之后的方法
func (b *BaseRouter) PostHandle(request suinterface.IRequest) {}
