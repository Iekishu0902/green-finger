package controller

import (
	"github.com/green-finger/microservices/plant/domain/entity"
	pb "github.com/green-finger/microservices/plant/infrastructure/rpc"
	"github.com/green-finger/microservices/plant/usecase/dto"
	uc "github.com/green-finger/microservices/plant/usecase/plant"
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
