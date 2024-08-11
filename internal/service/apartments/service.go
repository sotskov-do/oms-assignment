package apartments

import (
	"context"
	"errors"
	"fmt"

	"github.com/sotskov-do/oms-assignment/internal/models"
	"github.com/sotskov-do/oms-assignment/internal/storage"
)

//go:generate minimock -i github.com/sotskov-do/oms-assignment/internal/service/apartments.ApartmentsService -o ../mocks/
type ApartmentsService interface {
	GetApartments(ctx context.Context) (models.ApartmentSlice, error)
	GetApartment(ctx context.Context, id int) (*models.Apartment, error)
	GetApartmentsInBuilding(ctx context.Context, buildingId int) (models.ApartmentSlice, error)
	CreateApartment(ctx context.Context, apartment *models.Apartment) error
	DeleteApartment(ctx context.Context, id int) error
}

type Service struct {
	apartmentsStorage storage.ApartmentsStorage
}

func NewService(apartmentsStorage storage.ApartmentsStorage) *Service {
	return &Service{
		apartmentsStorage: apartmentsStorage,
	}
}

func (s *Service) GetApartments(ctx context.Context) (models.ApartmentSlice, error) {
	apartments, err := s.apartmentsStorage.GetApartments(ctx)
	if err != nil {
		return nil, err
	}

	return apartments, nil
}

func (s *Service) GetApartment(ctx context.Context, id int) (*models.Apartment, error) {
	if id <= 0 {
		return nil, errors.New("id less or equal 0")
	}

	apartment, err := s.apartmentsStorage.GetApartment(ctx, id)
	if err != nil {
		return nil, err
	}

	return apartment, nil
}

func (s *Service) GetApartmentsInBuilding(ctx context.Context, buildingId int) (models.ApartmentSlice, error) {
	if buildingId <= 0 {
		return nil, errors.New("building id less or equal 0")
	}

	apartmentsInBuilding, err := s.apartmentsStorage.GetApartmentsInBuilding(ctx, buildingId)
	if err != nil {
		return nil, err
	}

	return apartmentsInBuilding, nil
}

func (s *Service) CreateApartment(ctx context.Context, apartment *models.Apartment) error {
	err := s.apartmentsStorage.CreateApartment(ctx, apartment)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteApartment(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.New("id less or equal 0")
	}

	n, err := s.apartmentsStorage.DeleteApartment(ctx, id)
	if err != nil {
		return err
	}

	if n == 0 {
		return fmt.Errorf("no apartment with id [%v]", id)
	}

	return nil
}
