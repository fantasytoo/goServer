package tcp

type Message struct {
	Zip  uint16
	Code uint16
	Sid  uint32
	ID   uint32
	Data []byte
}

type MsgQueueData struct {
	Queue chan Message
}

// 获取消息ID
func (this *Message) GetMsgId() uint32 {
	return this.ID
}

// 获取消息内容
func (this *Message) GetData() []byte {
	return this.Data
}

// 获取消息内容
func (this *Message) GetIsPos() int64 {
	return 0
}
