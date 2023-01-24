package main

import (
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func run() {
	port := "8080"
	mux := runtime.NewServeMux()
	http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
}

func main() {
	run()
}
