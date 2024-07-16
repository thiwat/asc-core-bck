package main

import (
	"log"

	"asc-core/configs"
	"asc-core/event"
	"asc-core/ticket"
	"asc-core/user"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	api := app.Group("/api")
	v1 := api.Group("/v1")

	event.RestRouteV1(v1)
	user.RestRouteV1(v1)
	ticket.RestRouteV1(v1)

	log.Fatal(app.Listen(configs.GetPort()))
}
