package p2p

import (
	"fmt"
	"net"
	"time"

	"github.com/WhileCodingDoLearn/my_df_system/msg"
	handler "github.com/WhileCodingDoLearn/my_df_system/msghandler"
)

func StartClient(serverAddr string) {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Fehler beim Verbinden:", err)
		return
	}
	defer conn.Close()

	// Nachricht erstellen und senden

	msgH := msg.MsgHeader{
		Version:    1,
		MsgType:    handler.IdxPING,
		Method:     handler.IdxFETCH,
		Timestamp:  int(time.Now().Unix()),
		Timeout:    3 * time.Second,
		Domain:     "MyDomain",
		Endpoint:   "1",
		HasAuth:    false,
		HasPayload: false,
	}
	data, err := msg.EncodeMsgHeader(msgH)
	if err != nil {
		fmt.Println(err)
	}
	conn.Write(data)
	fmt.Println("Nachricht gesendet:", msgH)

	conn.SetReadDeadline(time.Now().Add(3 * time.Second))
	answer, err := msg.DecodeMsgHeader(conn)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(answer)
}
