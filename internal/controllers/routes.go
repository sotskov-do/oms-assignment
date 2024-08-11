package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sotskov-do/oms-assignment/internal/controllers/bms"
)

func SetupRoutes(
	app *fiber.App,
	bms *bms.BuildingManagementSystem,
) {
	app.Route("/buildings", func(api fiber.Router) {
		// GET /buildings: List all buildings (with or without the apartments)
		api.Get("/", bms.GetBuildingsHandler).Name("getAll")
		// GET /buildings/{id}: Get a single building by ID
		api.Get("/:id", bms.GetBuildingHandler).Name("getByID")
		// POST /buildings: Create a new building (update if already exist)
		api.Post("/", bms.CreateBuildingHandler).Name("create")
		// DELETE /buildings/{id}: Delete a building by ID
		api.Delete("/:id", bms.DeleteBuildingHandler).Name("delete")
	}, "buildings.")

	app.Route("/apartments", func(api fiber.Router) {
		// GET /apartments: List all apartments
		api.Get("/", bms.GetApartmentsHandler).Name("getAll")
		// GET /apartments/{id}: Get a single apartment by ID
		api.Get("/:id", bms.GetApartmentHandler).Name("getByID")
		// GET /apartments/building/{buildingId}: Get all apartments in a specific building
		api.Get("/building/:buildingId", bms.GetApartmentsInBuildingHandler).Name("getAllInBuilding")
		// POST /apartments: Create a new apartment (update if already exist)
		api.Post("/", bms.CreateApartmentHandler).Name("create")
		// DELETE /apartments/{id}: Delete an apartment by ID
		api.Delete("/:id", bms.DeleteApartmentHandler).Name("delete")
	}, "apartments.")
}
