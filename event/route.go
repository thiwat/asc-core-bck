package event

import (
	"net/http"

	"asc-core/types"

	fiber "github.com/gofiber/fiber/v2"
)

func RestRouteV1(router fiber.Router) {
	eventGroup := router.Group("event")

	eventGroup.Get("/", func(c *fiber.Ctx) error {
		page := int64(c.QueryInt("page", 1))
		pageSize := int64(c.QueryInt("page_size", 10))
		sort := c.Query("sort", "")

		res, err := ListEvent(page, pageSize, sort)

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(types.ErrorResponse{
				Code:    "list_error",
				Message: err.Error(),
			})
		}

		return c.Status(http.StatusOK).JSON(res)
	})

	eventGroup.Get("/:code", func(c *fiber.Ctx) error {

		code := c.Params("code")
		var event Event
		if err := c.BodyParser(&event); err != nil {
			return c.Status(http.StatusBadRequest).JSON(types.ErrorResponse{
				Code:    "invalid_request",
				Message: err.Error(),
			})
		}

		res, err := FindByCode(code)

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(types.ErrorResponse{
				Code:    "invalid_request",
				Message: err.Error(),
			})
		}

		return c.Status(http.StatusOK).JSON(res)
	})

	eventGroup.Post("/", func(c *fiber.Ctx) error {

		var event Event
		if err := c.BodyParser(&event); err != nil {
			return c.Status(http.StatusBadRequest).JSON(types.ErrorResponse{
				Code:    "invalid_request",
				Message: err.Error(),
			})
		}

		res, err := CreateEvent(event)

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(types.ErrorResponse{
				Code:    "invalid_request",
				Message: err.Error(),
			})
		}

		return c.Status(http.StatusOK).JSON(res)
	})
}
