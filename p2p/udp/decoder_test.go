package main

import (
	"bytes"
	"testing"
)

func TestParser(t *testing.T) {

	buff := bytes.NewBuffer(make([]byte, 1))
	buff.WriteRune(0)
	buff.WriteString(":")
	buff.Write(GenerateKey(32))
	buff.WriteString(":")
	buff.Write(GenerateKey(16))
	buff.WriteString(":")
	ack := AckParser{}
	Parse(buff.Bytes(), &ack)

}
