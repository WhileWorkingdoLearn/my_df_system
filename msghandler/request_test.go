package handler

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

	u.msgHeader.MsgType = IdxPING

	assert.Equal(t, "PING", u.MsgType())

	u.msgHeader.MsgType = IdxPONG

	assert.Equal(t, "PONG", u.MsgType())

	u.msgHeader.MsgType = IdxHEADER

	assert.Equal(t, "HEADER", u.MsgType())

	u.msgHeader.MsgType = IdxERROR

	assert.Equal(t, "ERROR", u.MsgType())

	u.msgHeader.MsgType = idxEND

	assert.Equal(t, "END", u.MsgType())
}
