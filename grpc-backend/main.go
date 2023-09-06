package main

import (
	"github.com/masharpik/grpc_tracer_test/utils/trace_jaeger"
)

func main() {
	prv, err := tracejaeger.NewProvider(tracejaeger.ProviderConfig{
		JaegerEndpoint: "http://meetme-app.ru:14268/api/traces",
		ServiceName:    "chatServer",
		Disabled:       false,
	})
}
