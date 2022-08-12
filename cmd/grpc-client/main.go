package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/tembleking/test-grpc-gateway/pkg/application/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("unable to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := proto.NewTestServiceClient(conn)

	response, err := client.Test(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("could not send request to the gRPC server: %v", err)
	}

	log.Printf("%s", response.Message)
}
