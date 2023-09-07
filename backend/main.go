package main

import (
	"context"
	"log"

	"github.com/masharpik/grpc_tracer_test/backend/pkg/server"
	tracejaeger "github.com/masharpik/grpc_tracer_test/utils/trace_jaeger"
)

func main() {
	ctx := context.Background()

	prv, err := tracejaeger.NewProvider("http://localhost:14268/tracer", "gRPC-Backend")
	if err != nil {
		log.Fatalf("tracejaeger.NewProvider: %v", err)
	}
	defer tracejaeger.Close(prv, ctx)

	router := server.Init()
	if err = server.Run(router, "0.0.0.0", "8080"); err != nil {
		log.Fatalf("server.Run: %v", err)
	}
}
