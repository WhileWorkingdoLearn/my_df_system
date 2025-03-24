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
