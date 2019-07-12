package suinterface

type IMsgHandle interface {
	DoMsgHandle(request IRequest)           //调度执行对应的Router的消息处理方法
	AddRouter(msgid uint32, router IRouter) //为消息添加处理的逻辑
}
