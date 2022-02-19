package main

import (
	"Tazhben/auth/config"
	"Tazhben/auth/internal/transport/http"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Fatalln(fmt.Sprintf("Service shut down: %v", run()))
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	errCh := make(chan error, 1)

	configs := config.NewConfigServer()

	go http.StartServer(ctx, errCh, configs)

	go shutDown(errCh)

	return <-errCh
}

func shutDown(errCh chan<- error) {
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)
	errCh <- fmt.Errorf("%s", <-sigCh)
}
