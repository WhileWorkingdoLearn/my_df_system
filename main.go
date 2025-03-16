package main

import (
	"fmt"
	"time"

	handler "github.com/WhileCodingDoLearn/my_df_system/msghandler"
	"github.com/WhileCodingDoLearn/my_df_system/p2p"
)

func main() {

	mux := handler.NewServerMux("Hello world Domain")
	mux.HandleFunc("/v1", handler.FETCH.String(), func(res handler.ResponseWriter, req *handler.Request) {
		fmt.Println(req.Endpoint())
	})

	sever := p2p.NewServer(p2p.Config{
		Timeout: 10 * time.Second,
		Port:    3000,
		Mux:     mux,
	})

	go func() {

		time.Sleep(1 * time.Second)

		p2p.StartClient(":3000")

		time.Sleep(1 * time.Second)

		sever.Close()

	}()

	sever.StartListening()
}
