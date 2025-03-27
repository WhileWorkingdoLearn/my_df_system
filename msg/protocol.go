package msg

type ParserState int

const (
	ReadHeader ParserState = iota
	ReadBody
	Done
	ErrorHeader
	ErrorBody
)

type ParserHandler map[byte]ParseFunc

type ParseFunc func(data []byte) (n int, done bool, err error)

func (pf ParseFunc) Handle(data []byte) (n int, done bool, err error) {
	parsed, done, err := pf(data)
	return parsed, done, err
}

type HeaderParserState int

const (
	Sep       = ":"
	HeaderEnd = "-:"
	MSgEnd    = "|:"
)

const (
	MsgType HeaderParserState = iota
	SenderId
	Key
	TimeStamp
	Version
	HeaderDone
)

type Header struct {
	MsgType   []byte
	SenderId  []byte
	Key       []byte
	TimeStamp []byte
	Version   []byte
}

type BodyParserState int

const (
	PackedId BodyParserState = iota
	PacketPos
	PrevPacked
	NextPacket
	PacketLength
	Packet
	PacketDone
)

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
)

type Message struct {
	header *Header
	Body   *Body
}
