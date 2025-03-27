package msg

type Encoder struct {
}

func (e *Encoder) EncodeString(data string) ([]byte, error) {
	return nil, nil
}

func (e *Encoder) parseString(data string) (int, bool, error) {
	return 0, false, nil
}

func (e *Encoder) Encode(data Message) ([]byte, error) {
	return nil, nil
}
