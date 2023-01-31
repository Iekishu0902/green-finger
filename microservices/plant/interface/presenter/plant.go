package presenter

import (
	"github.com/green-finger/microservices/plant/domain/entity"
	pb "github.com/green-finger/microservices/plant/infrastructure/rpc"
	"github.com/green-finger/microservices/plant/usecase/dto"
)

type PlantPresenter struct{}

func NewPlantPresenter() *PlantPresenter {
	return &PlantPresenter{}
}

func (p *PlantPresenter) GetPlant(ctx *entity.Context, in *dto.GetPlantOutputDto) error {

	ctx.Response = pb.GetPlantResponse{}
	return nil
}
