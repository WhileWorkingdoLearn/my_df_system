package main

import (
	"fmt"
	"log"

	transfer "github.com/WhileCodingDoLearn/my_df_system/protocol"
)

func main() {
	data, err := transfer.EncodeNodeMsgHeader(1, 2, "timestamp", "Lifespan", "domain", "endpoint", "auth")
	if err != nil {
		log.Fatal("")
	}
	fmt.Println(data)
	fmt.Println(len(data) * cap(data))
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
