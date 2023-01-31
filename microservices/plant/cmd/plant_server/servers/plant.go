package servers

import (
	"context"
	"log"

	"github.com/green-finger/microservices/plant/domain/entity"
	pb "github.com/green-finger/microservices/plant/infrastructure/rpc"
	"github.com/green-finger/microservices/plant/interface/controller"
)

type PlantServer struct {
	pb.UnimplementedPlantServer
	Controller *controller.PlantController
}

func (s *PlantServer) GetPlant(ctx context.Context, in *pb.GetPlantRequest) (res *pb.GetPlantResponse, err error) {
	err = nil
	res = nil
	contextEntity := &entity.Context{
		Context: ctx,
	}

	log.Printf("Controller Call")
	err = s.Controller.GetPlant(contextEntity, in)
	if err != nil {
		log.Printf("Controller Call FAILED")
		return
	}
	log.Printf("Controller Call End")

	log.Printf("Set ContextEntity Response")
	r := contextEntity.Response.(pb.GetPlantResponse)
	res = &r
	log.Printf("Set End")
	return
}
