package apartments

import (
	"context"
	"errors"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/sotskov-do/oms-assignment/internal/models"
	"github.com/sotskov-do/oms-assignment/internal/storage"
	storage_mocks "github.com/sotskov-do/oms-assignment/internal/storage/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/null/v8"
)

func Test_GetApartments(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                 string
		getApartmentsStorage func(mc *minimock.Controller) storage.ApartmentsStorage
		want                 models.ApartmentSlice
		wantErr              bool
	}{
		{
			name: "valid",
			getApartmentsStorage: func(mc *minimock.Controller) storage.ApartmentsStorage {
				return storage_mocks.NewApartmentsStorageMock(mc).
					GetApartmentsMock.
					Expect(minimock.AnyContext).
					Return(models.ApartmentSlice{
						{
							ID:         1,
							BuildingID: 1,
							Number:     null.String{Valid: true, String: "10"},
							Floor:      null.Int{Valid: true, Int: 2},
							SQMeters:   null.Int{Valid: true, Int: 20},
						},
						{
							ID:         2,
							BuildingID: 1,
							Number:     null.String{Valid: true, String: "11"},
							Floor:      null.Int{Valid: true, Int: 2},
							SQMeters:   null.Int{Valid: true, Int: 25},
						},
					}, nil)
			},
			want: models.ApartmentSlice{
				{
					ID:         1,
					BuildingID: 1,
					Number:     null.String{Valid: true, String: "10"},
					Floor:      null.Int{Valid: true, Int: 2},
					SQMeters:   null.Int{Valid: true, Int: 20},
				},
				{
					ID:         2,
					BuildingID: 1,
					Number:     null.String{Valid: true, String: "11"},
					Floor:      null.Int{Valid: true, Int: 2},
					SQMeters:   null.Int{Valid: true, Int: 25},
				},
			},
		},
		{
			name: "storageError",
			getApartmentsStorage: func(mc *minimock.Controller) storage.ApartmentsStorage {
				return storage_mocks.NewApartmentsStorageMock(mc).
					GetApartmentsMock.
					Expect(minimock.AnyContext).
					Return(nil, errors.New("storageError"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mc := minimock.NewController(t)
			apartmentsStorage := tt.getApartmentsStorage(mc)
			s := Service{apartmentsStorage: apartmentsStorage}

			got, err := s.GetApartments(context.Background())
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_GetApartment(t *testing.T) {
	t.Parallel()

	type args struct {
		id int
	}

	tests := []struct {
		name                 string
		args                 args
		getApartmentsStorage func(mc *minimock.Controller) storage.ApartmentsStorage
		want                 *models.Apartment
		wantErr              bool
	}{
		{
			name: "valid",
			args: args{
				id: 1,
			},
			getApartmentsStorage: func(mc *minimock.Controller) storage.ApartmentsStorage {
				return storage_mocks.NewApartmentsStorageMock(mc).
					GetApartmentMock.
					Expect(minimock.AnyContext, 1).
					Return(&models.Apartment{
						ID:         1,
						BuildingID: 1,
						Number:     null.String{Valid: true, String: "10"},
						Floor:      null.Int{Valid: true, Int: 2},
						SQMeters:   null.Int{Valid: true, Int: 20},
					}, nil)
			},
			want: &models.Apartment{
				ID:         1,
				BuildingID: 1,
				Number:     null.String{Valid: true, String: "10"},
				Floor:      null.Int{Valid: true, Int: 2},
				SQMeters:   null.Int{Valid: true, Int: 20},
			},
		},
		{
			name: "wrongID",
			args: args{
				id: 0,
			},
			getApartmentsStorage: func(mc *minimock.Controller) storage.ApartmentsStorage {
				return nil
			},
			wantErr: true,
		},
		{
			name: "storageError",
			args: args{
				id: 2,
			},
			getApartmentsStorage: func(mc *minimock.Controller) storage.ApartmentsStorage {
				return storage_mocks.NewApartmentsStorageMock(mc).
					GetApartmentMock.
					Expect(minimock.AnyContext, 2).
					Return(nil, errors.New("storageError"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mc := minimock.NewController(t)
			apartmentsStorage := tt.getApartmentsStorage(mc)
			s := Service{apartmentsStorage: apartmentsStorage}

			got, err := s.GetApartment(context.Background(), tt.args.id)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_GetApartmentsInBuilding(t *testing.T) {
	t.Parallel()

	type args struct {
		buildingId int
	}

	tests := []struct {
		name                 string
		args                 args
		getApartmentsStorage func(mc *minimock.Controller) storage.ApartmentsStorage
		want                 models.ApartmentSlice
		wantErr              bool
	}{
		{
			name: "valid",
			args: args{
				buildingId: 1,
			},
			getApartmentsStorage: func(mc *minimock.Controller) storage.ApartmentsStorage {
				return storage_mocks.NewApartmentsStorageMock(mc).
					GetApartmentsInBuildingMock.
					Expect(minimock.AnyContext, 1).
					Return(models.ApartmentSlice{
						{
							ID:         1,
							BuildingID: 1,
							Number:     null.String{Valid: true, String: "10"},
							Floor:      null.Int{Valid: true, Int: 2},
							SQMeters:   null.Int{Valid: true, Int: 20},
						},
						{
							ID:         2,
							BuildingID: 1,
							Number:     null.String{Valid: true, String: "11"},
							Floor:      null.Int{Valid: true, Int: 2},
							SQMeters:   null.Int{Valid: true, Int: 25},
						},
					}, nil)
			},
			want: models.ApartmentSlice{
				{
					ID:         1,
					BuildingID: 1,
					Number:     null.String{Valid: true, String: "10"},
					Floor:      null.Int{Valid: true, Int: 2},
					SQMeters:   null.Int{Valid: true, Int: 20},
				},
				{
					ID:         2,
					BuildingID: 1,
					Number:     null.String{Valid: true, String: "11"},
					Floor:      null.Int{Valid: true, Int: 2},
					SQMeters:   null.Int{Valid: true, Int: 25},
				},
			},
		},
		{
			name: "wrongID",
			args: args{
				buildingId: 0,
			},
			getApartmentsStorage: func(mc *minimock.Controller) storage.ApartmentsStorage {
				return nil
			},
			wantErr: true,
		},
		{
			name: "storageError",
			args: args{
				buildingId: 1,
			},
			getApartmentsStorage: func(mc *minimock.Controller) storage.ApartmentsStorage {
				return storage_mocks.NewApartmentsStorageMock(mc).
					GetApartmentsInBuildingMock.
					Expect(minimock.AnyContext, 1).
					Return(nil, errors.New("storageError"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mc := minimock.NewController(t)
			apartmentsStorage := tt.getApartmentsStorage(mc)
			s := Service{apartmentsStorage: apartmentsStorage}

			got, err := s.GetApartmentsInBuilding(context.Background(), tt.args.buildingId)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_CreateApartment(t *testing.T) {
	t.Parallel()

	type args struct {
		apartment *models.Apartment
	}

	tests := []struct {
		name                 string
		args                 args
		getApartmentsStorage func(mc *minimock.Controller) storage.ApartmentsStorage
		wantErr              bool
	}{
		{
			name: "valid",
			args: args{
				apartment: &models.Apartment{
					BuildingID: 1,
					Number:     null.String{Valid: true, String: "10"},
					Floor:      null.Int{Valid: true, Int: 2},
					SQMeters:   null.Int{Valid: true, Int: 20},
				},
			},
			getApartmentsStorage: func(mc *minimock.Controller) storage.ApartmentsStorage {
				return storage_mocks.NewApartmentsStorageMock(mc).
					CreateApartmentMock.
					Expect(minimock.AnyContext, &models.Apartment{
						BuildingID: 1,
						Number:     null.String{Valid: true, String: "10"},
						Floor:      null.Int{Valid: true, Int: 2},
						SQMeters:   null.Int{Valid: true, Int: 20},
					}).
					Return(nil)
			},
		},
		{
			name: "storageError",
			args: args{
				apartment: &models.Apartment{
					BuildingID: 1,
					Number:     null.String{Valid: true, String: "10"},
					Floor:      null.Int{Valid: true, Int: 2},
					SQMeters:   null.Int{Valid: true, Int: 20},
				},
			},
			getApartmentsStorage: func(mc *minimock.Controller) storage.ApartmentsStorage {
				return storage_mocks.NewApartmentsStorageMock(mc).
					CreateApartmentMock.
					Expect(minimock.AnyContext, &models.Apartment{
						BuildingID: 1,
						Number:     null.String{Valid: true, String: "10"},
						Floor:      null.Int{Valid: true, Int: 2},
						SQMeters:   null.Int{Valid: true, Int: 20},
					}).
					Return(errors.New("storageError"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mc := minimock.NewController(t)
			apartmentsStorage := tt.getApartmentsStorage(mc)
			s := Service{apartmentsStorage: apartmentsStorage}

			err := s.CreateApartment(context.Background(), tt.args.apartment)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
		})
	}
}

func Test_DeleteApartment(t *testing.T) {
	t.Parallel()

	type args struct {
		id int
	}

	tests := []struct {
		name                 string
		args                 args
		getApartmentsStorage func(mc *minimock.Controller) storage.ApartmentsStorage
		wantErr              bool
	}{
		{
			name: "valid",
			args: args{
				id: 1,
			},
			getApartmentsStorage: func(mc *minimock.Controller) storage.ApartmentsStorage {
				return storage_mocks.NewApartmentsStorageMock(mc).
					DeleteApartmentMock.
					Expect(minimock.AnyContext, 1).
					Return(1, nil)
			},
		},
		{
			name: "wrongID",
			args: args{
				id: 0,
			},
			getApartmentsStorage: func(mc *minimock.Controller) storage.ApartmentsStorage {
				return nil
			},
			wantErr: true,
		},
		{
			name: "noRowToDelete",
			args: args{
				id: 2,
			},
			getApartmentsStorage: func(mc *minimock.Controller) storage.ApartmentsStorage {
				return storage_mocks.NewApartmentsStorageMock(mc).
					DeleteApartmentMock.
					Expect(minimock.AnyContext, 2).
					Return(0, nil)
			},
			wantErr: true,
		},
		{
			name: "storageError",
			args: args{
				id: 2,
			},
			getApartmentsStorage: func(mc *minimock.Controller) storage.ApartmentsStorage {
				return storage_mocks.NewApartmentsStorageMock(mc).
					DeleteApartmentMock.
					Expect(minimock.AnyContext, 2).
					Return(0, errors.New("storageError"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mc := minimock.NewController(t)
			apartmentsStorage := tt.getApartmentsStorage(mc)
			s := Service{apartmentsStorage: apartmentsStorage}

			err := s.DeleteApartment(context.Background(), tt.args.id)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
		})
	}
}
