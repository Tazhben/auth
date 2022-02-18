package main

import (
	"fmt"
	"log"
)

func main() {
	log.Fatalln(fmt.Sprintf("Service shut down: %v", run()))
}

func run() error {
	//ctx, cancel := context.WithCancel(context.Background())
	//
	//defer cancel()
	//
	//errCh := make(chan error, 1)
	return nil
}
