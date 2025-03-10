package p2p

type NetworkNode interface {
	StartListening(port int) error
	ConnectToPeer(address string) error
	SendMsg(peer string, msg []byte) error
	ReceiveMsg() (sender string, msg []byte)
	Close()
}
