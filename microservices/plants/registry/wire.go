//go:build wireinject
// +build wireinject

package registry

import (
	"github.com/google/wire"
	"github.com/green-finger/microservices/plants/interface/controller"
	"github.com/green-finger/microservices/plants/interface/presenter"
	usecase "github.com/green-finger/microservices/plants/usecase/plants"
)

func InitPlant() *controller.PlantController {
	wire.Build(
		controller.NewPlantController,
		usecase.NewPlantInteractor,
		presenter.NewPlantPresenter,
		wire.Bind(new(usecase.PlantInputPort), new(*usecase.PlantInteractor)),
		wire.Bind(new(usecase.PlantOutputPort), new(*presenter.PlantPresenter)),
	)
	return nil
}
