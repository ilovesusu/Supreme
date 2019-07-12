package sunet

import (
	"fmt"
	"github.com/ilovesusu/Supreme/suinterface"
	"strconv"
)

type MsgHandle struct {
	//存放msgid对应的处理方法
	Apis map[uint32]suinterface.IRouter
}

//创建msgHanldle
func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis: make(map[uint32]suinterface.IRouter),
	}
}

//调度执行对应的Router的消息处理方法
func (m *MsgHandle) DoMsgHandle(request suinterface.IRequest) {
	//1.找到对应msgid对应的router
	router, ok := m.Apis[request.GetMsgID()]
	if !ok {
		fmt.Println("api msgId = ", request.GetMsgID(), " is not FOUND!")
		return
	}

	//2.执行对应处理方法
	router.PreHandle(request)
	router.Handle(request)
	router.PostHandle(request)
}

//为消息添加处理的逻辑
func (m *MsgHandle) AddRouter(msgid uint32, router suinterface.IRouter) {
	//1.判断当前msgid绑定的api方法是否重复
	if _, ok := m.Apis[msgid]; ok {
		msgid := strconv.FormatUint(uint64(msgid), 10)
		panic("repeated api , msgId = " + msgid)
	}
	//2.添加msgid与api方法的绑定
	m.Apis[msgid] = router
}
