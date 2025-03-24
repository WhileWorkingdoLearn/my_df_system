package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/WhileCodingDoLearn/my_df_system/msg"
)

type Stream struct {
	Data chan byte
	wg   sync.WaitGroup
}

func NewStream() *Stream {
	ch := make(chan byte, 16)
	return &Stream{Data: ch}
}

func (s *Stream) Start(data []byte) {
	s.wg.Wait()
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		defer close(s.Data)
		for _, item := range data {

			min := 10
			max := 20
			randomIntInRange := rand.Intn(max-min+1) + min
			time.Sleep(time.Duration(randomIntInRange) * time.Millisecond)

			s.Data <- item
		}

	}()
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

func (s *Stream) Read(p []byte) (n int, err error) {
	for i := range p {
		b, ok := <-s.Data
		if !ok {
			if n == 0 {
				return 0, io.EOF
			}
			return n, io.EOF
		}
		p[i] = b
		n++
	}
	return n, nil
}

func main() {
	msgHheader := msg.MsgHeader{
		Version:     1,
		MsgType:     msg.IdxPING,
		Error:       msg.ErrDomain,
		Method:      msg.IdxFETCH,
		Timestamp:   time.Now().UTC(),
		Timeout:     time.Duration(3) * time.Second,
		Domain:      "blabal",
		Endpoint:    "v1dv",
		HasAuth:     false,
		Auth:        "2qe2qeqe2",
		HasPayload:  true,
		PayloadType: msg.IdxJSON,
		PayloadSize: 128,
	}

	fmt.Println(msgHheader)
	encoded, err := msg.EncodeMsgHeader(msgHheader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(encoded)
	r := NewStream()
	r.Start(encoded)

	var msgH msg.MsgHeader

	dec := msg.NewDecoder(bytes.NewReader(encoded))
	errDec := dec.Decode(&msgH)
	if errDec != nil {

	}
	fmt.Println(msgH)
	fmt.Println("Done")
}
