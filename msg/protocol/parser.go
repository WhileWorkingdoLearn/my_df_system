package msg

import (
	"bytes"
	"fmt"
	"io"
)

type ParserState int

type Iterator func(data []byte, idx int) error

const (
	done    ParserState = -1
	initial ParserState = iota
	parsing
	parseErr
)

type Config struct {
	Buffersize  int
	FixedBuffer bool
	Separator   string
	End         string
}

func DefaultConfig() Config {
	return Config{
		Buffersize:  128,
		FixedBuffer: false,
		Separator:   string(Sep),
		End:         string(MSgEnd),
	}
}

type Parser struct {
	reader      io.Reader
	buffersize  int
	fixedBuffer bool
	state       ParserState
	stm         *stateMachine
	separator   string
	end         string
	err         error
}

func NewParser(cfg Config) *Parser {
	return &Parser{
		buffersize:  cfg.Buffersize,
		fixedBuffer: cfg.FixedBuffer,
		separator:   cfg.Separator,
		end:         cfg.End,
		stm:         &stateMachine{state: initialstate},
	}
}

func (parser *Parser) Parse(setter Iterator) error {
	buffer := make([]byte, parser.buffersize, parser.buffersize)
	read := 0

	for parser.state == parsing {
		if !parser.fixedBuffer && read >= len(buffer) {
			nBuff := make([]byte, len(buffer)*2, len(buffer)*2)
			copy(nBuff, buffer)
			buffer = nBuff
		}

		n, err := parser.reader.Read(buffer[read:])
		if err != nil {
			if err == io.EOF {
				parser.state = done
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

		if !parser.fixedBuffer && parsed > 0 {
			newbuff := make([]byte, read-parsed, read-parsed)
			copy(newbuff, buffer[parsed:read])
			buffer = newbuff
			read -= parsed
		}
	}

	return nil
}

func (parser *Parser) parse(data []byte, setter Iterator) (int, bool, error) {
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

type stateMachine struct {
	state stateValues
}

type stateValues int

const (
	end          stateValues = -1
	initialstate stateValues = iota - 1
)

func (sm *stateMachine) Pos() int {
	return int(sm.state)
}

func (sm *stateMachine) Next() {
	sm.state++
}

func (sm *stateMachine) Set(state stateValues) error {
	if state < 0 {
		return fmt.Errorf("no negative state allowed: ", state)
	}
	sm.state = state
	return nil
}

func (sm *stateMachine) Reset() {
	sm.state = 0
}

func (sm *stateMachine) Done() {
	sm.state = end
}
