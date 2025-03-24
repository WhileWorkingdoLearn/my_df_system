package msg

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Parse(buff []byte) (int, bool) {
	sep := bytes.Index(buff, []byte{':'})
	end := false
	if sep != -1 {
		if sep < len(buff)-2 && buff[sep+1] == byte('-') && buff[sep+2] == byte(':') {
			end = true
		}
		return sep, end
	}

	return 0, end
}

func TestParse(t *testing.T) {
	//Testing closing paragraphs
	b := []byte("aabb:-:c")
	parsed, end := Parse(b)
	assert.Equal(t, 4, parsed)
	assert.True(t, end)

	//Testing seperator
	b = []byte("aabb:c")
	parsed, end = Parse(b)
	assert.Equal(t, 4, parsed)
	assert.False(t, end)

	// Testing nor included
	b = []byte("aabbc")
	parsed, end = Parse(b)
	assert.Equal(t, 0, parsed)
	assert.False(t, end)

	//testing two seperator
	b = []byte("a:b:bc")
	parsed, end = Parse(b)
	assert.Equal(t, 1, parsed)
	assert.False(t, end)
}

func TestDecode(t *testing.T) {
	msgHheader := MsgHeader{
		Version:     1,
		MsgType:     IdxPING,
		Error:       ErrDomain,
		Method:      IdxFETCH,
		Timestamp:   time.Now().UTC(),
		Timeout:     time.Duration(3) * time.Second,
		Domain:      "blabal",
		Endpoint:    "v1dv",
		HasAuth:     false,
		Auth:        "",
		HasPayload:  false,
		PayloadType: 0,
		PayloadSize: 0,
	}

	encoded, err := EncodeMsgHeader(msgHheader)
	if err != nil {
		t.Fatal(err)
	}

	r := bytes.NewReader(encoded)
	var msgH MsgHeader
	dec := NewDecoder(r)
	errDec := dec.Decode(&msgH)
	require.Nil(t, errDec)
	assert.Equal(t, 89, dec.Parsed())
}
