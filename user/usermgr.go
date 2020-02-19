package user

import "sync"

type Manager struct {
	store sync.Map
}

var (
	once    sync.Once
	umserve *Manager
)

func init() {
	Service()
}

func Service() *Manager {
	once.Do(func() {
		umserve = &Manager{}
	})
	return umserve
}

func (um *Manager) Add(u *User) bool {
	_, ok := um.store.Load(u.UserId)
	if !ok {
		um.store.Store(u.UserId, u)
	}
	return true
}

func (um *Manager) Remove(userId int64) {
	um.store.Delete(userId)
}

func (um *Manager) RemoveWith(u *User) {
	um.Remove(u.UserId)
}

func (um *Manager) Get(userId int64) *User {
	v, ok := um.store.Load(userId)
	if ok {
		return v.(*User)
	}
	return nil
}
