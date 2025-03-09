package main

import "fmt"

type Node struct {
	id    string
	pings chan string
	pongs chan string
}

type Conne struct {
	Pings chan string
	Pongs chan string
}

func NewNode(c *Conne, id string) Node {
	return Node{
		id:    id,
		pings: c.Pings,
		pongs: c.Pongs,
	}
}

func (n *Node) Send(msg string) {
	n.pings <- msg
}

func (n *Node) ListenOnNode(oN Node) {
	go func() {
		msg := <-n.pings
		oN.pongs <- msg
	}()
}

func (n *Node) Listen() {

	msg := <-n.pongs
	fmt.Printf("id:%v msg:%s\n", n.id, msg)

}

//The pong function accepts one channel for receives (pings) and a second for sends (pongs).
/*
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
*/
func main() {
	/*
		pings := make(chan string, 1)
		pongs := make(chan string, 1)
		ping(pings, "passed message")
		pong(pings, pongs)
		fmt.Println(<-pongs)*/

	c := Conne{
		Pings: make(chan string, 1),
		Pongs: make(chan string, 1),
	}
	n1 := NewNode(&c, "1")
	n2 := NewNode(&c, "2")

	n1.Send("Hello World")
	n1.ListenOnNode(n2)
	n2.Listen()
}
