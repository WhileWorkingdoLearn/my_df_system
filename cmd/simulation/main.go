package main

import (
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
	ch := make(chan byte, 1)
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
			max := 100
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
	r := NewStream()
	r.Start(encoded)

	h, err := msg.DecodeMsgHeader(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(h)

}
