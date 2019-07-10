package net

import "github.com/ilovesusu/Supreme/suinterface"

type Request struct {
	conn suinterface.IConnection //已经建立的连接
	data []byte                  //客户端请求的数据
}

func (r *Request) GetConnection() suinterface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}
