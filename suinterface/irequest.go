package suinterface

type IRequest interface {
	GetConnection() IConnection // 得到当前conn
	GetData() []byte            // 获取conn 的 data
	GetMsgID() uint32           //获取请求的消息的ID
}
