package p2p

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/WhileCodingDoLearn/my_df_system/msg"
)

type Client struct {
	conn net.Conn
}

func NewClient(serverAddr string) *Client {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Fehler beim Verbinden:", err)
		return nil
	}

	conn.SetReadDeadline(time.Now().Add(3 * time.Second))

	return &Client{
		conn: conn,
	}
}

func (c *Client) SendMsg(msg []byte) error {
	_, err := c.conn.Write(msg)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Read(size int) ([]byte, error) {
	buff := make([]byte, size, size)
	n, err := c.conn.Read(buff)
	if err != nil {
		return buff[:n], err
	}
	return buff[:n], nil
}

func (c *Client) Close() {
	c.conn.Close()
}

func main() {
	// Nachricht erstellen und senden
	client := NewClient(":3000")

	msgH := msg.MsgHeader{
		Version:    1,
		MsgType:    msg.IdxPING,
		Method:     msg.IdxFETCH,
		Timestamp:  time.Now().UTC(),
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

	client.SendMsg(data)
	var msgHeader msg.MsgHeader
	err = msg.NewDecoder(client.conn).Decode(&msgHeader)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msgHeader)

	client.Close()

}
