package bms

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sotskov-do/oms-assignment/internal/models"
)

func (bms *BuildingManagementSystem) GetApartmentsHandler(c *fiber.Ctx) error {
	apartments, err := bms.apartmentsService.GetApartments(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(&fiber.Map{
				resultKey:   resultError,
				responseKey: err.Error(),
			})
	}

	return c.JSON(&fiber.Map{
		resultKey:   resultSuccess,
		responseKey: apartments,
	})
}

func (bms *BuildingManagementSystem) GetApartmentHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(&fiber.Map{
				resultKey:   resultError,
				responseKey: err.Error(),
			})
	}

	apartment, err := bms.apartmentsService.GetApartment(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(&fiber.Map{
				resultKey:   resultError,
				responseKey: err.Error(),
			})
	}

	return c.JSON(&fiber.Map{
		resultKey:   resultSuccess,
		responseKey: apartment,
	})
}

func (bms *BuildingManagementSystem) GetApartmentsInBuildingHandler(c *fiber.Ctx) error {
	buildingId, err := c.ParamsInt("buildingId", 0)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(&fiber.Map{
				resultKey:   resultError,
				responseKey: err.Error(),
			})
	}

	apartmentsInBuilding, err := bms.apartmentsService.GetApartmentsInBuilding(c.Context(), buildingId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(&fiber.Map{
				resultKey:   resultError,
				responseKey: err.Error(),
			})
	}

	return c.JSON(&fiber.Map{
		resultKey:   resultSuccess,
		responseKey: apartmentsInBuilding,
	})
}

func (bms *BuildingManagementSystem) CreateApartmentHandler(c *fiber.Ctx) error {
	var apartment *models.Apartment
	err := c.BodyParser(&apartment)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(&fiber.Map{
				resultKey:   resultError,
				responseKey: err.Error(),
			})
	}

	err = bms.apartmentsService.CreateApartment(c.Context(), apartment)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(&fiber.Map{
				resultKey:   resultError,
				responseKey: err.Error(),
			})
	}

	return c.JSON(&fiber.Map{
		resultKey: resultSuccess,
	})
}

func (bms *BuildingManagementSystem) DeleteApartmentHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(&fiber.Map{
				resultKey:   resultError,
				responseKey: err.Error(),
			})
	}

	err = bms.apartmentsService.DeleteApartment(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(&fiber.Map{
				resultKey:   resultError,
				responseKey: err.Error(),
			})
	}

	return c.JSON(&fiber.Map{
		resultKey: resultSuccess,
	})
}
