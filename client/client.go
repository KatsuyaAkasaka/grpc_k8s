package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	pb "github.com/KatsuyaAkasaka/grpc_k8s/pb"
	server "github.com/KatsuyaAkasaka/grpc_k8s/server"

	"google.golang.org/grpc"
)

type HelloClient struct{}

var conn *grpc.ClientConn

const helloWorld = true

func SendHello() string {
	if helloWorld {
		return "hello world"
	}
	ctx := context.Background()
	client := pb.NewHelloClient(conn)
	message := &pb.HelloRequest{User: "sakas"}
	res, _ := client.GetHelloWorld(ctx, message)
	return res.Message
}

func Connect() {
	var err error
	//sampleなのでwithInsecure
	conn, err = grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	res := SendHello()
	fmt.Fprintf(w, res)
}

func main() {
	server.GRPCStart()
	time.Sleep(1 * time.Second)
	Connect()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
