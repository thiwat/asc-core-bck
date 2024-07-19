package order

import (
	"net/http"

	"asc-core/middlewares"
	"asc-core/types"

	fiber "github.com/gofiber/fiber/v2"
)

func RestRouteV1(router fiber.Router) {
	orderGroup := router.Group("order")

	orderGroup.Use(middlewares.ValidateToken())

	orderGroup.Get("/me", func(c *fiber.Ctx) error {

		session := c.Locals("session").(types.Session)

		page := int64(c.QueryInt("page", 1))
		pageSize := int64(c.QueryInt("page_size", 10))
		sort := c.Query("sort", "")

		res, err := ListByUser(page, pageSize, sort, session)

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(types.ErrorResponse{
				Code:    "list_error",
				Message: err.Error(),
			})
		}

		return c.Status(http.StatusOK).JSON(res)
	})

	orderGroup.Get("/:order_no", func(c *fiber.Ctx) error {

		session := c.Locals("session").(types.Session)
		orderNo := c.Params("order_no")

		res, err := GetOrder(orderNo, session)

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(types.ErrorResponse{
				Code:    "get_error",
				Message: err.Error(),
			})
		}

		return c.Status(http.StatusOK).JSON(res)
	})

	orderGroup.Post("/", func(c *fiber.Ctx) error {

		session := c.Locals("session").(types.Session)
		var input PlaceOrderInput

		if err := c.BodyParser(&input); err != nil {
			return c.Status(http.StatusBadRequest).JSON(types.ErrorResponse{
				Code:    "invalid_request",
				Message: err.Error(),
			})
		}

		res, err := PlaceOrder(input, session)

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(types.ErrorResponse{
				Code:    "invalid_request",
				Message: err.Error(),
			})
		}

		return c.Status(http.StatusOK).JSON(res)
	})

	orderGroup.Patch("/slip", func(c *fiber.Ctx) error {

		session := c.Locals("session").(types.Session)
		var input UploadSlipInput

		if err := c.BodyParser(&input); err != nil {
			return c.Status(http.StatusBadRequest).JSON(types.ErrorResponse{
				Code:    "invalid_request",
				Message: err.Error(),
			})
		}

		res, err := UploadSlip(input, session)

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(types.ErrorResponse{
				Code:    "invalid_request",
				Message: err.Error(),
			})
		}

		return c.Status(http.StatusOK).JSON(res)
	})

	orderGroup.Patch("/approve", func(c *fiber.Ctx) error {

		var input ApprovePaymentInput

		if err := c.BodyParser(&input); err != nil {
			return c.Status(http.StatusBadRequest).JSON(types.ErrorResponse{
				Code:    "invalid_request",
				Message: err.Error(),
			})
		}

		res, err := ApprovePayment(input)

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(types.ErrorResponse{
				Code:    "invalid_request",
				Message: err.Error(),
			})
		}

		return c.Status(http.StatusOK).JSON(res)

	})
}
