package encoding

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEncoding(t *testing.T) {
	msgHheader := MsgHeader{
		Version:     1,
		MsgType:     2,
		Method:      3,
		Timestamp:   int(time.Now().Unix()),
		Timeout:     time.Duration(3) * time.Second,
		Domain:      "blabal",
		Endpoint:    "v1/dv",
		Auth:        "mykey",
		PayloadType: 1,
		PayloadSize: 10,
	}

	data, err := EncodeMsg(msgHheader)
	if err != nil {
		t.Fatal()
	}

	msg, err := DecodeMsgStream(bytes.NewBuffer(data))
	if err != nil {
		t.Fatal()
	}

	assert.Equal(t, msgHheader, msg, "")
}
