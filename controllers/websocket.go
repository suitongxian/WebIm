package controllers

import (
	"WebIM/services"
	"github.com/astaxie/beego/core/logs"
)

// WebSocketController handles WebSocket requests.
type WebSocketController struct {
	baseController
}

// Get method handles GET requests for WebSocketController.
func (this *WebSocketController) Get() {
	uname := this.GetString("uname")
	roomId := this.GetString("room_id")
	if len(uname) == 0 {
		this.Redirect("/", 302)
		return
	}

	this.TplName = "websocket.html"
	this.Data["IsWebSocket"] = true
	this.Data["UserName"] = uname
	this.Data["RoomId"] = roomId
}

func (this *WebSocketController) Join() {
	uname := this.GetString("uname")
	roomId := this.GetString("room_id")
	if len(uname) == 0 {
		this.Redirect("/", 302)
		return
	}
	ws := services.UpgradeReq(this.Ctx.ResponseWriter, this.Ctx.Request)
	if ws == nil {
		logs.Error("websocket fail")
		return
	}

	var chat *services.Chat
	chat = services.ChatServicesInstance.GetRoom(roomId, true)
	chat.Add(uname, ws)
}
