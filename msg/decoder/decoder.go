package msg

import (
	"io"

	msg "github.com/WhileCodingDoLearn/my_df_system/msg/protocol"
)

type Decoder struct {
	reader       io.Reader
	headerParser func(reader io.Reader, setter msg.Handler) error
	bodyParser   func(reader io.Reader, setter msg.Handler) error
}

type HeaderParserPosition int

const (
	MsgType HeaderParserPosition = iota
	SenderId
	Key
	TimeStamp
	Version
	HeaderDone
)

type BodyParserPosition int

const (
	PackedId BodyParserPosition = iota
	PacketPos
	PrevPacked
	NextPacket
	PacketLength
	Packet
	PacketDone
)
