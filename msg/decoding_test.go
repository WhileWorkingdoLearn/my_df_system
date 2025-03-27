package msg

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

	testb := buff.Bytes()
	hp := NewHeaderParser()

	parsed := 0
	n, done, err := hp.parseHeader(testb)
	fmt.Println(n)
	require.Nil(t, err)
	assert.False(t, done)
	assert.Equal(t, 2, n)

	parsed += n
	n, done, err = hp.parseHeader(testb[parsed:])

	require.Nil(t, err)
	assert.False(t, done)
	assert.Equal(t, "SenderId", string(hp.header.SenderId))
	parsed += n

	n, done, err = hp.parseHeader(testb[parsed:])

	fmt.Println(hp.header)
	require.Nil(t, err)
	assert.False(t, done)
	assert.Equal(t, "Key", string(hp.header.Key))
	parsed += n

	n, done, err = hp.parseHeader(testb[parsed:])

	fmt.Println(hp.header)
	require.Nil(t, err)
	assert.False(t, done)
	assert.Equal(t, "TimeStamp", string(hp.header.TimeStamp))
	parsed += n

	n, done, err = hp.parseHeader(testb[parsed:])

	fmt.Println(hp.header)
	require.Nil(t, err)
	assert.False(t, done)
	assert.Equal(t, "Version", string(hp.header.Version))
	parsed += n

	n, done, err = hp.parseHeader(testb[parsed:])
	fmt.Println(hp.header)
	require.Nil(t, err)
	assert.True(t, done)
	assert.Equal(t, 2, n)

}
