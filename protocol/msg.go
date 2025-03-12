package transfer

import (
	"time"
)

/*
	Version     int
	MsgType     int
	Method      string
	Timestamp   time.Time
	Timeout     time.Duration
	Domain      string
	Endpoint    string
	Auth        string
	PayloadType string
*/

type MSgContructor interface {
	MsgType(msgtype int) error
	Method(m int) error
	LifeTime(deltaTime time.Duration) error
	Domain(name string) error
	Endpoint(path string) error
	Auth(key string) error
	Payload([]byte)
}

type IntByteType int
type StringByteType string

func (ib IntByteType) To1Byte() [1]byte { return [1]byte{uint8(ib)} }

func (ib IntByteType) To4Bytes() [4]byte { return [4]byte{} }

func (ib IntByteType) To8Bytes() [8]byte { return [8]byte{} }

func (sb StringByteType) To4Byte() [4]byte   { return [4]byte{} }
func (sb StringByteType) To8Bytes() [8]byte  { return [8]byte{} }
func (sb StringByteType) To16Byte() [16]byte { return [16]byte{} }
func (sb StringByteType) To32Byte() [32]byte { return [32]byte{} }

type allowedMsgTypes struct {
	MsgTypePing  IntByteType
	MsgTypePong  IntByteType
	MsgTypeData  IntByteType
	MsgTypeAck   IntByteType
	MsgTypeError IntByteType
}

var MsgTypes = allowedMsgTypes{
	MsgTypePing:  0,
	MsgTypePong:  1,
	MsgTypeData:  2,
	MsgTypeAck:   3,
	MsgTypeError: 3,
}

type allowedDataTypes struct {
	None      IntByteType
	Json      IntByteType
	Text_Utf8 IntByteType
	Binary    IntByteType
	Picture   IntByteType
	Video     IntByteType
	Audio     IntByteType
	Custom    IntByteType
}

var DataTypes = allowedDataTypes{
	None:      data_unkown,
	Json:      data_json,
	Text_Utf8: data_text_utf8,
	Binary:    data_binary,
	Picture:   data_picture,
	Video:     data_video,
	Audio:     data_audio,
	Custom:    data_custom,
}

type methods struct {
	CREATE IntByteType
	FETCH  IntByteType
	UPDATE IntByteType
	DELETE IntByteType
}

var MethodTypes = methods{
	CREATE: 0,
	FETCH:  1,
	UPDATE: 2,
	DELETE: 3,
}

