package usecase

import (
	"github.com/green-finger/microservices/plants/domain/entity"
	"github.com/green-finger/microservices/plants/usecase/dto"
)

type PlantInputPort interface {
	GetPlant(ctx *entity.Context, dto *dto.GetPlantInputDto) error
}

type PlantOutputPort interface {
	GetPlant(ctx *entity.Context, dto *dto.GetPlantOutputDto) error
}
