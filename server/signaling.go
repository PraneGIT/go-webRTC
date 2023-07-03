package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var AllRooms Room

func CreateRoomHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

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

type BrodcastMsg struct {
	Message map[string]interface{}
	RoomId  string
	Client  *websocket.Conn
}

// make channel for brodcast
var broadcast = make(chan BrodcastMsg)

// brodcast message to all participants in a room
func brodcaster() {
	for {
		msg := <-broadcast
		for _, participant := range AllRooms.RoomMap[msg.RoomId] {
			if participant.Conn != msg.Client {
				err := participant.Conn.WriteJSON(msg.Message)
				if err != nil {
					log.Println(err)
					return
				}
			}
		}
	}
}

func JoinRoomHandler(w http.ResponseWriter, r *http.Request) {

	// get room id from url params
	roomid, ok := r.URL.Query()["roomID"]
	if !ok || len(roomid[0]) < 1 {
		log.Println("Url Param 'roomID' is missing")
		return
	}
	id := roomid[0]
	log.Println(id)

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

	go brodcaster()

	for {
		var brodcastMsg BrodcastMsg

		err = conn.ReadJSON(&brodcastMsg.Message)
		if err != nil {
			log.Fatal("Read Error:", err)
		}

		brodcastMsg.RoomId = id
		brodcastMsg.Client = conn
		broadcast <- brodcastMsg
	}

}
