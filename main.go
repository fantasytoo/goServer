/*
*
@author: Administrator
@Date: 2022-11-10-0010
@Note
*
*/
package main

import (
	"go_test/client"
	"go_test/network/message"
)

func main() {
	isClose := make(chan bool)
	message.InitMsgId()
	client.CreateClient("10.23.26.128:11001")
	//结束进程
	<-isClose
}
