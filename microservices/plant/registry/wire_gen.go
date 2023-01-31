// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package registry

import (
	"github.com/green-finger/microservices/plant/interface/controller"
	"github.com/green-finger/microservices/plant/interface/presenter"
	"github.com/green-finger/microservices/plant/usecase/plant"
)

// Injectors from wire.go:

func InitPlant() *controller.PlantController {
	plantPresenter := presenter.NewPlantPresenter()
	plantInteractor := usecase.NewPlantInteractor(plantPresenter)
	plantController := controller.NewPlantController(plantInteractor)
	return plantController
}