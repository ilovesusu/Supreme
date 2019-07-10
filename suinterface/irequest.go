package suinterface

type IRequest interface {
	GetConnection() IConnection // 得到当前conn
	GetData() []byte            // 获取conn 的 data
}
