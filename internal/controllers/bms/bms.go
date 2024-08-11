package bms

import (
	"github.com/sotskov-do/oms-assignment/internal/service/apartments"
	"github.com/sotskov-do/oms-assignment/internal/service/buildings"
)

const (
	resultKey   = "result"
	responseKey = "response"

	resultSuccess = "success"
	resultError   = "error"
)

type BuildingManagementSystem struct {
	apartmentsService apartments.ApartmentsService
	buildingsService  buildings.BuildingsService
}

func NewBuildingManagementSystem(
	apartmentsService apartments.ApartmentsService,
	buildingsService buildings.BuildingsService,
) *BuildingManagementSystem {
	return &BuildingManagementSystem{
		apartmentsService: apartmentsService,
		buildingsService:  buildingsService,
	}
}
