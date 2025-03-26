package main

import (
	"bytes"
	"io"
)

const (
	ACK uint8 = iota
	OPEN
	DATA
	CLOSE
	ERR
)

type ParseHandler map[uint8]Parser

type DataParser struct {
	parser  ParseHandler
	msgtype int
	done    bool
}

func (p *DataParser) Parse(r io.Reader) {
	for {

	}
}

type Parser interface {
	Parse(data []byte) (n int, done bool, err error)
	Done() bool
}

type MsgParser struct {
	step int
	done bool
}

type AckParser struct {
	step int
	done bool
}

func (p *AckParser) Parse(data []byte) (n int, done bool, err error) {
	for !p.done {
		clrf := bytes.Index(data, []byte("-"))
		if clrf == 0 {
			return 1, true, nil
		}
		if clrf > 0 {
			data = data[:clrf]
		}
		toParse := bytes.Index(data, []byte(":"))

		if toParse == -1 {
			return 0, false, nil
		}

		if toParse == 0 {
			return 1, false, nil
		}

	}
	return 0, true, nil
}

type OpenParser struct {
}

func (p *OpenParser) Parse(data []byte) (n int, done bool, err error) {
	return 0, true, nil
}

type StreamParser struct {
}

func (p *StreamParser) Parse(data []byte) (n int, done bool, err error) {
	return 0, true, nil
}

type CloseParser struct {
}

func (p *CloseParser) Parse(data []byte) (n int, done bool, err error) {
	return 0, true, nil
}

type ErrorParser struct {
}

func (p *ErrorParser) Parse(data []byte) (n int, done bool, err error) {
	return 0, true, nil
}
