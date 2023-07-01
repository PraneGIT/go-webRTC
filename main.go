package main

import (
	"github.com/PraneGIT/go-webRTC/server"
	"github.com/gofiber/fiber/v2"
)

func main() {

	server.AllRooms.Init()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, webRTC-go 👋!")
	})

	api := app.Group("/webRTC")
	api.Get("/create", server.CreateRoomHandler)

	api.Get("/join:roomId", server.JoinRoomHandler)

	app.Listen(":3000")
}
