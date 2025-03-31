package msg

import (
	"bytes"
	"testing"
)

func TestHeaderParser(t *testing.T) {

	buff := bytes.NewBuffer(make([]byte, 0))
	buff.WriteByte(10)
	buff.WriteRune(':')
	buff.WriteString("SenderId")
	buff.WriteRune(':')
	buff.WriteString("Key")
	buff.WriteRune(':')
	buff.WriteString("TimeStamp")
	buff.WriteRune(':')
	buff.WriteString("Version")
	buff.Write([]byte(":-:"))
	/*
		testb := buff.Bytes()
		hp := NewHeaderParser()
		hp.Parse(bytes.NewReader(testb), func(data []byte, idx int) error {

		})
	*/
}
