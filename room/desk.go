package room

import (
	"demo/gogame/user"
	"fmt"
	"sync"
)

type Desk struct {
	TableId     int
	PlayerCount int
	user        []user.User
	m           sync.Mutex
}

type IDesk interface {
	OnSitup(chairID int, user user.User)
	OnSitdown(chairID int, user user.User)

	onFrameMessage(v interface{}) bool

	OnGameMessage(v interface{}) bool
}

func (d *Desk) OnSitup(chairID int, user user.User) {
	fmt.Println(chairID, user)
}

func (d *Desk) OnSitdown(chairID int, user user.User) {
	fmt.Println(chairID, user)
}

func (d *Desk) onFrameMessage(v interface{}) bool {
	fmt.Println(v)
	return true
}

func (d *Desk) OnGameMessage(v interface{}) bool {
	fmt.Println(v)
	return true
}
