package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/green-finger/microservices/plants/cmd/plants_server/servers"
	pb "github.com/green-finger/microservices/plants/infrastructure/rpc"
	"github.com/green-finger/microservices/plants/registry"
	"github.com/green-finger/microservices/plants/utils/config"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

var plantsServerInst *servers.PlantServer

func run() error {
	ctx := context.Background()
	port := config.GetString("server_port")

	mux := runtime.NewServeMux()

	plantsServer := servers.PlantServer{}
	plantsServerInst = &plantsServer
	plantsServer.Controller = registry.InitPlant()
	err := pb.RegisterPlantHandlerServer(ctx, mux, &plantsServer)
	if err != nil {
		return err
	}
	log.Printf("gateway server listen at %v", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
}

func main() {
	run()
}
