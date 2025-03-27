package p2p

import (
	"fmt"
	"net"
	"time"
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

}
