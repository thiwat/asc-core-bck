package user

import (
	"net/http"

	"asc-core/middlewares"
	"asc-core/types"

	fiber "github.com/gofiber/fiber/v2"
)

func RestRouteV1(router fiber.Router) {
	authGroup := router.Group("auth")

	authGroup.Post("/login", func(c *fiber.Ctx) error {

		var input LoginInput

		if err := c.BodyParser(&input); err != nil {
			return c.Status(http.StatusBadRequest).JSON(types.ErrorResponse{
				Code:    "invalid_request",
				Message: err.Error(),
			})
		}

		res, err := Login(input)

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(types.ErrorResponse{
				Code:    "invalid_request",
				Message: err.Error(),
			})
		}

		return c.Status(http.StatusOK).JSON(res)
	})

	userGroup := router.Group("me")
	userGroup.Use(middlewares.ValidateToken())

	userGroup.Get("/", func(c *fiber.Ctx) error {

		session := c.Locals("session").(types.Session)

		profile, err := GetProfile(session)

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(types.ErrorResponse{
				Code:    "invalid_request",
				Message: err.Error(),
			})
		}

		return c.Status(http.StatusOK).JSON(profile)
	})
}
