package services

import (
	"WebIm/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"sync"
)

type Chat struct {
	RoomId   string
	UserList map[string]*User
	UserNum  int32
	rw       sync.RWMutex
}

// Add 加入房间
func (c *Chat) Add(username string, ws *websocket.Conn) {
	c.rw.RLock()
	user, ok := c.UserList[username]
	c.rw.RUnlock()
	if ok {
		fmt.Println("user is old")
		return
	}
	user = &User{Username: username, Ws: ws, RoomId: c.RoomId}
	go user.init()
	c.rw.Lock()
	c.UserList[username] = user
	c.UserNum++
	c.rw.Unlock()
}

// Leave 离开房间
func (c *Chat) Leave(username string) {
	c.rw.RLock()
	user, ok := c.UserList[username]
	c.rw.RUnlock()
	if !ok {
		return
	}
	user.Ws.Close()
	c.rw.Lock()
	delete(c.UserList, username)
	c.UserNum--
	c.rw.Unlock()
}

// Broadcast 广播消息
func (c *Chat) Broadcast(event models.Event) {
	data, err := json.Marshal(event)
	if err != nil {
		return
	}
	for k, v := range c.UserList {
		ws := v.Ws
		if ws != nil {
			if ws.WriteMessage(websocket.TextMessage, data) != nil {
				fmt.Println("send message fail")
				c.Leave(k)
			}
		}
	}
}

// PushByUname 根据username发送消息
func (c *Chat) PushByUname()  {
	//
}
