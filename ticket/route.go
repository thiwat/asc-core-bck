package ticket

import (
	"net/http"

	"asc-core/middlewares"
	"asc-core/types"

	fiber "github.com/gofiber/fiber/v2"
)

func RestRouteV1(router fiber.Router) {
	ticketGroup := router.Group("ticket")

	ticketGroup.Use(middlewares.ValidateToken())

	ticketGroup.Get("/", func(c *fiber.Ctx) error {

		session := c.Locals("session").(types.Session)

		page := int64(c.QueryInt("page", 1))
		pageSize := int64(c.QueryInt("page_size", 10))
		sort := c.Query("sort", "")

		res, err := ListTicketByUser(page, pageSize, sort, session)

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(types.ErrorResponse{
				Code:    "list_error",
				Message: err.Error(),
			})
		}

		return c.Status(http.StatusOK).JSON(res)
	})

	ticketGroup.Get("/:event", func(c *fiber.Ctx) error {

		session := c.Locals("session").(types.Session)
		event := c.Params("event")

		res, err := FindTicketByUser(event, session)

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(types.ErrorResponse{
				Code:    "get_error",
				Message: err.Error(),
			})
		}

		return c.Status(http.StatusOK).JSON(res)
	})

	ticketGroup.Post("/", func(c *fiber.Ctx) error {

		session := c.Locals("session").(types.Session)

		var input PurchaseTicketInput
		if err := c.BodyParser(&input); err != nil {
			return c.Status(http.StatusBadRequest).JSON(types.ErrorResponse{
				Code:    "invalid_request",
				Message: err.Error(),
			})
		}

		res, err := PurchaseTicket(input, session)

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(types.ErrorResponse{
				Code:    "invalid_request",
				Message: err.Error(),
			})
		}

		return c.Status(http.StatusOK).JSON(res)
	})
}
