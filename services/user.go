package services

import (
	"WebIm/models"
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

type User struct {
	Username string
	Ws *websocket.Conn
	RoomId string
}

func (u *User) init()  {
	for {
		_, msg, err := u.Ws.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		ChatServicesInstance.GetRoom(u.RoomId, false).Broadcast(newEvent(models.EVENT_MESSAGE, u.Username, string(msg)))
	}
}

func newEvent(ep models.EventType, user, msg string) models.Event {
	return models.Event{ep, user, int(time.Now().Unix()), msg}
}
