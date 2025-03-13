package main

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/WhileCodingDoLearn/my_df_system/msg"
)

func main() {

	msgHheader := msg.MsgHeader{
		Version:     1,
		MsgType:     2,
		Method:      3,
		Timestamp:   int(time.Now().Unix()),
		Timeout:     time.Duration(3) * time.Second,
		Domain:      "blabal",
		Endpoint:    "v1\\dv",
		HasAuth:     false,
		Auth:        "mykey",
		HasPayload:  false,
		PayloadType: 1,
		PayloadSize: 10,
	}

	data, err := msg.EncodeMsgHeader(msgHheader)
	if err != nil {
		log.Fatal("Ecode:", err)
	}

	fmt.Println(data)

	msg, err := msg.DecodeMsgHeader(bytes.NewBuffer(data))
	if err != nil {
		log.Fatal("Decode:", err)
	}

	fmt.Println(msgHheader)
	fmt.Println(msg)
}
