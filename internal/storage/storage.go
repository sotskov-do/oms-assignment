package storage

import (
	"context"

	"github.com/sotskov-do/oms-assignment/internal/models"
)

type ApartmentsStorage interface {
	GetApartments(ctx context.Context) (models.ApartmentSlice, error)
	GetApartment(ctx context.Context, id int) (*models.Apartment, error)
	GetApartmentsInBuilding(ctx context.Context, buildingId int) (models.ApartmentSlice, error)
	CreateApartment(ctx context.Context, apartment *models.Apartment) error
	DeleteApartment(ctx context.Context, id int) (int64, error)
}

type BuildingsStorage interface {
	GetBuildings(ctx context.Context) (models.BuildingSlice, error)
	GetBuilding(ctx context.Context, id int) (*models.Building, error)
	CreateBuilding(ctx context.Context, building *models.Building) error
	DeleteBuilding(ctx context.Context, id int) (int64, error)
}
