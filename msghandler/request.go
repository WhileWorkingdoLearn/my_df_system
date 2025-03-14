package handler

import (
	"context"
	"time"

	"github.com/WhileCodingDoLearn/my_df_system/msg"
)

type Request struct {
	msgHeader msg.MsgHeader
	ctx       context.Context
}

func (r *Request) MsgType() string { return convertMsgType(r.msgHeader.MsgType).String() }

func (r *Request) Method() string { return convertMethod(r.msgHeader.Method).String() }

func (r *Request) Timestamp() time.Time {
	t := time.Unix(int64(r.msgHeader.Timestamp), 0)
	return t.UTC()
}

func (r *Request) Domain() string {
	return r.msgHeader.Domain
}

func (r *Request) Endpoint() string {
	return r.msgHeader.Endpoint
}

func (r *Request) Authentification() string {
	return r.msgHeader.Auth
}

func (r *Request) PayloadType() string {
	return convertMediaType(r.msgHeader.PayloadType).String()
}

func (r *Request) PayloadSyize() int {
	return r.msgHeader.PayloadSize
}

func (r *Request) ReadPayload() []byte {
	return nil
}
