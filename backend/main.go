package main

import (
	"context"
	"log"

	"github.com/masharpik/grpc_tracer_test/backend/pkg/server"
	"github.com/masharpik/grpc_tracer_test/proto"
	tracejaeger "github.com/masharpik/grpc_tracer_test/utils/trace_jaeger"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()

	prv, err := tracejaeger.NewProvider("http://jaeger:14268/api/traces", "Backend")
	if err != nil {
		log.Fatalf("tracejaeger.NewProvider: %v", err)
	}
	defer tracejaeger.Close(prv, ctx)

	conn, err := grpc.Dial(
		"grpc-backend:8081",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)
	if err != nil {
		log.Fatalf("grpc.Dial: %v", err)
	}
	defer conn.Close()

	client := proto.NewGRPCBackendClient(conn)

	router := server.Init(client)

	if err = server.Run(router, "0.0.0.0", "8080"); err != nil {
		log.Fatalf("server.Run: %v", err)
	}
}
