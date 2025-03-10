package main

import (
	"github.com/WhileCodingDoLearn/my_df_system/p2p"
)

func main() {
	n1 := p2p.NewTCPServer()
	n1.StartListening(8080)

}
