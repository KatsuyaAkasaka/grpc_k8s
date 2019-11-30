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
func (hs *HelloServer) GetHelloWorld(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	message := "hello, " + req.User
	return &pb.HelloResponse{
		Message: message,
	}, nil
}

func GRPCStart() {
	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	helloS := &HelloServer{}
	// 実行したい実処理をseverに登録する
	pb.RegisterHelloServer(s, helloS)
	go func() {
		if err := s.Serve(listenPort); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	log.Printf("listen to grpc port %d", 19003)
}
