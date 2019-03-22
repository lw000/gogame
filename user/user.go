package user

import (
	"sync"
)

type User struct {
	UserId         int64 //用户Id
	Alisid         int64
	UserName       string //用户名
	Nickname       string
	RoomId         string //房间Id
	GameId         int64  //游戏Id
	DeskId         int64  //桌子Id
	Status         int    //用户状态
	LoginTimestamp int64  //登录时间

	ip string

	m sync.RWMutex
}

func NewUser() *User {
	return &User{}
}

func (u *User) Ip() string {
	return u.ip
}

func (u *User) SetIp(ip string) {
	u.ip = ip
}
