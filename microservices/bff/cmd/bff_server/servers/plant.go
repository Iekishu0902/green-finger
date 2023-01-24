package servers

import (
	"context"

	pb "github.com/green-finger/microservices/bff/infrastructure/rpc"
)

type PlantServer struct {
	pb.UnimplementedPlantServer
}

func (p *PlantServer) Sample(ctx context.Context, in *pb.SampleRequest) (res *pb.SampleResponse, err error) {
	return &pb.SampleResponse{ResponseMessage: "Hello Plant"}, nil
}
