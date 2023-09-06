package main

import (
	"github.com/masharpik/grpc_tracer_test/utils/trace_jaeger"
)

func main() {
	prv, err := trace_jaeger.NewProvider(trace_jaeger.ProviderConfig{
		JaegerEndpoint: "http://meetme-app.ru:14268/api/traces",
		ServiceName:    "chatServer",
		Disabled:       false,
	})
}
