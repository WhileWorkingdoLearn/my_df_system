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
		Error:       msg.ErrDomain,
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
	msg := MSG{}
	msg.Decode(r)
	fmt.Println("Done")
}

func (msg MSG) Decode(r io.Reader) {
	buff := make([]byte, 8, 8)
	readIndex := 0
	readDone := false

	for !msg.done {

		if readIndex >= len(buff) {
			newBuff := make([]byte, len(buff)*2, len(buff)*2)
			copy(newBuff, buff)
			buff = newBuff
		}

		if !readDone {
			n, err := r.Read(buff[readIndex:])
			if err != nil {
				if err == io.EOF {
					readDone = true
				} else {
					fmt.Println(err)
					readIndex += n
					continue
				}
			}
			readIndex += n
		}

		sep, hasEnd := msg.parse(buff[:readIndex])
		msg.done = hasEnd

		if sep > 0 {
			sep = sep + 1
			newbuff := make([]byte, readIndex-sep, readIndex-sep)
			copy(newbuff, buff[sep:readIndex])
			buff = newbuff
			readIndex = readIndex - sep
		}

	}
}

type MSG struct {
	step int
	done bool
}

func (msg *MSG) Done() bool {
	return msg.done
}

func (msg *MSG) parse(buff []byte) (int, bool) {
	sep := bytes.Index(buff, []byte{':'})
	end := false
	if sep != -1 {
		fmt.Printf("%v\n", buff[:sep])
		switch msg.step {
		case 0:

		case 1:
		case 2:

			fmt.Println("step: ", msg.step)
		default:
			fmt.Println("undefined")
		}
		msg.step += 1
		if sep < len(buff)-1 && buff[sep+1] == byte(':') {
			end = true
		}
		return sep, end
	}

	return 0, end
}
