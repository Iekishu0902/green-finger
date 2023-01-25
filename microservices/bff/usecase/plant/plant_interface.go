package usecase

import (
	"github.com/green-finger/microservices/bff/domain/entity"
	"github.com/green-finger/microservices/bff/usecase/dto"
)

type PlantInputPort interface {
	Sample(ctx *entity.Context, dto *dto.SampleInputDto) error
}

type PlantOutputPort interface {
	Sample(ctx *entity.Context, dto *dto.SampleOutputDto) error
}
