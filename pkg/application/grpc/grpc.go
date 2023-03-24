package grpc

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/tembleking/test-grpc-gateway/pkg/application/proto"
)

type server struct {
}

func (s *server) ReceiveName(ctx context.Context, request *proto.TestRequestWithName) (*proto.TestResponse, error) {
	return &proto.TestResponse{Message: "Hello " + request.Name}, nil
}

func (s *server) Test(ctx context.Context, empty *emptypb.Empty) (*proto.TestResponse, error) {
	return &proto.TestResponse{Message: "Hello world!!!"}, nil
}

func NewServer() *grpc.Server {
	grpcServer := grpc.NewServer()
	server := &server{}
	proto.RegisterTestServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	return grpcServer
}
