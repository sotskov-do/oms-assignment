package bms

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sotskov-do/oms-assignment/internal/models"
)

func (bms *BuildingManagementSystem) GetBuildingsHandler(c *fiber.Ctx) error {
	buildings, err := bms.buildingsService.GetBuildings(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(&fiber.Map{
				resultKey:   resultError,
				responseKey: err.Error(),
			})
	}

	return c.JSON(&fiber.Map{
		resultKey:   resultSuccess,
		responseKey: buildings,
	})
}

func (bms *BuildingManagementSystem) GetBuildingHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(&fiber.Map{
				resultKey:   resultError,
				responseKey: err.Error(),
			})
	}

	building, err := bms.buildingsService.GetBuilding(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(&fiber.Map{
				resultKey:   resultError,
				responseKey: err.Error(),
			})
	}

	return c.JSON(&fiber.Map{
		resultKey:   resultSuccess,
		responseKey: building,
	})
}

func (bms *BuildingManagementSystem) CreateBuildingHandler(c *fiber.Ctx) error {
	var building *models.Building
	err := c.BodyParser(&building)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(&fiber.Map{
				resultKey:   resultError,
				responseKey: err.Error(),
			})
	}

	err = bms.buildingsService.CreateBuilding(c.Context(), building)
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

func (bms *BuildingManagementSystem) DeleteBuildingHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(&fiber.Map{
				resultKey:   resultError,
				responseKey: err.Error(),
			})
	}

	err = bms.buildingsService.DeleteBuilding(c.Context(), id)
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
