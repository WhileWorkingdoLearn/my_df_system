package msg

type ParserState int

type Handler func(data []byte, idx int) error

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

type Parser struct {
	buffersize  int
	fixedBuffer bool
	state       ParserState
	stm         *state
	separator   string
	end         string
	err         error
}

type state struct {
	state stateValues
}

type stateValues int

const (
	end          stateValues = -1
	initialstate stateValues = iota - 1
)
