package main

import (
	sim "github.com/WhileCodingDoLearn/my_df_system/simulation"
)

func main() {
	buff := make([]byte, 4)
	_, err := sim.NewReader().Read(buff)
	if err != nil {

	}
	/*
	   ch := make(chan []byte, 1)

	   go func() {

	   		nMsg := msg.MsgHeader{
	   			Version:     1,
	   			MsgType:     msg.IdxDATA,
	   			Error:       msg.None,
	   			Method:      msg.IdxFETCH,
	   			Timestamp:   int(time.Now().Unix()),
	   			Timeout:     4 * time.Second,
	   			Domain:      "Test",
	   			Endpoint:    "/version",
	   			HasAuth:     false,
	   			Auth:        "",
	   			HasPayload:  true,
	   			PayloadType: msg.IdxJSON,
	   			PayloadSize: 10,
	   			Checksum:    0,
	   		}

	   		data, err := msg.EncodeMsgHeader(nMsg)
	   		if err != nil {
	   			log.Fatal(err)
	   		}
	   		ch <- data

	   		time.Sleep(2 * time.Second)
	   		payload := []byte("Hello World")
	   		ch <- payload
	   	}()

	   data := <-ch

	   r := bytes.NewReader(data)
	   msg, err := msg.DecodeMsgHeader(r)

	   	if err != nil {
	   		log.Fatal(err)
	   	}

	   fmt.Println(msg)
	   payload := <-ch

	   fmt.Println(string(payload))
	   close(ch)
	*/
}
