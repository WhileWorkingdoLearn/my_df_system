package msg

import (
	"bytes"
	"fmt"
	"io"

	"github.com/WhileCodingDoLearn/my_df_system/msg/protocol"
)

func NewDecoder(reader io.Reader) *Decoder {
	return &Decoder{reader: reader}
}

func (decoder *Decoder) Decode(msg *msg.Message) error {
	//readHeader
	err := decoder.parse(decoder.reader, ReadHeader, decoder.headerparser.parseHeader)
	if err != nil {
		decoder.state = ErrorHeader
		decoder.err = err
		return err
	}

	bodyparser, found := decoder.persers[decoder.headerparser.header.MsgType[0]]
	if !found {
		return fmt.Errorf("no parser vor version found: %v", string(decoder.headerparser.header.MsgType[0]))
	}
	err = decoder.parse(decoder.reader, ReadBody, bodyparser)
	if err != nil {
		decoder.state = ErrorBody
		decoder.err = err
		return err
	}

	return nil

}

func (decoder *Decoder) parse(r io.Reader, state DecoderState, parser ParseFunc) error {
	buffer := make([]byte, 8, 8) // Anfangsgröße des Buffers
	read := 0

	for decoder.state == state {
		if read >= len(buffer) {
			nBuff := make([]byte, len(buffer)*2, len(buffer)*2)
			copy(nBuff, buffer)
			buffer = nBuff
		}

		n, err := r.Read(buffer[read:])
		if err != nil {
			if err == io.EOF {
				decoder.state = ReadDone
			} else {
				fmt.Println(err)
				return err
			}
		}
		read += n

		parsed, done, err := parser.Handle(buffer[:read])

		if err != nil {
			decoder.state = ReadDone
			return fmt.Errorf("parsing error: %w", err)
		}

		if done {
			decoder.state = ReadDone
			return nil
		}

		fmt.Println("buff before:", len(buffer))
		if parsed > 0 {
			newbuff := make([]byte, read-parsed, read-parsed)
			copy(newbuff, buffer[parsed:read])
			buffer = newbuff
			read -= parsed
		}
		fmt.Println("buff after:", len(buffer))
	}

	return nil
}

func NewHeaderParser() *HeaderParser {
	return &HeaderParser{state: MsgType, header: &msg.Header{}}
}

func (hp *HeaderParser) parseHeader(data []byte) (n int, done bool, err error) {
	dataToParse := bytes.Index(data, []byte(msg.Sep))
	if dataToParse == -1 {
		return 0, false, nil
	}

	end := bytes.Index(data, []byte("-:"))
	if end == 0 {
		hp.state = HeaderDone
		return 2, true, nil
	}

	if dataToParse == 0 {
		hp.state++
		return 1, false, nil
	}

	data = data[:dataToParse]

	err = hp.WriteHeader(data)
	if err != nil {
		hp.state = HeaderDone
		return 0, false, err
	}

	hp.state++

	return dataToParse + 1, false, nil
}

func (hp *HeaderParser) WriteHeader(data []byte) error {

	switch hp.state {
	case MsgType:
		if len(data) != int(msg.Onebyte) {
			return fmt.Errorf("malformed header data %v, want : %v byte", data, msg.Onebyte)
		}
		hp.header.MsgType = data
	case SenderId:
		if len(data) != int(msg.Onebyte) {
			return fmt.Errorf("malformed header data %v, want : %v byte", data, msg.Onebyte)
		}
		hp.header.SenderId = data
	case Key:
		if len(data) != int(msg.Onebyte) {
			return fmt.Errorf("malformed header data %v, want : %v byte", data, msg.Onebyte)
		}
		hp.header.Key = data
	case TimeStamp:
		if len(data) != int(msg.Onebyte) {
			return fmt.Errorf("malformed header data %v, want : %v byte", data, msg.Onebyte)
		}
		hp.header.TimeStamp = data
	case Version:
		if len(data) != int(msg.Onebyte) {
			return fmt.Errorf("malformed header data %v, want : %v byte", data, msg.Onebyte)
		}
		hp.header.Version = data
	default:
		hp.state = HeaderDone
		return fmt.Errorf("unkwon parser state %v", hp.state)
	}
	return nil
}

func NewBodyParser() *BodyParser {
	return &BodyParser{state: PackedId}
}

func (bp *BodyParser) parseBody(data []byte) (n int, done bool, err error) {
	dataToParse := bytes.Index(data, []byte(msg.Sep))
	if dataToParse == -1 {
		fmt.Println("Break no : found")
		return 0, false, nil
	}

	end := bytes.Index(data, []byte("-:"))
	if end == 0 {
		bp.state = PacketDone
		return 2, true, nil
	}

	if dataToParse == 0 {
		fmt.Println("Break found : at 0")
		bp.state++
		return 1, false, nil
	}

	data = data[:dataToParse]
	err = bp.WriteBody(data)

	if err != nil {
		bp.state = PacketDone
		return 0, false, err
	}

	bp.state++

	return dataToParse + 1, false, nil
}

func (bp *BodyParser) WriteBody(data []byte) error {
	switch bp.state {
	case PackedId:
		bp.body.PackedId = data
	case PacketPos:
		bp.body.PacketPos = data
	case PrevPacked:
		bp.body.PrevPacked = data
	case NextPacket:
		bp.body.NextPacket = data
	case PacketLength:
		bp.body.PacketLength = data
	case Packet:
		bp.body.Packet = data
	default:
		return fmt.Errorf("unknown parser state %v", bp.state)
	}
	return nil
}
