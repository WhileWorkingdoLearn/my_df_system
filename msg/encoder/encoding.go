package msg

import (
	"io"

	parser "github.com/WhileCodingDoLearn/my_df_system/msg/protocol"
)

func NewEncoder(r io.Reader) Encoder {
	header := parser.NewHeaderParser()

	body := parser.NewBodyParser()
	return Encoder{headerParser: header.Parse, bodyParser: body.Parse}

}

func (enc *Encoder) EncodeMsg(msg parser.Message) ([]byte, error) {
	err := enc.headerParser(enc.reader,func(data []byte, idx int) error {
		return nil
	})
	if err != nil {
		return nil, err
	}

	err = enc.bodyParser(enc.reader, func(data []byte, idx int) error {
		return nil
	})

	return nil, err
}

func (enc *Encoder) DecodeMsgFromtring(data string) ([]byte, error) {
	err := enc.headerParser(enc.reader, func(data []byte, idx int) error {
		return nil
	})
	if err != nil {
		return nil, err
	}

	err = enc.bodyParser(enc.reader, func(data []byte, idx int) error {
		return nil
	})

	return nil, err
}
