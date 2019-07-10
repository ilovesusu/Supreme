package suinterface

type IRouter interface {
	PreHandle(request IRequest)  //处理conn之前的方法
	Handle(request IRequest)     //处理conn时的主方法
	PostHandle(request IRequest) //处理conn之后的方法
}
