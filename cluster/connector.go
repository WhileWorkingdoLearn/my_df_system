package cluster

import (
	"fmt"
)

type dns struct {
	nodes map[string]Node
}

var socket map[string]Node

func init() {
	socket = make(map[string]Node)

	n := NewNode("localhost:3000")
	socket[":3000"] = n

}

func NewServer() *dns {
	return &dns{nodes: socket}
}

func (c *dns) ConnectTo(port string) (Connection, error) {
	node, found := c.nodes[port]
	if !found {
		return Connection{}, fmt.Errorf("id not found")
	}
	node.Start()
	newConn := Connection{
		adress: node.adress,
		To:     node.send,
		From:   node.listen,
	}

	return newConn, nil
}

type Connection struct {
	adress string
	To     chan<- RCP
	From   <-chan RCP
}

func (c Connection) GetAdress() string {
	return c.adress
}

func (c Connection) Listen() {
	go func() {
		for {
			n := <-c.From
			fmt.Println(n)
		}
	}()
}

func (c Connection) Send(msg RCP) {
	c.To <- msg
}
