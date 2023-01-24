package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/green-finger/microservices/bff/cmd/bff_server/servers"
	pb "github.com/green-finger/microservices/bff/infrastructure/rpc"
	"github.com/green-finger/microservices/bff/utils/config"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

var plantServerInst *servers.PlantServer

func run() error {
	ctx := context.Background()
	port := config.GetString("server_port")

	mux := runtime.NewServeMux()

	plantServer := servers.PlantServer{}
	plantServerInst = &plantServer

	err := pb.RegisterPlantHandlerServer(ctx, mux, &plantServer)
	if err != nil {
		return err
	}
	log.Printf("gateway server listen at %v", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
}

func main() {
	run()
}
