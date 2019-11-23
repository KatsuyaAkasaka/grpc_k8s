package client

import (
	"context"
	"fmt"
	"log"

	pb "github.com/KatsuyaAkasaka/grpc_k8s/pb"

	"google.golang.org/grpc"
)

type HelloClient struct{}

func Client() {
	log.Print("client")
	//sampleなのでwithInsecure
	conn, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	ctx := context.Background()
	client := pb.NewHelloClient(conn)
	message := &pb.HelloRequest{User: "sakas"}
	res, err := client.GetHelloWorld(ctx, message)
	fmt.Printf("result:%#v \n", res)
	fmt.Printf("error::%#v \n", err)
}
