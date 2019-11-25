package main

import (
	"sync"

	client "github.com/KatsuyaAkasaka/grpc_k8s/client"
	server "github.com/KatsuyaAkasaka/grpc_k8s/server"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		server.GRPCStart()
		wg.Done()
	}()
	wg.Wait()
	client.Start()
}
