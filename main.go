package main

import (
	"github.com/WeslleyRibeiro-1999/api-star-wars/db"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	// ROTAS CRIADAS
	app.Get("/planets", db.GetPlanetID)
	app.Get("/planet/:id?", db.GetPlanetID)
	app.Post("/planet", db.CreatePlanet)
	app.Delete("/planet/:id", db.DeletePlanet)

	app.Listen(db.Port)
}
