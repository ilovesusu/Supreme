package sunet

import "github.com/ilovesusu/Supreme/suinterface"

type Request struct {
	conn suinterface.IConnection //已经建立的连接
	data suinterface.IMessage    //客户端请求的数据
}

func (r *Request) GetConnection() suinterface.IConnection {
	return r.conn
}

//获取请求消息的数据
func (r *Request) GetData() []byte {
	return r.data.GetData()
}

//获取请求的消息的ID
func (r *Request) GetMsgID() uint32 {
	return r.data.GetMsgId()
}
