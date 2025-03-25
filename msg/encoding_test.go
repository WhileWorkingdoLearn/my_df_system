package msg

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncoding(t *testing.T) {
	msgHheader := MsgHeader{
		Version:     1,
		MsgType:     2,
		Method:      3,
		Timestamp:   time.Now().UTC(),
		Timeout:     time.Duration(3) * time.Second,
		Domain:      "blabal",
		Endpoint:    "v1\\dv",
		HasAuth:     false,
		Auth:        "mykey",
		HasPayload:  false,
		PayloadType: 1,
		PayloadSize: 10,
	}
	data, err := EncodeMsgHeader(msgHheader)
	require.Nil(t, err)
	assert.Equal(t, len(data), 102)
	assert.Equal(t, len(data), 102)
	assert.Equal(t, data[1], byte(':'))
	assert.Equal(t, data[len(data)-3], byte(':'))
	assert.Equal(t, data[len(data)-2], byte('-'))
	assert.Equal(t, data[len(data)-1], byte(':'))
}

func TestEncodingFromString(t *testing.T) {
	data := "1:1:2:-:"
	p := parser{}
	parsed, err := p.parse(data)
	assert.NoError(t, err)
	assert.NotNil(t, parsed)
	assert.Equal(t, 1, parsed)
	data1 := "1:1:2:-:"
	bdata, err := EncodeMsgHeaderFromString(data1)
	require.NoError(t, err)
	assert.Equal(t, bdata[0], byte(1))
	assert.Equal(t, bdata[1], byte(1))
	assert.Equal(t, bdata[2], byte(1))
}
