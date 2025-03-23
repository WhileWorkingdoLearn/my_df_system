package msg

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func TestEncoding(t *testing.T) {
	msgHheader := MsgHeader{
		Version:     1,
		MsgType:     2,
		Method:      3,
		Timestamp:   time.Now().UTC(),
		Timeout:     time.Duration(3) * time.Second,
		Domain:      "blabal",
		Endpoint:    "v1\\dv",
		HasAuth:     false,
		Auth:        "mykey",
		HasPayload:  false,
		PayloadType: 1,
		PayloadSize: 10,
	}
	data, err := EncodeMsgHeader(msgHheader)
	if err != nil {
		t.Fatal()
	}
	fmt.Println(data)
	msg, err := DecodeMsgHeader(bytes.NewBuffer(data))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(msg)
	//assert.Equal(t, msgHheader, msg, "")
}
