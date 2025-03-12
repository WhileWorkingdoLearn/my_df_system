package main

import (
	"bytes"
	"fmt"
	"log"
	"time"

	transfer "github.com/WhileCodingDoLearn/my_df_system/encoding"
)

func main() {

	msgHheader := transfer.MsgHeader{
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

	data, err := transfer.EncodeMsg(msgHheader)
	if err != nil {
		log.Fatal()
	}

	fmt.Println(data)

	msg, err := transfer.DecodeMsgStream(bytes.NewBuffer(data))
	if err != nil {
		log.Fatal()
	}

	fmt.Println(msg)
}

/*
duration := time.Duration(2) * time.Second

	// Create a buffer to hold the 8-byte field
	buf := new(bytes.Buffer)

	// Write the duration as an int64 to the buffer
	err := binary.Write(buf, binary.LittleEndian, duration.Milliseconds())
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	// Get the 8-byte field
	byteField := buf.Bytes()

	// Print the byte field
	fmt.Printf("8-byte field: %x %v\n", byteField, byteField)*/
