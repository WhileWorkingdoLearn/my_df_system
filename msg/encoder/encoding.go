package msg

import (
	"bytes"
	"fmt"
	"strings"

	protocol "github.com/WhileCodingDoLearn/my_df_system/msg/protocol"
)

func NewEncoder() *Encoder {
	return &Encoder{state: EncodeHeader, buffer: bytes.NewBufferString("")}
}

func (e *Encoder) EncodeString(data string) ([]byte, error) {
	headerEnd := strings.Index(data, string(protocol.HeaderEnd))
	if headerEnd == -1 {
		return nil, fmt.Errorf("no header provided")
	}
	end := strings.Index(data, string(protocol.MSgEnd))
	if end == -1 {
		return nil, fmt.Errorf("no message end provided")
	}
	/// Parse Header
	readHeader := 0
	e.compile = InitHeader
	for e.state == EncodeHeader {
		parsed, done, err := e.parseString(data[readHeader:], protocol.HeaderEnd, e.WriteHeader)
		if err != nil {
			e.state = Err
			return nil, err
		}
		readHeader += parsed
		if done {
			e.state = EncodeBody
		}
	}

	e.compile = InitBody
	readBody := readHeader
	for e.state == EncodeBody {
		parsed, done, err := e.parseString(data[readBody:], protocol.MSgEnd, e.WriteBody)
		if err != nil {
			e.state = Err
			return nil, err
		}
		readBody += parsed
		if done {
			e.state = EncodeDone
		}
	}

	return e.buffer.Bytes(), nil
}

func (e *Encoder) parseString(data string, limiter protocol.ProtocolToken, encode EncodeType) (int, bool, error) {

	dataToParse := strings.Index(data, string(protocol.Sep))
	if dataToParse == -1 {
		return 0, false, nil
	}

	end := strings.Index(data, string(limiter))
	if end == 0 {
		return 2, true, nil
	}

	if dataToParse == 0 {
		e.compile++
		return 1, false, nil
	}
	e.compile++
	data = data[:dataToParse]
	encode(data)
	return dataToParse + 1, false, nil
}

func (e *Encoder) WriteHeader(data string) error {
	switch e.compile {
	case MsgType:
		{
			fmt.Print(data + ":")
		}
	case SenderId:
		{
			fmt.Print(data + ":")
		}
	case Key:
		{
			fmt.Print(data + ":")
		}
	case TimeStamp:
		{
			fmt.Print(data + ":")
		}
	case Version:
		{
			fmt.Print(data + ":-:")
		}
	case HeaderDone:
		{
			//fmt.Println(data + "-:")
		}
	default:
		e.state = EncodeDone
	}
	return nil
}

func (e *Encoder) WriteBody(data string) error {
	switch e.compile {
	case BodyId:
		{
			fmt.Print(data + ":")
		}
	case BodyPos:
		{
			fmt.Print(data + ":")
		}
	case PrevMsg:
		{
			fmt.Print(data + ":")
		}
	case NextMsg:
		{
			fmt.Print(data + ":")
		}
	case BdoyLength:
		{
			fmt.Print(data + ":")
		}
	case Body:
		{
			fmt.Println(data + ":|:")
		}
	case BodyDone:
		{

		}
	default:
		e.state = EncodeDone
	}
	return nil
}

func (e *Encoder) Encode(data protocol.Message) ([]byte, error) {
	return nil, nil
}
