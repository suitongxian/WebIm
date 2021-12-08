package services

import (
	"net/http"

	"github.com/astaxie/beego/core/logs"
	"github.com/gorilla/websocket"
)

func UpgradeReq(w http.ResponseWriter, r *http.Request) *websocket.Conn {
	ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "Not a websocket handshake", 400)
		return nil
	} else if err != nil {
		logs.Error("Cannot setup WebSocket connection:", err)
		return nil
	}
	return ws
}
