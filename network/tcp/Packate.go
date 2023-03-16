package tcp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"go_test/network"
	"go_test/network/utils"
)

const (
	ConstMsgLength = 2 //消息长度
	ConstMsgIdLen  = 12
)

type packet struct {
}

func NewPacket() network.IPacket {
	return &packet{}
}

func (dp *packet) GetHeadLen() uint16 {
	return ConstMsgLength
}

// //封包
// func (dp *packet) Pack(msg network.TransitData, ispos int64) []byte {
// 	msgLength := int32(len(msg.Data) + ConstMsgIdLen)
// 	dataLen := utils.IntToBytes(msgLength)
// 	dataId := utils.IntToBytes(msg.MsgId)
// 	return append(append(dataLen, dataId...), msg.Data...)
// }

// 封包
func (dp *packet) Pack(msg network.TransitData, ispos int64) []byte {
	msgLength := uint16(len(msg.Data) + ConstMsgIdLen)
	zip := utils.UintShortToBytes(0)
	code := utils.UintShortToBytes(1)
	sid := utils.UIntToBytes(1)
	fmt.Println(msg.MsgId)
	mid := utils.UIntToBytes(uint32(msg.MsgId))
	msdData := append(append(append(zip, code...), append(sid, mid...)...), msg.Data...)
	return append(utils.UintShortToBytes(msgLength), msdData...)
}

// 解包
func (dp *packet) Unpack(binaryData []byte) (network.IMessage, error, func(conn network.IConn)) {
	//创建一个从输入二进制数据的ioReader
	dataBuff := bytes.NewReader(binaryData)
	//只解压head的信息，得到dataLen和msgID
	msg := &Message{}
	//读zip
	if err := binary.Read(dataBuff, binary.BigEndian, &msg.Zip); err != nil {
		return nil, err, nil
	}
	//读code
	if err := binary.Read(dataBuff, binary.BigEndian, &msg.Code); err != nil {
		return nil, err, nil
	}
	//读sid
	if err := binary.Read(dataBuff, binary.BigEndian, &msg.Sid); err != nil {
		return nil, err, nil
	}
	//读msgID
	if err := binary.Read(dataBuff, binary.BigEndian, &msg.ID); err != nil {
		return nil, err, nil
	}
	fmt.Println("消息解包", msg)
	//读msgData
	msg.Data = binaryData[ConstMsgIdLen:]
	//这里只需要把head的数据拆包出来就可以了，然后再通过head的长度，再从conn读取一次数据
	return msg, nil, nil
}

// // 解包
// func (dp *packet) Unpack(binaryData []byte) (network.IMessage, error, func(conn network.IConn)) {
// 	//创建一个从输入二进制数据的ioReader
// 	dataBuff := bytes.NewReader(binaryData)
// 	//只解压head的信息，得到dataLen和msgID
// 	msg := &Message{}
// 	//读msgID
// 	if err := binary.Read(dataBuff, binary.BigEndian, &msg.ID); err != nil {
// 		return nil, err, nil
// 	}
// 	//读msgID
// 	msg.Data = binaryData[ConstMsgIdLen:]
// 	//这里只需要把head的数据拆包出来就可以了，然后再通过head的长度，再从conn读取一次数据
// 	return msg, nil, nil
// }
