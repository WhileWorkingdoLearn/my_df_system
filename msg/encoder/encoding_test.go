package msg

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStringParser(t *testing.T) {
	/*
		// Test 1 : missing header
		teststring1 := "adwdadwad:dwadwadwad"
		testEncoder1 := NewEncoder()
		done, err := testEncoder1.parseString(teststring1)
		require.Error(t, err)
		assert.False(t, done)
		assert.Equal(t, Err, testEncoder1.state)
		assert.Equal(t, fmt.Errorf("no header closure provided"), testEncoder1.err)

		// Test 2 : missing header
		teststring2 := msg.HeaderEnd + "adwdadwad:dwadwadwad"
		testEncoder2 := NewEncoder()
		done, err = testEncoder2.parseString(teststring2)
		require.Error(t, err)
		assert.False(t, done)
		assert.Equal(t, Err, testEncoder2.state)
		assert.Equal(t, fmt.Errorf("no header provided"), testEncoder2.err)

		// Test 3 : missing message closure
		teststring3 := "adwdadwad:-:1233445567"
		testEncoder3 := NewEncoder()
		done, err = testEncoder3.parseString(teststring3)
		require.Error(t, err)
		assert.False(t, done)
		assert.Equal(t, Err, testEncoder3.state)
		assert.Equal(t, fmt.Errorf("no message closure provided"), testEncoder3.err)
	*/
	teststring4 := "MsgType:SenderId:Key:TimeStamp:Version:-:BodyId:BodyPos:PrevMsg:NextMsg:BdoyLength:Body:|:"
	testEncoder4 := NewEncoder()
	data, err := testEncoder4.EncodeString(teststring4)
	require.Nil(t, err)
	assert.Equal(t, "adwdadwad94321", string(data))

}
