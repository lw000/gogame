package platform

import (
	"fmt"
	"gogame/user"
	"sync"
)

type Platform struct {
	Id       int64
	Name     string
	Rooms    []*Room
	m        sync.Mutex
	mgrServe *user.Manager
}

func NewPlatform(Id int64, name string) *Platform {
	return &Platform{
		Id:       Id,
		Name:     name,
		mgrServe: user.Service(),
	}
}

func (p *Platform) Start() error {
	p.CreateRoom()

	return nil
}

func (p *Platform) Stop() error {

	return nil
}

func (p *Platform) CreateRoom() {
	for i := 0; i < 100; i++ {
		p.Rooms = append(p.Rooms, NewRoom(i+1, fmt.Sprintf("room_%d", i+1)))
	}
}

func (p *Platform) DestroyRoom() {

}
