package msg

import (
	"io"

	msg "github.com/WhileCodingDoLearn/my_df_system/msg/protocol"
)

type Encoder struct {
	reader       io.Reader
	headerParser func(reader io.Reader, setter msg.Handler) error
	bodyParser   func(reader io.Reader, setter msg.Handler) error
}

type EncoderStringState int

const (
	EncodeHeader EncoderStringState = iota
	EncodeBody
	EncodeDone
	Err
)

type EncodeType func(data string) error

type CompileState int

const (
	InitHeader CompileState = iota
	MsgType
	SenderId
	Key
	TimeStamp
	Version
	HeaderDone
	InitBody
	BodyId
	BodyPos
	PrevMsg
	NextMsg
	PayloadLength
	Payload
	BodyDone
)
