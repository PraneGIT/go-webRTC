package main

import (
	"net/http"

	"github.com/PraneGIT/go-webRTC/server"
)

func main() {

	server.AllRooms.Init()

	http.HandleFunc("/create", server.CreateRoomHandler)

	http.HandleFunc("/join", server.JoinRoomHandler)

	println("server running on port 3000...")
	http.ListenAndServe(":3000", nil)

}
