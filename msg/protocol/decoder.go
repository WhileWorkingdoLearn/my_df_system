package msg

import "io"

type Decoder struct {
	headerParser Parser
	bodyParser   Parser
}

func NewDecoder(r io.Reader) Decoder {
	header := Parser{
		reader:      r,
		buffersize:  128,
		state:       initial,
		fixedBuffer: true,
		stm:         &stateMachine{state: initialstate},
		separator:   string(Sep),
		end:         string(HeaderEnd),
	}

	body := Parser{
		reader:      r,
		buffersize:  128,
		state:       initial,
		fixedBuffer: false,
		stm:         &stateMachine{state: initialstate},
		separator:   string(Sep),
		end:         string(HeaderEnd),
	}
	return Decoder{headerParser: header, bodyParser: body}

}

func (dec *Decoder) DecodeMsg(msg *Message) error {
	err := dec.headerParser.Parse(func(data []byte, idx int) error {
		return nil
	})
	if err != nil {
		return err
	}

	err = dec.bodyParser.Parse(func(data []byte, idx int) error {
		return nil
	})

	return err
}

func (dec *Decoder) DecodeMsgToString() (string, error) {
	err := dec.headerParser.Parse(func(data []byte, idx int) error {
		return nil
	})
	if err != nil {
		return "", err
	}

	err = dec.bodyParser.Parse(func(data []byte, idx int) error {
		return nil
	})

	return "", err
}

func (dec *Decoder) DecodeHeader(msgheader *Header) error {
	err := dec.headerParser.Parse(func(data []byte, idx int) error {
		return nil
	})
	return err
}

func (dec *Decoder) DecodeHeaderToString() (string, error) {
	return "", nil
}

func (dec *Decoder) DecodeBody(msgheader *Header) error {
	err := dec.bodyParser.Parse(func(data []byte, idx int) error {
		return nil
	})
	return err
}

func (dec *Decoder) DecodeBodyToString() (string, error) {
	err := dec.bodyParser.Parse(func(data []byte, idx int) error {
		return nil
	})
	return "", err
}
