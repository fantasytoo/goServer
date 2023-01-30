/**
@author: Administrator
@Date: 2022-12-01-0001
@Note
**/
package client

import (
	"go_test/log"
	"go_test/network"
	"go_test/network/router"
	"go_test/network/tcp"
	"time"
)

type ClientStruct struct {
	conn    network.IConn
	IsClose chan bool
}

func CreateClient(addr string) {
	route := router.RouterMgr
	client := &ClientStruct{
		IsClose: make(chan bool, 1),
	}
	defer close(client.IsClose)
	tcpConn := tcp.NewClient(tcp.NewPacket(), addr, network.WithMax(10, time.Now().UnixNano()/1e6, nil), route, client.Stop).Connect()
	if tcpConn == nil {
		client.IsClose <- true
	} else {
		client.conn = tcpConn
		route.AddRouterByString("111", client.getRoleData)
	}
	<-client.IsClose
	log.AppLogger.Info("客户端关闭")
}

func (*ClientStruct) getRoleData(conn network.IConn, data []byte) {

}

func (server *ClientStruct) Stop(conn network.IConn) {
	server.IsClose <- true
}
