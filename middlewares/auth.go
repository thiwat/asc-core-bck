package middlewares

import (
	"asc-core/db"
	"asc-core/types"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	fiber "github.com/gofiber/fiber/v2"
)

func ValidateToken() fiber.Handler {
	return func(c *fiber.Ctx) error {

		tokens := c.GetReqHeaders()["Authorization"]

		if len(tokens) < 1 {
			return c.Status(http.StatusBadRequest).JSON(types.ErrorResponse{
				Code:    "invalid_request",
				Message: "Access Denided",
			})
		}

		token := tokens[0]
		token = strings.Replace(token, "Bearer ", "", 1)
		token = strings.Replace(token, "TOKEN_", "", 1)

		var session types.Session

		sessionData, _ := db.GetKey("TOKEN_" + token)

		err := json.Unmarshal([]byte(sessionData), &session)

		fmt.Println(session)

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(types.ErrorResponse{
				Code:    "invalid_request",
				Message: "Access Denided",
			})
		}

		c.Locals("session", session)

		return c.Next()
	}
}
