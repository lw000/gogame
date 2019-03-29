package ggrdsex

import (
	"errors"
	log "github.com/alecthomas/log4go"
	"strings"
	"sync"
	"time"
	// "gopkg.in/redis.v4"
	"github.com/go-redis/redis"
)

type RdsServer struct {
	done   chan struct{}
	Client *redis.Client
}

var (
	rdsSrv  *RdsServer
	rdsOnce sync.Once
)

func SharedRdsServerInsance() *RdsServer {
	rdsOnce.Do(func() {
		rdsSrv = NewResServer()
	})
	return rdsSrv
}

func NewResServer() *RdsServer {
	return &RdsServer{
		done: make(chan struct{}, 1),
	}
}

func (rs *RdsServer) connect(host string, psd string, db int64, poolSize int64, minIdleConns int64) (*RdsServer, error) {
	if rs == nil {
		return nil, errors.New("object is nil")
	}

	if len(host) == 0 {
		return nil, errors.New("redis host is empty")
	}

	rs.Client = redis.NewClient(&redis.Options{
		Addr:         host,
		Password:     psd,
		DB:           int(db),
		PoolSize:     int(poolSize),
		MinIdleConns: int(minIdleConns),
		MaxConnAge:   time.Hour * time.Duration(2),
	})

	pong := rs.Client.Ping()
	if strings.ToUpper(pong.Val()) != "PONG" {
		return nil, errors.New(pong.Err().Error())
	}

	//TODO:心跳 PING <-> PONG
	//go rs.runPingPong()

	return rs, nil
}

func (rs *RdsServer) ConnectRedis(c *RdsConfigStruct) (*RdsServer, error) {
	if rs == nil {
		return nil, errors.New("object is nil")
	}

	if c == nil {
		return nil, errors.New("config is nil")
	}

	return rs.connect(c.Host, c.Psd, c.Db, c.PoolSize, c.MinIdleConns)
}

func (rs *RdsServer) DisconnectRedis() error {
	if rs == nil {
		return errors.New("object is nil")
	}

	rs.done <- struct{}{}

	er := rs.Client.Close()
	if er != nil {
		return er
	}

	return nil
}

func (rs *RdsServer) Pipe() *redis.Pipeline {
	pipe := rs.Client.Pipeline().(*redis.Pipeline)
	return pipe
}

func (rs *RdsServer) ListenKeyExpired() {
	parameter := "notify-keyspace-events"
	rdscfg, er := rs.Client.ConfigGet(parameter).Result()
	if er != nil {
		log.Error(er)
		return
	}

	v := rdscfg[1]
	if v.(string) == "" {
		s, er := rs.Client.ConfigSet(parameter, "Ex").Result()
		if er != nil {
			log.Error(er)
			return
		}
		log.Info(s)
	}

	log.Info(rdscfg)

	Pubsub := rs.Client.Subscribe("__keyevent@0__:expired")
	for {
		msg, er := Pubsub.ReceiveMessage()
		if er != nil {
			log.Error(er)
		}
		log.Info(msg)
	}
}

func (rs *RdsServer) runPingPong() {
	defer func() {
		log.Error("client redis exit")
	}()

	ticker := time.NewTicker(time.Minute * time.Duration(1))
	for {
		select {
		case <-ticker.C:
			s, er := rs.Client.Ping().Result()
			if er != nil {
				log.Error(er)
				return
			}
			if strings.ToUpper(s) != "PONG" {
				log.Error("%s", s)
				return
			}

			log.Info(s)

		case <-rs.done:
			ticker.Stop()
			close(rs.done)
			return
		default:
		}
	}
}
