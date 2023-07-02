package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

var AllRooms Room

func CreateRoomHandler(c *fiber.Ctx) error {

	type resp struct {
		RoomID string `json:"room_id"`
	}

	roomId := AllRooms.CreateRoom()

	log.Println(AllRooms.RoomMap)
	c.JSON(resp{RoomID: roomId})

	return nil
}

func JoinRoomHandler(c *fiber.Ctx) error {
	return nil
}
