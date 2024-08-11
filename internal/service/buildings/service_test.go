package buildings

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

func Test_GetBuildings(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                string
		getBuildingsStorage func(mc *minimock.Controller) storage.BuildingsStorage
		want                models.BuildingSlice
		wantErr             bool
	}{
		{
			name: "valid",
			getBuildingsStorage: func(mc *minimock.Controller) storage.BuildingsStorage {
				return storage_mocks.NewBuildingsStorageMock(mc).
					GetBuildingsMock.
					Expect(minimock.AnyContext).
					Return(models.BuildingSlice{
						{
							ID:      1,
							Name:    "building_1",
							Address: null.String{Valid: true, String: "address_1"},
						},
						{
							ID:      2,
							Name:    "building_2",
							Address: null.String{Valid: true, String: "address_2"},
						},
					}, nil)
			},
			want: models.BuildingSlice{
				{
					ID:      1,
					Name:    "building_1",
					Address: null.String{Valid: true, String: "address_1"},
				},
				{
					ID:      2,
					Name:    "building_2",
					Address: null.String{Valid: true, String: "address_2"},
				},
			},
		},
		{
			name: "storageError",
			getBuildingsStorage: func(mc *minimock.Controller) storage.BuildingsStorage {
				return storage_mocks.NewBuildingsStorageMock(mc).
					GetBuildingsMock.
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
			buildingsStorage := tt.getBuildingsStorage(mc)
			s := Service{buildingsStorage: buildingsStorage}

			got, err := s.GetBuildings(context.Background())
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_GetBuilding(t *testing.T) {
	t.Parallel()

	type args struct {
		id int
	}

	tests := []struct {
		name                string
		args                args
		getBuildingsStorage func(mc *minimock.Controller) storage.BuildingsStorage
		want                *models.Building
		wantErr             bool
	}{
		{
			name: "valid",
			args: args{
				id: 1,
			},
			getBuildingsStorage: func(mc *minimock.Controller) storage.BuildingsStorage {
				return storage_mocks.NewBuildingsStorageMock(mc).
					GetBuildingMock.
					Expect(minimock.AnyContext, 1).
					Return(&models.Building{
						ID:      1,
						Name:    "building_1",
						Address: null.String{Valid: true, String: "address_1"},
					}, nil)
			},
			want: &models.Building{
				ID:      1,
				Name:    "building_1",
				Address: null.String{Valid: true, String: "address_1"},
			},
		},
		{
			name: "wrongID",
			args: args{
				id: 0,
			},
			getBuildingsStorage: func(mc *minimock.Controller) storage.BuildingsStorage {
				return nil
			},
			wantErr: true,
		},
		{
			name: "storageError",
			args: args{
				id: 2,
			},
			getBuildingsStorage: func(mc *minimock.Controller) storage.BuildingsStorage {
				return storage_mocks.NewBuildingsStorageMock(mc).
					GetBuildingMock.
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
			buildingsStorage := tt.getBuildingsStorage(mc)
			s := Service{buildingsStorage: buildingsStorage}

			got, err := s.GetBuilding(context.Background(), tt.args.id)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_CreateBuilding(t *testing.T) {
	t.Parallel()

	type args struct {
		building *models.Building
	}

	tests := []struct {
		name                string
		args                args
		getBuildingsStorage func(mc *minimock.Controller) storage.BuildingsStorage
		wantErr             bool
	}{
		{
			name: "valid",
			args: args{
				building: &models.Building{
					Name:    "building_1",
					Address: null.String{Valid: true, String: "address_1"},
				},
			},
			getBuildingsStorage: func(mc *minimock.Controller) storage.BuildingsStorage {
				return storage_mocks.NewBuildingsStorageMock(mc).
					CreateBuildingMock.
					Expect(minimock.AnyContext, &models.Building{
						Name:    "building_1",
						Address: null.String{Valid: true, String: "address_1"},
					}).
					Return(nil)
			},
		},
		{
			name: "storageError",
			args: args{
				building: &models.Building{
					ID:      1,
					Name:    "building_1",
					Address: null.String{Valid: true, String: "address_1"},
				},
			},
			getBuildingsStorage: func(mc *minimock.Controller) storage.BuildingsStorage {
				return storage_mocks.NewBuildingsStorageMock(mc).
					CreateBuildingMock.
					Expect(minimock.AnyContext, &models.Building{
						ID:      1,
						Name:    "building_1",
						Address: null.String{Valid: true, String: "address_1"},
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
			buildingsStorage := tt.getBuildingsStorage(mc)
			s := Service{buildingsStorage: buildingsStorage}

			err := s.CreateBuilding(context.Background(), tt.args.building)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
		})
	}
}

func Test_DeleteBuilding(t *testing.T) {
	t.Parallel()

	type args struct {
		id int
	}

	tests := []struct {
		name                string
		args                args
		getBuildingsStorage func(mc *minimock.Controller) storage.BuildingsStorage
		wantErr             bool
	}{
		{
			name: "valid",
			args: args{
				id: 1,
			},
			getBuildingsStorage: func(mc *minimock.Controller) storage.BuildingsStorage {
				return storage_mocks.NewBuildingsStorageMock(mc).
					DeleteBuildingMock.
					Expect(minimock.AnyContext, 1).
					Return(1, nil)
			},
		},
		{
			name: "wrongID",
			args: args{
				id: 0,
			},
			getBuildingsStorage: func(mc *minimock.Controller) storage.BuildingsStorage {
				return nil
			},
			wantErr: true,
		},
		{
			name: "noRowToDelete",
			args: args{
				id: 2,
			},
			getBuildingsStorage: func(mc *minimock.Controller) storage.BuildingsStorage {
				return storage_mocks.NewBuildingsStorageMock(mc).
					DeleteBuildingMock.
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
			getBuildingsStorage: func(mc *minimock.Controller) storage.BuildingsStorage {
				return storage_mocks.NewBuildingsStorageMock(mc).
					DeleteBuildingMock.
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
			buildingsStorage := tt.getBuildingsStorage(mc)
			s := Service{buildingsStorage: buildingsStorage}

			err := s.DeleteBuilding(context.Background(), tt.args.id)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
		})
	}
}
