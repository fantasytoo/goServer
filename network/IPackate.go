package network

type IPacket interface {
	GetHeadLen() uint16
	Pack(data TransitData, isPos int64) []byte
	Unpack(binaryData []byte) (IMessage, error, func(conn IConn))
}
