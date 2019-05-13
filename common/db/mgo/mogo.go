package ggmgo

import (
	"fmt"
	log "github.com/alecthomas/log4go"
	"gopkg.in/mgo.v2"
	tymgoconfig "tuyue/tuyue_common/db/mgo/config"

	//"gopkg.in/mgo.v2/bson"
	"time"
)

type Mongo struct {
	session *mgo.Session
}

func (m *Mongo) Open(cfg *tymgoconfig.JsonConfigStruct) error {
	dailInfo := &mgo.DialInfo{
		Addrs:     cfg.Address,
		Direct:    false,
		Timeout:   time.Second * 15,
		Database:  cfg.Db,
		Source:    "",
		Username:  "",
		Password:  "",
		PoolLimit: 1024,
	}
	var er error
	m.session, er = mgo.DialWithInfo(dailInfo)
	if er != nil {
		log.Error(fmt.Sprintf("mgo dail error[%s]\n", er.Error()))
		return er
	}
	// set mode
	m.session.SetMode(mgo.Monotonic, true)

	return nil
}

//func (m *Mongo)Session() *mgo.Session{
//	return m.session.Copy()
//}

func (m *Mongo) Close() {
	m.session.Close()
}

func (m *Mongo) DB(db string) *mgo.Database {
	return m.session.DB(db)
}
