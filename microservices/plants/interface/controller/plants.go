package controller

import (
	"github.com/green-finger/microservices/plants/domain/entity"
	pb "github.com/green-finger/microservices/plants/infrastructure/rpc"
	"github.com/green-finger/microservices/plants/usecase/dto"
	uc "github.com/green-finger/microservices/plants/usecase/plants"
)

type PlantController struct {
	PlantInputPort uc.PlantInputPort
}

func NewPlantController(inputPort uc.PlantInputPort) *PlantController {
	return &PlantController{
		PlantInputPort: inputPort,
	}
}

func (c *PlantController) GetPlant(ctx *entity.Context, in *pb.GetPlantRequest) error {
	dto := dto.GetPlantInputDto{}
	return c.PlantInputPort.GetPlant(ctx, &dto)
}
