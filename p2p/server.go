package p2p

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/WhileCodingDoLearn/my_df_system/msg"
	nmsgp "github.com/WhileCodingDoLearn/my_df_system/p2p/mux"
)

/*
	TCP implementation of

*/

type Config struct {
	Mux     nmsgp.SeverMux
	Timeout time.Duration
	Port    int
}

type TCPNode struct {
	cfg      Config
	peer     map[string]net.Conn
	peerLock sync.RWMutex
	listener net.Listener
	receive  chan msg.MsgHeader
	quit     chan struct{}
	wg       sync.WaitGroup
}

func NewServer(cfg Config) TCPNode {
	return TCPNode{
		peer:    make(map[string]net.Conn),
		receive: make(chan msg.MsgHeader, 10),
		quit:    make(chan struct{}),
		cfg:     cfg,
	}
}

func (n *TCPNode) StartListening() error {

	ln, err := net.Listen("tcp", fmt.Sprintf(":%v", n.cfg.Port))
	if err != nil {
		return err
	}

	if n.listener != nil {
		return fmt.Errorf("node is already listening")
	}

	n.listener = ln
	fmt.Println("Node is listening on Port :", n.cfg.Port)

	for {
		conn, err := n.listener.Accept()
		if err != nil {
			fmt.Println(err)
			return err
		}
		n.wg.Add(1)
		conn.SetDeadline(time.Now().Add(10 * time.Second))
		go n.handleConnection(conn)
	}

}

func (n *TCPNode) ConnectToPeer(address string) error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	n.peerLock.Lock()
	defer n.peerLock.Unlock()
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

func (node *TCPNode) addConnection(conn net.Conn) string {
	peerAddr := conn.RemoteAddr().String()
	node.peerLock.Lock()
	defer node.peerLock.Unlock()
	node.peer[peerAddr] = conn
	return peerAddr
}

func (node *TCPNode) removeConnection(conn string) {
	node.peerLock.Lock()
	defer node.peerLock.Unlock()
	delete(node.peer, conn)
}

func (node *TCPNode) handleConnection(conn net.Conn) {
	defer node.wg.Done()

	id := node.addConnection(conn)

	defer node.removeConnection(id)

	msg, err := msg.DecodeMsgHeader(conn)
	if err != nil {
		conn.Close()
	}

	ctx, cancel := context.WithTimeout(context.Background(), msg.Timeout)
	req := nmsgp.NewRequest(msg, ctx)
	res := nmsgp.NewResponse()
	defer cancel()
	f, _ := node.cfg.Mux.Handler(req)

	f.ForwardMsg(res, req)

	node.sendResponse(res, conn)

}

func (node *TCPNode) sendResponse(res *nmsgp.Response, conn net.Conn) {

}
