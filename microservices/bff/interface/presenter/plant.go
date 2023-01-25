package presenter

import (
	"github.com/green-finger/microservices/bff/domain/entity"
	pb "github.com/green-finger/microservices/bff/infrastructure/rpc"
	"github.com/green-finger/microservices/bff/usecase/dto"
)

type PlantPresenter struct{}

func NewPlantPresenter() *PlantPresenter {
	return &PlantPresenter{}
}

func (p *PlantPresenter) Sample(ctx *entity.Context, in *dto.SampleOutputDto) error {

	ctx.Response = pb.SampleResponse{
		ResponseMessage: in.Result,
	}
	return nil
}
