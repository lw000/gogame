package desk

import (
	"fmt"
	"gogame/user"
	"sync"
)

type Desk struct {
	TableId     int
	PlayerCount int
	user        []user.User
	m           sync.Mutex
}

type IDesk interface {
	OnSitUp(chairID int, user user.User)
	OnSitDown(chairID int, user user.User)

	onFrameMessage(v interface{}) bool

	OnGameMessage(v interface{}) bool
}

func (d *Desk) OnSitUp(chairID int, user user.User) {
	fmt.Println(chairID, user)
}

func (d *Desk) OnSitDown(chairID int, user user.User) {
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
