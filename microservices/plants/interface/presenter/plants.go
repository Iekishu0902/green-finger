package presenter

import (
	"github.com/green-finger/microservices/plants/domain/entity"
	pb "github.com/green-finger/microservices/plants/infrastructure/rpc"
	"github.com/green-finger/microservices/plants/usecase/dto"
)

type PlantPresenter struct{}

func NewPlantPresenter() *PlantPresenter {
	return &PlantPresenter{}
}

func (p *PlantPresenter) GetPlant(ctx *entity.Context, in *dto.GetPlantOutputDto) error {

	ctx.Response = pb.GetPlantResponse{}
	return nil
}
