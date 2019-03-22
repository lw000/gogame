package moyrds

import (
	"errors"
	"time"

	// "fmt"
	"sync"

	log "github.com/alecthomas/log4go"

	// "gopkg.in/redis.v4"
	"github.com/go-redis/redis"
)

type RdsServer struct {
	host   string
	Client *redis.Client
}

var (
	ins  *RdsServer
	once sync.Once
)

func SharedRdsServerInsance() *RdsServer {
	once.Do(func() {
		ins = &RdsServer{}
	})

	return ins
}

func (this *RdsServer) ConnectRedis(host string, psd string) (*RdsServer, error) {
	if len(host) == 0 {
		log.Error("host is empty")
		return nil, errors.New("host is empty")
	}

	this.Client = redis.NewClient(&redis.Options{
		Addr:         host,
		Password:     psd,
		DB:           0,
		PoolSize:     20,
		MinIdleConns: 5,
		MaxConnAge:   time.Hour * time.Duration(2),
		// PoolTimeout: 240 * time.Second,
	})

	pong := this.Client.Ping()
	if pong.Val() != "PONG" {
		return nil, errors.New("connected faild")
	}

	return this, nil
}

func (this *RdsServer) DisconnectRedis() {
	this.Client.Close()
}
