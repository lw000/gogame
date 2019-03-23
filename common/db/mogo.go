package ggdb

import (
	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	// "sync"
)

type TyMogo struct {
	Db *mgo.Session
}

// var (
// 	ins *mgo.Session
// 	mu  sync.Mutex
// )

// func SharedTymogoInstance() *TyMogo {
// 	if ins == nil {
// 		mu.Lock()
// 		defer mu.Unlock()

// 		if ins == nil {
// 			ins = &TyMogo{}
// 		}
// 	}

// 	return ins
// }

func NewMogo() *mgo.Session {
	session, err := mgo.Dial("")
	if err != nil {
		return nil
	}
	defer session.Clone()

	session.SetMode(mgo.Monotonic, true)

	return session
}
