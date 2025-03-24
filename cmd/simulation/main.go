package main

import (
	"bytes"
	"encoding/binary"
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
		HasAuth:     true,
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

	var msgH msg.MsgHeader

	dec := NewDecoder(r)
	errDec := dec.Decode(&msgH)
	if errDec != nil {

	}
	fmt.Println(msgH)
	fmt.Println("Done")
}

type Decoder struct {
	reader io.Reader
	buffer []byte
	parsed int
	step   int
	read   int
	done   bool
	err    error
}

func NewDecoder(reader io.Reader) *Decoder {
	return &Decoder{reader: reader, buffer: make([]byte, 8, 8)}
}

func (decoder *Decoder) Parsed() int {
	return decoder.parsed
}

func (decoder *Decoder) Decode(msgHeader *msg.MsgHeader) error {
	readIndex := 0
	readDone := false

	for !decoder.done {

		if readIndex >= len(decoder.buffer) {
			newBuff := make([]byte, len(decoder.buffer)*2, len(decoder.buffer)*2)
			copy(newBuff, decoder.buffer)
			decoder.buffer = newBuff
		}

		if !readDone {
			n, err := decoder.reader.Read(decoder.buffer[readIndex:])
			if err != nil {
				if err == io.EOF {
					readDone = true
				} else {
					fmt.Println(err)
					decoder.err = err
					return err
				}
			}
			decoder.read += n
			readIndex += n
		}

		sep, hasEnd, errParse := decoder.parse(decoder.buffer[:readIndex], msgHeader)
		if errParse != nil {
			fmt.Println(errParse)
		}
		decoder.parsed += sep
		decoder.done = hasEnd

		if sep > 0 {
			sep = sep + 1
			newbuff := make([]byte, readIndex-sep, readIndex-sep)
			copy(newbuff, decoder.buffer[sep:readIndex])
			decoder.buffer = newbuff
			readIndex = readIndex - sep
		}

	}
	return nil
}

func (msg *Decoder) Done() bool {
	return msg.done
}

const (
	version int = iota
	msgtype
	errtype
	methodType
	timestamp
	timeout
	domain
	endpoint
	hasAuth
	auth
	haspayload
	payloadtype
	payloadlength
)

func (decoder *Decoder) parse(buff []byte, msgHeader *msg.MsgHeader) (int, bool, error) {
	separator := bytes.Index(buff, []byte{':'})
	closingtag := false
	if separator != -1 {
		if separator < len(buff)-2 && buff[separator+1] == byte('-') && buff[separator+2] == byte(':') {
			closingtag = true
		}
		data := buff[:separator]

		switch decoder.step {
		case version:

			if len(data) < 1 {
				return separator, closingtag, fmt.Errorf("error with version")
			}
			vers := int(data[0])
			msgHeader.Version = vers

		case msgtype:

			if len(data) < 1 {
				return separator, closingtag, fmt.Errorf("error with msgtype")
			}
			msgTpye := int(data[0])
			msgHeader.MsgType = msgTpye

		case errtype:

			if len(data) < 4 {
				return separator, closingtag, fmt.Errorf("error with errortype")
			}
			errT := int(binary.BigEndian.Uint16(data))
			msgHeader.Error = errT

		case methodType:

			if len(data) < 1 {
				return separator, closingtag, fmt.Errorf("error with methodtype")
			}
			mType := int(data[0])
			msgHeader.Method = mType

		case timestamp:

			if len(data) < 8 {
				return separator, closingtag, fmt.Errorf("error with timestamp")
			}
			tstmp := binary.BigEndian.Uint64(data)
			msgHeader.Timestamp = time.Unix(int64(tstmp), 0).UTC()

		case timeout:

			if len(data) < 4 {
				return separator, closingtag, fmt.Errorf("error with timeout")
			}
			lifetime := int(binary.BigEndian.Uint32(data))
			msgHeader.Timeout = time.Duration(lifetime) * time.Second

		case domain:

			if len(data) < 4 {
				return separator, closingtag, fmt.Errorf("error with domain")
			}
			dom := string(data)
			msgHeader.Domain = dom

		case endpoint:

			if len(data) < 4 {
				return separator, closingtag, fmt.Errorf("error with endpoint")
			}
			ep := string(data)
			msgHeader.Endpoint = ep

		case hasAuth:

			if len(data) < 1 {
				return separator, closingtag, fmt.Errorf("error with hasAuth")
			}
			authFlag := int(data[0])
			if auth == 0 {
				decoder.step += 1
			}
			msgHeader.HasAuth = authFlag == 1

		case auth:

			if len(data) < 4 {
				return separator, closingtag, fmt.Errorf("error with authkey")
			}
			auth := string(data)
			msgHeader.Auth = auth

		case haspayload:

			if len(data) < 1 {
				return separator, closingtag, fmt.Errorf("error with hasPayload")
			}
			payloadflag := int(data[0])
			if payloadflag == 0 {
				decoder.step += 2
			}
			msgHeader.HasPayload = payloadflag == 1

		case payloadtype:
			if len(data) < 1 {
				return separator, closingtag, fmt.Errorf("error with payloadtype")
			}
			ptype := int(data[0])
			msgHeader.PayloadType = ptype

		case payloadlength:

		default:
			fmt.Println("undefined")
		}
		decoder.step += 1

		return separator, closingtag, nil
	}

	return 0, closingtag, nil
}
