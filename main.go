package main

import (
	"fmt"
	"log"
	"time"

	"github.com/WhileCodingDoLearn/my_df_system/msg"
	sim "github.com/WhileCodingDoLearn/my_df_system/simulation"
)

func main() {

	msgHheader := msg.MsgHeader{
		Version:     1,
		MsgType:     msg.IdxPING,
		Error:       msg.None,
		Method:      msg.IdxFETCH,
		Timestamp:   int(time.Now().Unix()),
		Timeout:     time.Duration(3) * time.Second,
		Domain:      "blabal",
		Endpoint:    "v1\\dv",
		HasAuth:     false,
		Auth:        "mykey",
		HasPayload:  false,
		PayloadType: 0,
		PayloadSize: 0,
	}

	//buff := make([]byte, 32)
	encoded, err := msg.EncodeMsgHeader(msgHheader)
	if err != nil {
		log.Fatal(err)
	}

	r := sim.NewStream(encoded)
	n, err := msg.DecodeMsgHeader(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(n)

}
