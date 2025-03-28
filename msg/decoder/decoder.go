package msg

import (
	"io"

	msg "github.com/WhileCodingDoLearn/my_df_system/msg/protocol"
)

type DecoderState int

const (
	ReadHeader DecoderState = iota
	ReadBody
	ReadDone
	ErrorHeader
	ErrorBody
)

type ParserHandler map[byte]ParseFunc

type ParseFunc func(data []byte) (n int, done bool, err error)

func (pf ParseFunc) Handle(data []byte) (n int, done bool, err error) {
	parsed, done, err := pf(data)
	return parsed, done, err
}

type Decoder struct {
	reader       io.Reader
	persers      ParserHandler
	headerparser HeaderParser
	state        DecoderState
	err          error
}

type HeaderParserState int

const (
	MsgType HeaderParserState = iota
	SenderId
	Key
	TimeStamp
	Version
	HeaderDone
)

type HeaderParser struct {
	header *msg.Header
	state  HeaderParserState
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

type BodyParser struct {
	body  *msg.Body
	state BodyParserState
}
