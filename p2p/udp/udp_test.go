package main

/*
func NewParser() *Parser {
	buff := bytes.NewBuffer(make([]byte, 0))
	p := NewParseHandler(DefaultConfig())
	return &Parser{buff: buff, parsehandler: p, state: 0, Done: false}
}



const (
	ParsingInitialized int = iota
	StreamData
	ParsingErr
)

func (p *Parser) parseByState(data []byte) {
	switch p.state {
	case ParsingInitialized:
		{
			idx := uint8(data[0])
			fmt.Println(idx)
			val, found := p.parsehandler[idx]
			if !found {
				p.state = ParsingErr
				return
			}
			n, _ := val(data)
			fmt.Println(n)
			p.state = n
		}
	case StreamData:
		{
		}
	case 10:
	default:
		fmt.Println("unknown type: ", string(data))
	}
}
func TestParseByState(t *testing.T) {

	//Test 1 : Parsing unkown number
	p := NewParser()
	p.parseByState([]byte{10})
	assert.Equal(t, ParsingErr, p.state)

	//Test 1 : Parsing acknowlege succesfully
	p = NewParser()
	p.parseByState([]byte{1})
	assert.Equal(t, StreamData, p.state)
}

type ParseHandler map[uint8]ACKFunc

type ACKFunc func(data []byte) (int, error)

func (f ACKFunc) Handle(data []byte) {
	num, err := f(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v,%v", string(data), num)
}

type Config struct {
	ACK   ACKFunc
	DATA  ACKFunc
	CLOSE ACKFunc
	ERR   ACKFunc
}

func DefaultConfig() Config {
	return Config{
		ACK:   func(data []byte) (int, error) { return StreamData, nil },
		DATA:  func(data []byte) (int, error) { return 1, nil },
		CLOSE: func(data []byte) (int, error) { return 2, nil },
		ERR:   func(data []byte) (int, error) { return 3, nil },
	}
}

func NewParseHandler(cfg Config) ParseHandler {
	parsehandler := make(ParseHandler)
	parsehandler[ACK] = cfg.ACK
	parsehandler[DATA] = cfg.DATA
	parsehandler[CLOSE] = cfg.CLOSE
	parsehandler[ERR] = cfg.ERR
	return parsehandler
}

func TestParserHandler(t *testing.T) {
	dfCFG := DefaultConfig()
	pg := NewParseHandler(dfCFG)

	parser, found := pg[0]
	require.True(t, found)
	parser.Handle([]byte("Hello World"))

	assert.Nil(t, parser)
}

func TestParser(t *testing.T) {
	str := "abcd:efg::de:-awda"
	//Test 1: First section of data. No Seperator or Delimiter
	p := NewParser()
	parsed := 0
	n, done, err := p.parse([]byte(str[:4]))
	require.Nil(t, err)
	assert.False(t, done)
	assert.Equal(t, 0, n)

	parsed += n
	//Test 2: First section of data. Seperator. No Delimiter
	n, done, err = p.parse([]byte(str[parsed:5]))
	require.Nil(t, err)
	assert.False(t, done)
	assert.Equal(t, 5, n)
	parsed += n
	//Test 3: Second section of data. Seperator. No delimiter
	n, done, err = p.parse([]byte(str[parsed:10]))
	require.Nil(t, err)
	assert.False(t, done)
	assert.Equal(t, 4, n)
	parsed += n
	//Test 3: Second section of data. Seperator. No delimiter
	n, done, err = p.parse([]byte(str[parsed:12]))
	require.Nil(t, err)
	assert.False(t, done)
	assert.Equal(t, 1, n)
	parsed += n
	//Test 4: Second section of data. Seperator. Delimiter
	n, done, err = p.parse([]byte(str[parsed : len(str)-len("awda")]))
	require.Nil(t, err)
	assert.False(t, done)
	assert.Equal(t, 3, n)

	//End of Data
	n, done, err = p.parse([]byte(str[len(str)-len("-awda"):]))
	require.Nil(t, err)
	assert.True(t, done)
	assert.Equal(t, 2, n)

}
*/
