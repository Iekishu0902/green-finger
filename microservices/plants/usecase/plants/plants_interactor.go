package usecase

import (
	"github.com/green-finger/microservices/plants/domain/entity"
	"github.com/green-finger/microservices/plants/usecase/dto"
)

type PlantInteractor struct {
	PlantOutputPort PlantOutputPort
}

func NewPlantInteractor(
	plantsOutputPort PlantOutputPort,
) *PlantInteractor {
	return &PlantInteractor{
		PlantOutputPort: plantsOutputPort,
	}
}

func (i *PlantInteractor) GetPlant(ctx *entity.Context, iDto *dto.GetPlantInputDto) error {
	resultMessage := "SUCCESSFUL CALL"
	return i.PlantOutputPort.GetPlant(ctx, &dto.GetPlantOutputDto{Result: resultMessage})
}
