package server

import (
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Participants is a struct that holds the websocket connection and the host info
type Participant struct {
	Conn *websocket.Conn
	Host bool
}

// Room is a struct that holds the room name and the type Participants mapped to it
type Room struct {
	RoomMap map[string][]Participant
	Mutex   sync.RWMutex
}

// inits room
func (r *Room) Init() {
	r.RoomMap = make(map[string][]Participant)
	log.Println("new room init")
}

// create a new room
func (r *Room) CreateRoom() string {

	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	//to create a room id like google-meet
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, 6)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	roomId := string(b)

	new_room := make([]Participant, 0)
	r.RoomMap[roomId] = new_room

	log.Println("Created Room with RoomID: ", roomId)
	return roomId
}

// get participants in a room
func (r *Room) GetParticipants(roomId string) []Participant {

	r.Mutex.RLock()
	defer r.Mutex.RUnlock()

	return r.RoomMap[roomId]
}

// add new participants in a room
func (r *Room) AddParticipants(roomId string, host bool, conn *websocket.Conn) {

	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	participant := Participant{
		Conn: conn,
		Host: host,
	}

	room, err := r.RoomMap[roomId]
	//return error if no room
	if err == false {
		log.Println("No room to add participant")
		return
	}

	log.Println("Inserting into Room with RoomID: ", roomId)
	room = append(room, participant)
}

// delete room
func (r *Room) DeleteRoom(roomId string) {

	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	delete(r.RoomMap, roomId)
}
