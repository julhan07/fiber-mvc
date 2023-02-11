package routes

import (
	"app-test-fiber/controller"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {

	v1 := app.Group("/v1")
	v1.Get("/user", controller.GetAll)
	v1.Post("/user", controller.Create)
	v1.Get("/user/:id", controller.GetById)
	v1.Put("/user/:id", controller.GetAll)
	v1.Delete("/user/:id", controller.GetAll)

}
