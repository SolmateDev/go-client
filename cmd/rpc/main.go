package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	solmateRpc "github.com/SolmateDev/go-client/rpc"
	basic "github.com/SolmateDev/go-client/rpc/basic"
)

func main() {
	config, err := basic.ConfigFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	//log.Printf("config=%+v", config)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sigChan := make(chan os.Signal, 1)
	go loopSignal(sigChan, cancel)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	serverErrorC := make(chan error, 1)
	log.Print("starting server")
	err = solmateRpc.Run(ctx, config, serverErrorC)
	if err != nil {
		log.Fatal(err)
	}

	err = <-serverErrorC
	if err != nil {
		log.Fatal(err)
	}

	log.Print("exiting without an error")

}

func loopSignal(sigChan <-chan os.Signal, cancel context.CancelFunc) {
	<-sigChan
	log.Print("received signal for the server to shutdown")
	cancel()
}
