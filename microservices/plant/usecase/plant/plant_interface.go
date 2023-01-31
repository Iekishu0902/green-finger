package usecase

import (
	"github.com/green-finger/microservices/plant/domain/entity"
	"github.com/green-finger/microservices/plant/usecase/dto"
)

type PlantInputPort interface {
	GetPlant(ctx *entity.Context, dto *dto.GetPlantInputDto) error
}

type PlantOutputPort interface {
	GetPlant(ctx *entity.Context, dto *dto.GetPlantOutputDto) error
}
