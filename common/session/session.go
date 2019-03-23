package session

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/url"
	"strconv"
	"sync"
	"time"
	// "github.com/gin-contrib/sessions"
	// "github.com/gin-contrib/sessions/cookie"
	// "github.com/gin-contrib/sessions/memstore"
	// "github.com/gin-contrib/sessions/redis"
)

// TySession
type TySession struct {
	SessionID        string
	LastTimeAccessed time.Time
	Values           map[interface{}]interface{}
}

// TySession会话管理
type TyNewSessionManager struct {
	CookieName  string // 客户端cookie名称
	MaxLifeTime int64  // 垃圾回收时间
	m           sync.Mutex
	Sessions    map[string]*TySession // 保存session指针
}

// var (
// 	ins *TyNewSessionManager
// 	mu  sync.Mutex
// )

// func GetSessionManager() *TyNewSessionManager {
// 	if ins == nil {
// 		mu.Lock()
// 		defer mu.Unlock()

// 		if ins == nil {
// 			ins = &TyNewSessionManager{}
// 		}
// 	}

// 	return ins
// }

type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxLifeTime int64)
}

type Session interface {
	Set(key, value interface{}) error // set session value
	Get(key interface{}) interface{}  // get session value
	Delete(key interface{}) error     // delete session value
	SessionID() string                // back current sessionID
}

func NewSessionManager(cookieName string, maxLifeTime int64) *TyNewSessionManager {
	mgr := &TyNewSessionManager{
		CookieName:  cookieName,
		MaxLifeTime: maxLifeTime,
		Sessions:    make(map[string]*TySession),
	}

	mgr.sessionGC()

	return mgr
}

func (this *TyNewSessionManager) newSessionId() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func (this *TyNewSessionManager) NewSessionID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		nano := time.Now().UnixNano()
		return strconv.FormatInt(nano, 10)
	}

	return base64.URLEncoding.EncodeToString(b)
}

func (this *TyNewSessionManager) StartSession() string {
	newSessionID := url.QueryEscape(this.NewSessionID())

	var session *TySession = &TySession{
		SessionID:        newSessionID,
		LastTimeAccessed: time.Now(),
		Values:           make(map[interface{}]interface{}),
	}
	this.Sessions[newSessionID] = session

	return newSessionID
}

func (this *TyNewSessionManager) EndSession() {

}

func (this *TyNewSessionManager) SetSessionValue(sessionID string, key interface{}, value interface{}) {
	this.m.Lock()
	defer this.m.Unlock()

	if session, ok := this.Sessions[sessionID]; ok {
		session.Values[key] = value
	}
}

func (this *TyNewSessionManager) GetSessionValue(sessionID string, key interface{}) (interface{}, bool) {
	this.m.Lock()
	defer this.m.Unlock()

	if session, ok := this.Sessions[sessionID]; ok {
		if val, ok := session.Values[key]; ok {
			return val, ok
		}
	}

	return nil, false
}

func (this *TyNewSessionManager) GetSessionIDList() []string {
	this.m.Lock()
	defer this.m.Unlock()

	sessionIDArray := make([]string, 0)

	for k, _ := range this.Sessions {
		sessionIDArray = append(sessionIDArray, k)
	}

	return sessionIDArray[0:len(sessionIDArray)]
}

func (this *TyNewSessionManager) GetLastAccessTime(sessionID string) time.Time {
	this.m.Lock()
	defer this.m.Unlock()

	if session, ok := this.Sessions[sessionID]; ok {
		return session.LastTimeAccessed
	}

	return time.Now()
}

func (this *TyNewSessionManager) sessionGC() {
	this.m.Lock()
	defer this.m.Unlock()

	for sessionid, session := range this.Sessions {
		if session.LastTimeAccessed.Unix()+this.MaxLifeTime < time.Now().Unix() {
			delete(this.Sessions, sessionid)
		}
	}

	time.AfterFunc(time.Duration(this.MaxLifeTime)*time.Second, func() {
		this.sessionGC()
	})
}
