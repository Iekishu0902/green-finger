package servers

import (
	"context"
	"log"

	"github.com/green-finger/microservices/bff/domain/entity"
	pb "github.com/green-finger/microservices/bff/infrastructure/rpc"
	"github.com/green-finger/microservices/bff/interface/controller"
)

type PlantServer struct {
	pb.UnimplementedPlantServer
	Controller *controller.PlantController
}

func (s *PlantServer) Sample(ctx context.Context, in *pb.SampleRequest) (res *pb.SampleResponse, err error) {
	err = nil
	res = nil
	contextEntity := &entity.Context{
		Context: ctx,
	}

	log.Printf("Controller Call")
	err = s.Controller.Sample(contextEntity, in)
	if err != nil {
		log.Printf("Controller Call FAILED")
		return
	}
	log.Printf("Controller Call End")

	log.Printf("Set ContextEntity Response")
	r := contextEntity.Response.(pb.SampleResponse)
	res = &r
	log.Printf("Set End")
	return
}
