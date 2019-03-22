package main

import (
	"demo/gogame/cmd/platformsrv/room"
	"demo/gogame/user"
	"fmt"
	"log"
	"sync"
)

type Platform struct {
	Id    int64
	Name  string
	Rooms []*room.Room
	m     sync.Mutex
	usrv  *user.Manager
}

func NewPlatform(Id int64, name string) *Platform {
	return &Platform{
		Id:   Id,
		Name: name,
		usrv: user.Service(),
	}
}

func (p *Platform) Start() error {

	return nil
}

func (p *Platform) Stop() error {

	return nil
}

func (p *Platform) CreateRoom() {
	for i := 0; i < 100; i++ {
		p.Rooms = append(p.Rooms, room.NewRoom(i+1, fmt.Sprintf("room_%d", i+1)))
	}
}

func (p *Platform) DestroyRoom() {

}

func main() {
	platform := NewPlatform(1, "魔游无线游戏")
	if er := platform.Start(); er != nil {
		log.Println(er)
		return
	}

	platform.CreateRoom()
}
