package transfer

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func TestEncoding(t *testing.T) {

	Header := MsgHeader{
		Version:     1,
		MsgType:     2,
		Method:      3,
		Timestamp:   int(time.Now().Unix()),
		Timeout:     5 * time.Second,
		Domain:      "example.com",
		Endpoint:    "/api/v1",
		Auth:        "secure_auth",
		PayloadType: 5,
		PayloadSize: 512,
	}

	// Nachricht kodieren
	encodedData, err := EncodeMsg(Header)
	if err != nil {
		fmt.Println("Encoding-Fehler:", err)
		t.Log()
	}

	// Nachricht dekodieren

	decodedMsg, err := DecodeMsgStream(bytes.NewReader(encodedData))
	if err != nil {
		t.Log()
	}

	fmt.Println()
	fmt.Println(decodedMsg)

}
