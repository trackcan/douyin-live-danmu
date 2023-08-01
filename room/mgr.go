package room

import (
	"strings"
	"sync"
)

type RoomMgr struct {
	m      map[string]*Room
	locker sync.RWMutex
}

func InitRoomMgr() *RoomMgr {
	mgr := RoomMgr{}
	mgr.init()
	return &mgr
}

func (p *RoomMgr) init() {
	p.m = make(map[string]*Room)
}

func (p *RoomMgr) Add(liveUrl, reportUrl string) (*Room, error) {
	roomID := parseRoomIDFromUrl(liveUrl)
	p.locker.RLock()
	defer p.locker.RUnlock()

	if value, ok := p.m[roomID]; ok {
		return value, nil
	}

	r, err := NewRoom(liveUrl,reportUrl)
	if err != nil {
		return nil, err
	}
	r.DisconnectedCB = p.disconnectedCB
	_ = r.Connect()
	p.m[roomID] = r
	return r, nil
}

func (p *RoomMgr) Get(roomID string) *Room {
	return p.m[roomID]
}

func (p *RoomMgr) disconnectedCB(roomID string) {
	p.Remove(roomID)
}

func (p *RoomMgr) Remove(liveUrl string) {
	roomID := parseRoomIDFromUrl(liveUrl)
	p.locker.RLock()
	defer p.locker.RUnlock()
	if value, ok := p.m[roomID]; ok {
		value.Close()
		delete(p.m, roomID)
		return
	}
}

func parseRoomIDFromUrl(url string) string {
	n := strings.LastIndex(url, "/")
	if n > 0 {
		url = url[n+1:]
	}
	n = strings.Index(url, "?")
	if n > 0 {
		url = url[:n]
	}
	return url
}
