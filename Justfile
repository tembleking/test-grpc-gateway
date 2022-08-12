

build-proto:
    #!/usr/bin/env bash
    cd pkg/application/proto
    buf mod update
    buf generate
	
deps:
	go mod tidy
	go install github.com/bufbuild/buf/cmd/buf
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
