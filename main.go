package main

import (
	"bytes"
	"fmt"
	"log"

	msg "github.com/WhileCodingDoLearn/my_df_system/msg/protocol"
)

func main() {
	buff := bytes.NewBuffer(make([]byte, 0))
	buff.WriteString("Type")
	buff.WriteRune(':')
	buff.WriteString("SenderId")
	buff.WriteRune(':')
	buff.WriteString("Key")
	buff.WriteRune(':')
	buff.WriteString("TimeStamp")
	buff.WriteRune(':')
	buff.WriteString("Version")
	buff.Write([]byte(":-:"))

	testb := buff.Bytes()
	hp := msg.NewHeaderParser()
	err := hp.Parse(bytes.NewReader(testb), func(data []byte, idx int) error {
		fmt.Printf("%s , idx: %v\n", data, idx)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

}
