package handler

import "time"

type Request interface {
	Version() int
	MsgType() int
	LifeTime() time.Duration
	Auth() string
	Size() int
	Index() int
	Payload() []byte
}

type request struct {
	version     int
	msgType     int
	timeout     time.Duration
	auth        string
	payloadType string
	size        int
	index       int
	payload     []byte
}

func (r *request) Version() int {
	return r.version
}

func (r *request) MsgType() int {
	return r.msgType
}

func (r *request) LifeTime() time.Duration {
	return r.timeout
}

func (r *request) Auth() string {
	return r.auth
}

func (r *request) Size() int {
	return r.index
}

func (r *request) Index() int {
	return r.index
}

func (r *request) Payload() []byte {
	return r.payload
}
