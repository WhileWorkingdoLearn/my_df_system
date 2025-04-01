package msg

import (
	"bytes"
	"fmt"
	"io"
	"math/rand/v2"
	"testing"
	"time"

	msg "github.com/WhileCodingDoLearn/my_df_system/msg/protocol"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const Tokens = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateId(length int) []byte {
	key := make([]byte, length)
	for idx := range key {
		key[idx] = Tokens[rand.IntN(len(Tokens))]
	}
	return key
}

/*
	type Header struct {
		MsgType    []byte
		SenderId   []byte
		SessionId []byte
		TimeStamp  []byte
		Version    []byte
	}
*/
func GenerateHeader() io.Reader {
	buff := bytes.NewBuffer(make([]byte, 0))
	//Type
	buff.WriteByte(1)
	buff.WriteByte(':')
	//SenderId
	buff.Write(GenerateId(16))
	buff.WriteByte(':')
	//SessionId
	buff.Write(GenerateId(16))
	buff.WriteByte(':')
	//Timestamp
	timeUnix := time.Now().AddDate(200, 0, 0).Unix()
	fmt.Println(fmt.Sprint(timeUnix))
	buff.WriteString(fmt.Sprint(timeUnix))
	buff.WriteByte(':')
	//Version
	buff.WriteByte(1)
	buff.WriteByte(':')
	buff.WriteString(string(msg.HeaderEnd))
	return buff
}

/*
	ts, err := strconv.Atoi(string(Amsg.Header.TimeStamp))
	if err != nil {
		log.Fatal(err)
	}
*/

func TestDecodeHeader(t *testing.T) {
	dec := NewDecoder(GenerateHeader())
	var Amsg msg.Message
	err := dec.DecodeMsg(&Amsg)
	require.Nil(t, err)
	fmt.Println(Amsg)
	assert.Equal(t, []byte{1}, Amsg.Header.MsgType)
	assert.Equal(t, 16, len(Amsg.Header.SenderId))
}
