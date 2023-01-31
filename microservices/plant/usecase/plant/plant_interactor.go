package usecase

import (
	"github.com/green-finger/microservices/plant/domain/entity"
	"github.com/green-finger/microservices/plant/usecase/dto"
)

type PlantInteractor struct {
	PlantOutputPort PlantOutputPort
}

func NewPlantInteractor(
	plantOutputPort PlantOutputPort,
) *PlantInteractor {
	return &PlantInteractor{
		PlantOutputPort: plantOutputPort,
	}
}

func (i *PlantInteractor) GetPlant(ctx *entity.Context, iDto *dto.GetPlantInputDto) error {
	resultMessage := "SUCCESSFUL CALL"
	return i.PlantOutputPort.GetPlant(ctx, &dto.GetPlantOutputDto{Result: resultMessage})
}
