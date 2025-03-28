package msg

import (
	"bytes"
)

type EncoderStringState int

const (
	EncodeHeader EncoderStringState = iota
	EncodeBody
	EncodeDone
	Err
)

type Encoder struct {
	buffer  *bytes.Buffer
	state   EncoderStringState
	compile CompileState
	err     error
}

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
	BdoyLength
	Body
	BodyDone
)
