package msg

type ProtocolToken string

const (
	Sep       ProtocolToken = ":"
	HeaderEnd ProtocolToken = "-:"
	MSgEnd    ProtocolToken = "|:"
)

type Header struct {
	MsgType   []byte
	SenderId  []byte
	Key       []byte
	TimeStamp []byte
	Version   []byte
}

type Body struct {
	PackedId     []byte
	PacketPos    []byte
	PrevPacked   []byte
	NextPacket   []byte
	PacketLength []byte
	Packet       []byte
}

type ByteLength int

const (
	Onebyte ByteLength = iota + 1
	TwoBytes
	FourBytes = TwoBytes * 2
)

type Message struct {
	Header Header
	Body   Body
}
