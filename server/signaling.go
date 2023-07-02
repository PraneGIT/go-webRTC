package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

var AllRooms Room

func CreateRoomHandler(w http.ResponseWriter, r *http.Request) {

	type resp struct {
		RoomID string `json:"room_id"`
	}

	roomId := AllRooms.CreateRoom()

	log.Println(AllRooms.RoomMap)
	json.NewEncoder(w).Encode(resp{RoomID: roomId})
}

var upgrader = websocket.Upgrader{
	// Allow origin connections from any domain.
	// This is useful for localhost development.
	CheckOrigin: func(r *http.Request) bool { return true },
}

func JoinRoomHandler(w http.ResponseWriter, r *http.Request) {

	// get room id from url params
	p := strings.Split(r.URL.Path, "/")
	id := p[0]

	// check if room exists
	if _, ok := AllRooms.RoomMap[id]; !ok {
		log.Println("Room does not exist", id)
		return
	}

	// upgrade connection to websocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// add participant to room
	AllRooms.AddParticipants(id, false, conn)

}
