package server

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/masharpik/grpc_tracer_test/proto"
	"go.opentelemetry.io/otel"
)

var client proto.GRPCBackendClient

func Init(cl proto.GRPCBackendClient) *mux.Router {
	client = cl

	r := mux.NewRouter()

	r.HandleFunc("/api/test", Test).Methods("GET")
	return r
}

func Run(handler http.Handler, host string, port string) error {
	log.Println("Backend server starts")
	return http.ListenAndServe(host+":"+port, handler)
}

func Test(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	ctxg, span := otel.GetTracerProvider().Tracer("Backend").Start(ctx, "in http-test")
	defer span.End()

	_, err := client.Test(ctxg, &proto.TestRequest{})
	if err != nil {
		log.Printf("client.Test: %v", err)
	}
}
