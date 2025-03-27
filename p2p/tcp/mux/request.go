package nmsgp

import (
	"context"
	"time"

	"github.com/WhileCodingDoLearn/my_df_system/msg"
)

type Request struct {
	msgHeader msg.Message
	ctx       context.Context
}

func NewRequest(msg msg.Message, ctx context.Context) *Request {
	return &Request{msgHeader: msg, ctx: ctx}
}

func (r *Request) Method() string { return "" }

func (r *Request) Timestamp() time.Time {
	return time.Now()
}
