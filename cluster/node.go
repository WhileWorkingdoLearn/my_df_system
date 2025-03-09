package cluster

import (
	"fmt"
	"time"
)

type Node struct {
	adress string
	listen <-chan RCP
	send   chan<- RCP
}

func NewNode(port string) Node {
	chListen := make(chan RCP, 1)
	chSend := make(chan RCP, 1)
	return Node{
		adress: port,
		listen: chListen,
		send:   chSend,
	}
}

func (n *Node) Start() {
	msg := <-n.listen
	fmt.Printf("Message for Node %v : msg %v ", n.adress, msg)
	time.Sleep(time.Duration(2) * time.Second)
	n.send <- RCP{From: n.adress, Payload: []byte("Hello Back")}
}
