package buildings

import (
	"context"
	"errors"

	"github.com/sotskov-do/oms-assignment/internal/models"
	"github.com/sotskov-do/oms-assignment/internal/storage"
)

//go:generate minimock -i github.com/sotskov-do/oms-assignment/internal/service/buildings.BuildingsService -o ../mocks/
type BuildingsService interface {
	GetBuildings(ctx context.Context) (models.BuildingSlice, error)
	GetBuilding(ctx context.Context, id int) (*models.Building, error)
	CreateBuilding(ctx context.Context, building *models.Building) error
	DeleteBuilding(ctx context.Context, id int) error
}

type Service struct {
	buildingsStorage storage.BuildingsStorage
}

func NewService(buildingsStorage storage.BuildingsStorage) *Service {
	return &Service{
		buildingsStorage: buildingsStorage,
	}
}

func (s *Service) GetBuildings(ctx context.Context) (models.BuildingSlice, error) {
	buildings, err := s.buildingsStorage.GetBuildings(ctx)
	if err != nil {
		return nil, err
	}

	return buildings, nil
}

func (s *Service) GetBuilding(ctx context.Context, id int) (*models.Building, error) {
	if id <= 0 {
		return nil, errors.New("id less or equal 0")
	}

	building, err := s.buildingsStorage.GetBuilding(ctx, id)
	if err != nil {
		return nil, err
	}

	return building, nil
}

func (s *Service) CreateBuilding(ctx context.Context, building *models.Building) error {
	err := s.buildingsStorage.CreateBuilding(ctx, building)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteBuilding(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.New("id less or equal 0")
	}

	n, err := s.buildingsStorage.DeleteBuilding(ctx, id)
	if err != nil {
		return err
	}

	if n == 0 {
		return errors.New("no building with such id")
	}

	return nil
}
