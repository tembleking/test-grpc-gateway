package main

import (
	"context"
	"fmt"
	"log"
	"net"
	gohttp "net/http"
	"os"
	"os/signal"

	"github.com/tembleking/test-grpc-gateway/pkg/application/grpc"
	"github.com/tembleking/test-grpc-gateway/pkg/application/http"
)

const (
	grpcPort = ":9090"
	httpPort = ":8080"
)

func main() {
	ctx, exit := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer exit()

	listener, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Println("failed to listen: %v", err)
		exit()
	}

	grpcServer := grpc.NewServer()
	go func() {
		fmt.Printf("Starting gRPC server on %s\n", grpcPort)
		err = grpcServer.Serve(listener)
		if err != nil {
			log.Println("failed to serve: %v", err)
			exit()
		}
	}()
	defer grpcServer.GracefulStop()

	go func() {
		fmt.Printf("Starting HTTP server on %s\n", httpPort)
		httpServer, err := http.NewRouter(ctx, listener.Addr().String())
		if err != nil {
			log.Printf("failed to create http server: %v", err)
			exit()
		}

		log.Print(gohttp.ListenAndServe(httpPort, httpServer))
		exit()
	}()

	<-ctx.Done()
	fmt.Println("exiting")
}
