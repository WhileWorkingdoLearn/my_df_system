package msg

import (
	"fmt"
	"io"

	msg "github.com/WhileCodingDoLearn/my_df_system/msg/protocol"
)

func NewDecoder(r io.Reader) Decoder {
	header := msg.NewHeaderParser()

	body := msg.NewBodyParser()

	return Decoder{reader: r, headerParser: header.Parse, bodyParser: body.Parse}

}

func WriteHeader(data []byte, idx HeaderParserPosition, header *msg.Header) error {
	if header == nil {
		return fmt.Errorf("no header provided")
	}
	switch idx {
	case MsgType:
		if len(data) != 1 {
			return fmt.Errorf("malformed header data %v, want : %v byte", data, msg.Onebyte)
		}
		header.MsgType = data
	case SenderId:
		if len(data) != int(msg.Onebyte) {
			return fmt.Errorf("malformed header data %v, want : %v byte", data, msg.Onebyte)
		}
		header.SenderId = data
	case Key:
		if len(data) != int(msg.Onebyte) {
			return fmt.Errorf("malformed header data %v, want : %v byte", data, msg.Onebyte)
		}
		header.Key = data
	case TimeStamp:
		if len(data) != int(msg.Onebyte) {
			return fmt.Errorf("malformed header data %v, want : %v byte", data, msg.Onebyte)
		}
		header.TimeStamp = data
	case Version:
		if len(data) != int(msg.Onebyte) {
			return fmt.Errorf("malformed header data %v, want : %v byte", data, msg.Onebyte)
		}
		header.Version = data
	default:

		return fmt.Errorf("unkwon parser state %v", idx)
	}
	return nil
}

func (dec *Decoder) DecodeMsg(msg *msg.Message) error {
	err := dec.DecodeHeader(&msg.Header)
	if err != nil {
		return err
	}

	err = dec.bodyParser(dec.reader, func(data []byte, idx int) error {
		return nil
	})

	return err
}

func (dec *Decoder) DecodeMsgToString() (string, error) {
	err := dec.headerParser(dec.reader, func(data []byte, idx int) error {
		return nil
	})
	if err != nil {
		return "", err
	}

	err = dec.bodyParser(dec.reader, func(data []byte, idx int) error {
		return nil
	})

	return "", err
}

func (dec *Decoder) DecodeHeader(msgheader *msg.Header) error {
	err := dec.headerParser(dec.reader, func(data []byte, idx int) error {
		return WriteHeader(data, HeaderParserPosition(idx), msgheader)
	})
	return err
}

func (dec *Decoder) DecodeHeaderToString() (string, error) {
	return "", nil
}

func (dec *Decoder) DecodeBody(msgheader *msg.Header) error {
	err := dec.bodyParser(dec.reader, func(data []byte, idx int) error {
		return nil
	})
	return err
}

func (dec *Decoder) DecodeBodyToString() (string, error) {
	err := dec.bodyParser(dec.reader, func(data []byte, idx int) error {
		return nil
	})
	return "", err
}
