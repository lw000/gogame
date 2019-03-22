package user

import "sync"

type UserManager struct {
	store sync.Map
}

var (
	once sync.Once
	umgr *UserManager
)

func UserMgr() *UserManager {
	once.Do(func() {
		umgr = newUserManager()
	})

	return umgr
}

func newUserManager() *UserManager {
	return &UserManager{}
}

func (um *UserManager) AddWith(u *User) bool {
	_, ok := um.store.Load(u.UserId)
	if !ok {
		um.store.Store(u.UserId, u)
	}
	return true
}

func (um *UserManager) Remove(userId int64) {
	um.store.Delete(userId)
}

func (um *UserManager) RemoveWith(u *User) {
	um.Remove(u.UserId)
}

func (um *UserManager) Get(userId int64) *User {
	v, ok := um.store.Load(userId)
	if ok {
		return v.(*User)
	}
	return nil
}
