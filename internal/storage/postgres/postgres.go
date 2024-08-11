package postgres

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/sotskov-do/oms-assignment/internal/models"
)

type PostgresDatabase struct {
	psqlClient *sql.DB
}

func New(ctx context.Context, pgconn string) (*PostgresDatabase, error) {
	db, err := sql.Open("postgres", pgconn+"?sslmode=disable")
	if err != nil {
		return nil, err
	}

	return &PostgresDatabase{psqlClient: db}, nil
}

func (pdb *PostgresDatabase) Ping(ctx context.Context) error {
	err := pdb.psqlClient.Ping()
	if err != nil {
		return err
	}
	return nil
}

func (pdb *PostgresDatabase) Stop(ctx context.Context) error {
	err := pdb.psqlClient.Close()
	if err != nil {
		return err
	}
	return nil
}

/* Apartments */

func (pdb *PostgresDatabase) GetApartments(ctx context.Context) (models.ApartmentSlice, error) {
	a, err := models.Apartments().All(ctx, pdb.psqlClient)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (pdb *PostgresDatabase) GetApartment(ctx context.Context, id int) (*models.Apartment, error) {
	a, err := models.Apartments(qm.Where("id=?", id)).One(ctx, pdb.psqlClient)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (pdb *PostgresDatabase) GetApartmentsInBuilding(ctx context.Context, buildingId int) (models.ApartmentSlice, error) {
	a, err := models.Apartments(qm.Where("building_id=?", buildingId)).All(ctx, pdb.psqlClient)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (pdb *PostgresDatabase) CreateApartment(ctx context.Context, apartment *models.Apartment) error {
	err := apartment.Upsert(ctx, pdb.psqlClient, true, []string{}, boil.Infer(), boil.Infer())
	if err != nil {
		return err
	}

	return nil
}

func (pdb *PostgresDatabase) DeleteApartment(ctx context.Context, id int) (int64, error) {
	n, err := models.Apartments(qm.Where("id=?", id)).DeleteAll(ctx, pdb.psqlClient)
	if err != nil {
		return 0, err
	}

	return n, nil
}

/* Buildings */

func (pdb *PostgresDatabase) GetBuildings(ctx context.Context) (models.BuildingSlice, error) {
	b, err := models.Buildings().All(ctx, pdb.psqlClient)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (pdb *PostgresDatabase) GetBuilding(ctx context.Context, id int) (*models.Building, error) {
	b, err := models.Buildings(qm.Where("id=?", id)).One(ctx, pdb.psqlClient)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (pdb *PostgresDatabase) CreateBuilding(ctx context.Context, building *models.Building) error {
	err := building.Upsert(ctx, pdb.psqlClient, true, []string{}, boil.Infer(), boil.Infer())
	if err != nil {
		return err
	}

	return nil
}

func (pdb *PostgresDatabase) DeleteBuilding(ctx context.Context, id int) (int64, error) {
	n, err := models.Buildings(qm.Where("id=?", id)).DeleteAll(ctx, pdb.psqlClient)
	if err != nil {
		return 0, err
	}

	return n, nil
}
