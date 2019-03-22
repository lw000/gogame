package platform

import (
	"demo/gogame/room"
	"fmt"
	"sync"
)

type Platform struct {
	Pid   int
	Name  string
	Rooms []*room.Room
	m     sync.Mutex
}

func NewPlatform(pid int, name string) *Platform {
	return &Platform{
		Pid:  pid,
		Name: name,
	}
}

func (p *Platform) CreateRoom() {
	for i := 0; i < 100; i++ {
		p.Rooms = append(p.Rooms, room.NewRoom(i+1, fmt.Sprintf("room_%d", i+1)))
	}
}

func (p *Platform) AddRoom(r *room.Room) {

}

func (p *Platform) RemoveRoom(rid int) {

}

func (p *Platform) DestroyRoom() {

}
