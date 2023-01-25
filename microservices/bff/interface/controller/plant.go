package controller

import (
	"github.com/green-finger/microservices/bff/domain/entity"
	pb "github.com/green-finger/microservices/bff/infrastructure/rpc"
	"github.com/green-finger/microservices/bff/usecase/dto"
	uc "github.com/green-finger/microservices/bff/usecase/plant"
)

type PlantController struct {
	SystemInputPort uc.PlantInputPort
}

func NewPlantController(inputPort uc.PlantInputPort) *PlantController {
	return &PlantController{
		SystemInputPort: inputPort,
	}
}

func (c *PlantController) Sample(ctx *entity.Context, in *pb.SampleRequest) error {
	dto := dto.SampleInputDto{
		Message: in.Message,
	}
	return c.SystemInputPort.Sample(ctx, &dto)
}
