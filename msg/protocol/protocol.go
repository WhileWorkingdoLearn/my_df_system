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
	SessionId []byte
	TimeStamp []byte
	Version   []byte
	Checksum  []byte
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
	FourBytes                 = TwoBytes * 27
	SixteenBytes   ByteLength = 16
	ThirtytwoBytes ByteLength = 32
	SixtyfourBytes ByteLength = 64
)

type Message struct {
	Header Header
	Body   Body
}
