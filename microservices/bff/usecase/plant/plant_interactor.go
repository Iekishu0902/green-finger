package usecase

import (
	"github.com/green-finger/microservices/bff/domain/entity"
	"github.com/green-finger/microservices/bff/usecase/dto"
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

func (i *PlantInteractor) Sample(ctx *entity.Context, iDto *dto.SampleInputDto) error {
	resultMessage := "SUCCESSFUL CALL"
	return i.PlantOutputPort.Sample(ctx, &dto.SampleOutputDto{Result: resultMessage})
}
