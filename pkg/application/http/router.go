package http

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/tembleking/test-grpc-gateway/pkg/application/proto"
)

func NewRouter(ctx context.Context, grpcListenAddress string) (http.Handler, error) {
	mux := runtime.NewServeMux()

	options := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := proto.RegisterTestServiceHandlerFromEndpoint(ctx, mux, grpcListenAddress, options)
	if err != nil {
		return nil, err
	}

	return mux, nil
}
