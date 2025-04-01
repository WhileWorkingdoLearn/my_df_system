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

type HeaderDataPosition int

const (
	MsgType HeaderDataPosition = iota
	SenderId
	SessionId
	TimeStamp
	Version
	HeaderDone
)

type BodyDataPosition int

const (
	PackedId BodyDataPosition = iota
	PacketPos
	PrevPacked
	NextPacket
	PacketLength
	Packet
	PacketDone
)
