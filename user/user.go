package user

type User struct {
	UserId         int64  // 用户Id
	AlisId         int64  // 用户Id
	UserName       string // 用户名
	Nickname       string // 用户昵称
	RoomId         string // 房间Id
	GameId         int64  // 游戏Id
	DeskId         int64  // 桌子Id
	Status         int    // 用户状态
	LoginTimestamp int64  // 登录时间

	ip string
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
