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
		Timestamp:   time.Now().UTC(),
		Timeout:     time.Duration(3) * time.Second,
		Domain:      "blabal",
		Endpoint:    "v1dv",
		HasAuth:     false,
		Auth:        "",
		HasPayload:  false,
		PayloadType: 0,
		PayloadSize: 0,
	}

	fmt.Println(msgHheader)
	encoded, err := msg.EncodeMsgHeader(msgHheader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(encoded)
	r := sim.NewStream()
	r.Start(encoded)

	h, err := msg.DecodeMsgHeader(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(h)

}
