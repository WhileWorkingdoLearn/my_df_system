package msg

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"time"
)

//Message Decoder. Converts a Message Header from its binary representation to a Msg Struct format.
//This Decoder throws an error if it encounters an error with byte size aligment.
//If an error it thrown, the whole message is seen as invalid.
//
//Example:
//	r := bytes.NewReader(/*header bytes*/)
//	var msgH msg.MsgHeader
//	dec := NewDecoder(r)
//	errDec := dec.Decode(&msgH)
//	if errDec != nil {
//		fmt.Println(errDec)
//	}

type Decoder struct {
	reader io.Reader
	buffer *bytes.Buffer
	parsed int
	step   int
	read   int
	done   bool
	err    error
}

func NewDecoder(reader io.Reader) *Decoder {
	return &Decoder{reader: reader, buffer: bytes.NewBuffer(make([]byte, 0))}
}

func (decoder *Decoder) Parsed() int {
	return decoder.parsed
}

func (decoder *Decoder) Buffered() []byte {
	return decoder.buffer.Bytes()
}

func (msg *Decoder) Done() bool {
	return msg.done
}

func (decoder *Decoder) Decode(msgHeader *MsgHeader) error {
	if msgHeader == nil {
		return fmt.Errorf("message header is nil")
	}
	readIndex := 0
	readDone := false
	readbuffer := make([]byte, 8, 8)
	for !decoder.done {

		if readIndex >= len(readbuffer) {
			newBuff := make([]byte, len(readbuffer)*2, len(readbuffer)*2)
			copy(newBuff, readbuffer)
			readbuffer = newBuff
		}

		if !readDone {
			n, err := decoder.reader.Read(readbuffer[readIndex:])
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

		parsed, hasEnd, errParse := decoder.parse(readbuffer[:readIndex], msgHeader)
		if errParse != nil {
			decoder.err = errParse
			return errParse
		}
		fmt.Println("Parsed: ", parsed)
		decoder.parsed += parsed
		decoder.done = hasEnd
		if hasEnd {
			decoder.buffer.Reset()
			fmt.Println(msgHeader)
		}
		if parsed > 0 {
			parsed = parsed + 1
			newbuff := make([]byte, readIndex-parsed, readIndex-parsed)
			copy(newbuff, readbuffer[parsed:readIndex])
			readbuffer = newbuff
			readIndex = readIndex - parsed
		}

	}
	return nil
}

func (decoder *Decoder) parse(buff []byte, msgHeader *MsgHeader) (int, bool, error) {
	separator := bytes.Index(buff, []byte{':'})
	closingtag := false
	if separator != -1 {
		if separator < len(buff)-2 && buff[separator+1] == byte('-') && buff[separator+2] == byte(':') {
			closingtag = true
		}
		data := buff[:separator]
		parsed := len(data)
		fmt.Printf("data l: %v icx: %v\n", len(data), decoder.step)
		switch decoder.step {
		case version:

			if len(data) != 1 {
				return parsed, closingtag, fmt.Errorf("error with version")
			}
			vers := int(data[0])
			msgHeader.Version = vers
			decoder.buffer.Write(data)
		case msgtype:

			if len(data) != 1 {
				return parsed, closingtag, fmt.Errorf("error with msgtype")
			}
			msgTpye := int(data[0])
			msgHeader.MsgType = msgTpye
			decoder.buffer.Write(data)

		case errtype:

			if len(data) != 4 {
				return parsed, closingtag, fmt.Errorf("error with errortype")
			}
			errT := int(binary.BigEndian.Uint16(data))
			msgHeader.Error = errT
			decoder.buffer.Write(data)

		case methodType:
			if len(data) != 1 {
				return parsed, closingtag, fmt.Errorf("error with methodtype")
			}
			mType := int(data[0])
			msgHeader.Method = mType
			decoder.buffer.Write(data)

		case timestamp:
			if len(data) != 8 {
				return parsed, closingtag, fmt.Errorf("error with timestamp")
			}
			tstmp := binary.BigEndian.Uint64(data)
			msgHeader.Timestamp = time.Unix(int64(tstmp), 0).UTC()
			decoder.buffer.Write(data)
		case timeout:

			if len(data) != 4 {
				return parsed, closingtag, fmt.Errorf("error with timeout")
			}
			lifetime := int(binary.BigEndian.Uint32(data))
			msgHeader.Timeout = time.Duration(lifetime) * time.Second
			decoder.buffer.Write(data)
		case domain:

			if len(data) > 32 {
				return parsed, closingtag, fmt.Errorf("error with domain")
			}
			dom := string(data)
			msgHeader.Domain = dom
			decoder.buffer.Write(data)

		case endpoint:

			if len(data) > 32 {
				return parsed, closingtag, fmt.Errorf("error with endpoint")
			}
			ep := string(data)
			msgHeader.Endpoint = ep
			decoder.buffer.Write(data)

		case hasAuth:

			if len(data) != 1 {
				return parsed, closingtag, fmt.Errorf("error with hasAuth")
			}
			authFlag := int(data[0])
			if auth == 0 {
				decoder.step += 1
			}
			msgHeader.HasAuth = authFlag == 1
			decoder.buffer.Write(data)

		case auth:
			if len(data) > 4 {
				return parsed, closingtag, fmt.Errorf("error with authkey")
			}
			auth := string(data)
			msgHeader.Auth = auth
			decoder.buffer.Write(data)

		case haspayload:
			if len(data) != 1 {
				return parsed, closingtag, fmt.Errorf("error with hasPayload")
			}
			payloadflag := int(data[0])
			if payloadflag == 0 {
				decoder.step += 2
			}
			msgHeader.HasPayload = payloadflag == 1
			decoder.buffer.Write(data)

		case payloadtype:
			if len(data) != 1 {
				return parsed, closingtag, fmt.Errorf("error with payloadtype")
			}
			ptype := int(data[0])
			msgHeader.PayloadType = ptype
			decoder.buffer.Write(data)

		case payloadlength:
			if len(data) != 8 {
				return parsed, closingtag, fmt.Errorf("error with payload")
			}
			plsize := binary.BigEndian.Uint64(data)
			msgHeader.PayloadSize = int(plsize)
			decoder.buffer.Write(data)

		case checksum:
			if len(data) != 4 {
				return parsed, closingtag, fmt.Errorf("error with checksum")
			}
			checksum := binary.BigEndian.Uint32(data)
			valid := ValidateCRC32Checksum(decoder.buffer.Bytes(), checksum)
			if !valid {
				return parsed, closingtag, fmt.Errorf(" invalid msg frame")
			}
			msgHeader.Checksum = int(checksum)

		default:
			fmt.Println("undefined")
		}
		decoder.step += 1

		return parsed, closingtag, nil
	}

	return 0, closingtag, nil
}
