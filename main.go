package main

import (
	"time"

	"github.com/WhileCodingDoLearn/my_df_system/p2p/connection"
)

func main() {
	conn := connection.Connection{}
	go func() {
		time.Sleep(1 * time.Second)
		conn.State(connection.Close)
	}()
	conn.Listen()
}
