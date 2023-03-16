package router

import (
	"go_test/network"
	"go_test/network/utils"
	"sync"

	"google.golang.org/protobuf/proto"
)

type router struct {
	sync.RWMutex
	routers        map[uint32]func(conn network.IConn, content []byte)
	msgQueueRouter map[int32]func(content []byte)
}

var RouterMgr = _NewRouterMgr()

func _NewRouterMgr() IRouter {
	return &router{
		routers: make(map[uint32]func(conn network.IConn, content []byte)),
	}
}

func (r *router) AddRouter(msgObj proto.Message, handler func(conn network.IConn, content []byte)) {
	r.Lock()
	defer r.Unlock()
	protocolNum := utils.GetProtoId(msgObj)
	r.routers[protocolNum] = handler
}

func (r *router) RegisterMQ(msgQueueName string, handler func(content []byte)) {
	r.Lock()
	defer r.Unlock()
	protocolNum := utils.ProtocalNumber(msgQueueName)
	r.msgQueueRouter[protocolNum] = handler
}

func (r *router) ExecRouterFunc(conn network.IConn, message network.TransitData) {
	r.Lock()
	defer r.Unlock()
	handler := r.routers[message.MsgId]
	if handler != nil {
		handler(conn, message.Data)
	}
}
