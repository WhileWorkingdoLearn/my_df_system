package transfer

import "time"

type ResponseWriter interface {
	Version(v int)
	MsgType(msgType int)
	LifeTime(deltaT time.Duration)
	Auth(t string)
	Index(index int)
	Payload(data []byte)
	Write() []byte
}

type writer struct {
	msg ByteNodeMsg
}

func (rw *writer) Version(v int) {
	rw.msg.Version = [1]byte{uint8(v)}
}

func (rw *writer) MsgType(msgType int) {
	rw.msg.MsgType = [1]byte{uint8(msgType)}
}

func (rw *writer) LifeTime(deltaT time.Duration) {
	t := uint8(time.Now().Add(time.Duration(deltaT.Seconds())).Unix())
	rw.msg.Timestamp = [8]byte{t}
}

func (rw *writer) Auth(key string) {

}

func (rw *writer) Index(index int) {
	rw.msg.ChunkIndex = [4]byte{uint8(index)}
}

func (rw *writer) Payload(data []byte) {
	rw.msg.ChunkData = data
}

func (rw *writer) Write() []byte {
	return nil
}
