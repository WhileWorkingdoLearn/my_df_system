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

func WriteHeader(data []byte, idx HeaderDataPosition, header *msg.Header) error {
	if header == nil {
		return fmt.Errorf("no header provided")
	}
	switch idx {
	case MsgType:
		if len(data) != int(msg.Onebyte) {
			return fmt.Errorf("msgtype - malformed header data have:  %v bytes, want : %v bytes", len(data), msg.Onebyte)
		}
		header.MsgType = data
	case SenderId:
		if len(data) != int(msg.SixteenBytes) {
			return fmt.Errorf("senderid - malformed header data have:  %v bytes, want : %v bytes", len(data), msg.SixteenBytes)
		}
		header.SenderId = data
	case SessionId:
		if len(data) != int(msg.SixteenBytes) {
			return fmt.Errorf("sessionKey -malformed header data have:  %v bytes, want : %v bytes", len(data), msg.SixteenBytes)
		}
		header.SessionId = data
	case TimeStamp:
		if len(data) != 10 {
			return fmt.Errorf("timestamp - malformed header data have:  %v bytes, want : %v bytes", len(data), 10)
		}
		header.TimeStamp = data
	case Version:
		if len(data) != int(msg.Onebyte) {
			return fmt.Errorf("version - malformed header data have:  %v bytes, want : %v bytes", len(data), msg.Onebyte)
		}
		header.Version = data
	default:

		return fmt.Errorf("unkwon parser state %v", idx)
	}
	return nil
}

func (dec *Decoder) DecodeMsg(msg *msg.Message) error {
	err := dec.decodeHeader(&msg.Header)
	if err != nil {
		return err
	}

	err = dec.bodyParser(dec.reader, func(data []byte, idx int) error {
		return nil
	})

	return err
}

func (dec *Decoder) decodeHeader(msgheader *msg.Header) error {
	err := dec.headerParser(dec.reader, func(data []byte, idx int) error {
		return WriteHeader(data, HeaderDataPosition(idx), msgheader)
	})
	return err
}

func (dec *Decoder) decodeBody(msgbody *msg.Body) error {
	err := dec.bodyParser(dec.reader, func(data []byte, idx int) error {
		return nil
	})
	return err
}
