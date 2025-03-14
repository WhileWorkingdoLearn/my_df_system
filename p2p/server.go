package p2p

import (
	"fmt"
	"net"

	"github.com/WhileCodingDoLearn/my_df_system/msg"
	handler "github.com/WhileCodingDoLearn/my_df_system/msghandler"
)

type TCPNode struct {
	peer     map[string]net.Conn
	listener net.Listener
	receive  chan msg.MsgHeader
	quit     chan struct{}
	mux      handler.SeverMux
}

func NewServer() TCPNode {
	return TCPNode{
		peer:    make(map[string]net.Conn),
		receive: make(chan msg.MsgHeader, 10),
		quit:    make(chan struct{}),
	}
}

func (n *TCPNode) StartListening(port int) error {

	ln, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		return err
	}

	if n.listener != nil {
		return fmt.Errorf("node is already listening")
	}

	n.listener = ln
	fmt.Println("Node is listening on Port :", port)

	go func() {
		for {
			conn, err := n.listener.Accept()
			if err != nil {
				fmt.Println(err)
				continue
			}

			go n.handleConnection(conn)
		}
	}()
	return nil

}

func (n *TCPNode) ConnectToPeer(address string) error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	n.peer[address] = conn
	fmt.Println("Connected to ", address)
	return nil
}

func (n *TCPNode) SendMsg(peer string, msg []byte) error {
	/*
		conn, found := n.peer[peer]
		if !found {
			return fmt.Errorf("peer not found %s", peer)
		}
		_, err := conn.Write(ConvertToBinary(0x01, msg))
		return err*/
	return nil
}

func (n *TCPNode) ReceiveMsg() (msg.MsgHeader, error) {

	msg := <-n.receive

	return msg, nil

}

func (n *TCPNode) Close() {
	close(n.quit)
	n.listener.Close()
	for _, conn := range n.peer {
		conn.Close()
	}
}

func (node *TCPNode) handleConnection(conn net.Conn) {
	peerAddr := conn.RemoteAddr().String()
	node.peer[peerAddr] = conn
	defer conn.Close()
	msgHeader, err := msg.DecodeMsgHeader(conn)
	if err != nil {
		delete(node.peer, peerAddr)
		conn.Close()
	}
	node.receive <- msgHeader
	//buffer := make([]byte, 1024)

}
