package nmsgp

import (
	"context"
	"testing"

	"github.com/WhileCodingDoLearn/my_df_system/msg"
	"github.com/stretchr/testify/assert"
)

func TestConverMsgTypesToString(t *testing.T) {
	u := Request{
		msg.MsgHeader{},
		context.Background(),
	}

	u.msgHeader.MsgType = msg.IdxPING

	assert.Equal(t, "PING", u.MsgType())

	u.msgHeader.MsgType = msg.IdxPONG

	assert.Equal(t, "PONG", u.MsgType())

	u.msgHeader.MsgType = msg.IdxHEADER

	assert.Equal(t, "HEADER", u.MsgType())

	u.msgHeader.MsgType = msg.IdxERROR

	assert.Equal(t, "ERROR", u.MsgType())

	u.msgHeader.MsgType = msg.IdxEND

	assert.Equal(t, "END", u.MsgType())
}
