package main

import (
	"sync"

	"github.com/KatsuyaAkasaka/grpc_k8s/client"
	"github.com/KatsuyaAkasaka/grpc_k8s/server"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		server.Server()
		wg.Done()
	}()
	wg.Wait()
	client.Client()
}
