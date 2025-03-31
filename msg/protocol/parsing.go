package msg

import (
	"bytes"
	"fmt"
	"io"
)

func DefaultConfig() Config {
	return Config{
		Buffersize:  128,
		FixedBuffer: false,
		Separator:   string(Sep),
		End:         string(MSgEnd),
	}
}

func NewHeaderParser() *Parser {
	return &Parser{buffersize: 128,
		fixedBuffer: false,
		separator:   string(Sep),
		end:         string(HeaderEnd),
		stm:         &state{state: initialstate},
		state:       initial}
}

func NewBodyParser() *Parser {
	return &Parser{buffersize: 128,
		fixedBuffer: false,
		separator:   string(Sep),
		end:         string(MSgEnd),
		stm:         &state{state: initialstate},
	}
}

func NewParser(cfg Config) *Parser {
	return &Parser{
		buffersize:  cfg.Buffersize,
		fixedBuffer: cfg.FixedBuffer,
		separator:   cfg.Separator,
		end:         cfg.End,
		stm:         &state{state: initialstate},
	}
}

func (parser *Parser) Parse(reader io.Reader, setter Handler) error {
	buffer := make([]byte, parser.buffersize, parser.buffersize)
	read := 0
	parser.state = parsing
	for parser.state == parsing {
		if !parser.fixedBuffer && read >= len(buffer) {
			nBuff := make([]byte, len(buffer)*2, len(buffer)*2)
			copy(nBuff, buffer)
			buffer = nBuff
		}

		n, err := reader.Read(buffer[read:])
		if err != nil {
			if err == io.EOF {
				
			} else {
				parser.state = parseErr
				parser.err = err
				return err
			}
		}
		read += n

		parsed, isDone, err := parser.parse(buffer[:read], setter)
		if err != nil {
			parser.state = parseErr
			parser.err = err
			return err
		}

		if isDone {
			parser.state = done
			return nil
		}

		if parsed > 0 && !parser.fixedBuffer {
			newbuff := make([]byte, read-parsed, read-parsed)
			copy(newbuff, buffer[parsed:read])
			buffer = newbuff
			read -= parsed
		}

	}

	return nil
}

func (parser *Parser) Stop() {
	parser.state = done
}

func (parser *Parser) parse(data []byte, setter Handler) (int, bool, error) {
	dataToParse := bytes.Index(data, []byte(parser.separator))
	if dataToParse == -1 {
		return 0, false, nil
	}

	end := bytes.Index(data, []byte(parser.end))
	if end == 0 {
		parser.stm.Done()
		return 2, true, nil
	}

	if dataToParse == 0 {
		parser.stm.Next()
		return 1, false, nil
	}

	err := setter(data[:dataToParse], parser.stm.Pos())
	if err != nil {
		parser.stm.Done()
		return 0, false, err
	}

	parser.stm.Next()
	return dataToParse + 1, false, nil

}

func (sm *state) Pos() int {
	return int(sm.state)
}

func (sm *state) Next() {
	sm.state++
}

func (sm *state) Set(state stateValues) error {
	if state < 0 {
		return fmt.Errorf("no negative state allowed: ", state)
	}
	sm.state = state
	return nil
}

func (sm *state) Reset() {
	sm.state = 0
}

func (sm *state) Done() {
	sm.state = end
}
