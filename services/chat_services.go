package services

import "sync"

var ChatServicesInstance *ChatServices

func init() {
	// 这个地方使用sync.Once来实现单例模式也可以
	ChatServicesInstance = &ChatServices{
		ChatMap: make(map[string]*Chat),
		RoomNum: 0,
	}
}

type ChatServices struct {
	ChatMap map[string]*Chat
	RoomNum int32
	rw      sync.RWMutex
}

// GetRoom 获取房间实例
func (cs *ChatServices) GetRoom(roomId string, isCreate bool) *Chat {
	cs.rw.RLock()
	room, ok := cs.ChatMap[roomId]
	cs.rw.RUnlock()
	if !ok && isCreate {
		room = cs.AddRoom(roomId)
	}
	return room
}

// AddRoom 新增房间
func (cs *ChatServices) AddRoom(roomId string) *Chat {
	var room = &Chat{
		RoomId: roomId,
		UserList: make(map[string]*User),
		UserNum: 0,
	}
	cs.rw.Lock()
	cs.ChatMap[roomId] = room
	cs.RoomNum++
	cs.rw.Unlock()
	return room
}

// RemoveRoom 删除房间
func (cs *ChatServices) RemoveRoom(roomId string)  {
	cs.rw.Lock()
	delete(cs.ChatMap, roomId)
	cs.RoomNum--
	cs.rw.Unlock()
}