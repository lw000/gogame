package user

import (
	"demo/gogame/network"
	"sync"
)

type User struct {
	UserId         int64  //用户Id
	UserName       string //用户名
	RoomId         string //房间Id
	GameId         int64  //游戏Id
	DeskId         int64  //桌子Id
	Status         int    //用户状态
	LoginTimestamp int64  //登录时间

	ip string

	client *network.Client
	m      sync.RWMutex
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

func (u *User) GetClient() *network.Client {
	u.m.RLock()
	defer u.m.RUnlock()

	return u.client
}

func (u *User) AttachClient(client *network.Client) {
	u.m.RLock()
	defer u.m.RUnlock()

	u.client = client
}

func (u *User) DetachClient() {
	u.m.RLock()
	defer u.m.RUnlock()

	u.client = nil
}
