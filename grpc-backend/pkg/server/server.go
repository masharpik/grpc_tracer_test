package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"

	"github.com/masharpik/grpc_tracer_test/proto"
)

type Server struct {
	host string
	port string

	proto.UnimplementedGRPCBackendServer
}

func InitServer(host string, port string) Server {
	var server = Server{
		host: host,
		port: port,
	}

	return server
}

func (server *Server) RunServer() (err error) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", server.host, server.port))
	if err != nil {
		return
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()))
	proto.RegisterGRPCBackendServer(s, server)

	log.Printf("GRPC-Backend успешно запущен\n")
	return s.Serve(lis)
}

func (server *Server) Test(ctx context.Context, in *proto.TestRequest) (resp *proto.TestResponse, err error) {
	_, span := otel.GetTracerProvider().Tracer("gRPC-Backend").Start(ctx, "in grpc-test")
	defer span.End()

	return &proto.TestResponse{}, nil
}
