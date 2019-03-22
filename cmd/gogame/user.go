package main

type User struct {
	Id       int64
	Alisid   int64
	Name     string
	Nickname string
}

type UserManager struct {
	count int64
	users map[int64]User
}

type IUserManager interface {
	Add(u User)
	Remove(u User)

	FindWithId(id int64) User
	FindWithAlisid(id int64) User
	FindWithName(name string) User
}

func (m *UserManager) Add(u User) {

}

func (m *UserManager) Remove(u User) {

}

func (m *UserManager) FindWithId(id int64) User {
	return m.users[id]
}

func (m *UserManager) FindWithAlisid(id int64) User {
	return m.users[id]
}

func (m *UserManager) FindWithName(name string) User {
	return User{}
}
