package server

import (
	"context"
	"log"
	"net"

	pb "github.com/KatsuyaAkasaka/grpc_k8s/pb"
	"google.golang.org/grpc"
)

// HelloServer hello world server struct
type HelloServer struct{}

// GetHelloWorld get hello world message
func (hs *HelloServer) GetHelloWorld(ctx context.Context, req *pb.HelloRequest) (res *pb.HelloResponse, err error) {
	message := req.User + ", hello!"
	return &pb.HelloResponse{
		Message: message,
	}, nil
}

// Server server func
func Server() {
	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	helloS := &HelloServer{}
	// 実行したい実処理をseverに登録する
	pb.RegisterHelloServer(server, helloS)
	server.Serve(listenPort)
}
