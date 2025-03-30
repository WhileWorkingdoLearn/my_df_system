package msg

import "io"

type Encoder struct {
	headerParser Parser
	bodyParser   Parser
}

func NewEncoder(r io.Reader) Encoder {
	header := Parser{
		reader:      r,
		buffersize:  128,
		state:       initial,
		fixedBuffer: false,
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
	return Encoder{headerParser: header, bodyParser: body}

}

func (dec *Decoder) EncodeMsg(msg Message) ([]byte, error) {
	err := dec.headerParser.Parse(func(data []byte, idx int) error {
		return nil
	})
	if err != nil {
		return nil, err
	}

	err = dec.bodyParser.Parse(func(data []byte, idx int) error {
		return nil
	})

	return nil, err
}

func (dec *Decoder) DecodeMsgFromtring(data string) ([]byte, error) {
	err := dec.headerParser.Parse(func(data []byte, idx int) error {
		return nil
	})
	if err != nil {
		return nil, err
	}

	err = dec.bodyParser.Parse(func(data []byte, idx int) error {
		return nil
	})

	return nil, err
}
