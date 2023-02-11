package controller

import (
	"app-test-fiber/entity"
	"app-test-fiber/model"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetAll(c *fiber.Ctx) error {

	data := map[string]string{
		"name":  "Julhan",
		"email": "kamapsi.@emai.com",
	}

	return c.JSON(data)
}

func GetById(c *fiber.Ctx) error {

	id := c.Params("id")
	name := c.Query("name")

	data := entity.UserEntity{
		ID:    id,
		Name:  name,
		Email: "test@emai.com",
	}

	return c.JSON(data)
}

func Create(c *fiber.Ctx) error {

	var userBody entity.UserEntity

	if err := c.BodyParser(&userBody); err != nil {
		return c.JSON(entity.Response{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		})
	}

	user, err := model.Create(&userBody)
	if err != nil {
		return c.JSON(entity.Response{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		})
	}

	return c.JSON(entity.Response{
		StatusCode: http.StatusCreated,
		Data:       user,
		Message:    "success",
	})

}
