package main

import (
	"context"
	"log"

	"github.com/masharpik/grpc_tracer_test/grpc-backend/pkg/server"
	tracejaeger "github.com/masharpik/grpc_tracer_test/utils/trace_jaeger"
)

func main() {
	ctx := context.Background()

	prv, err := tracejaeger.NewProvider("http://jaeger:14268/api/traces", "gRPC-Backend")
	if err != nil {
		log.Fatalf("tracejaeger.NewProvider: %v", err)
	}
	defer tracejaeger.Close(prv, ctx)

	server := server.InitServer("0.0.0.0", "8081")

	if err = server.RunServer(); err != nil {
		log.Fatalf("server.RunServer: %v", err)
	}
}
