package ggrdsex

import (
	"errors"
	log "github.com/alecthomas/log4go"
	"strings"
	"time"
	tyrdsexconfig "tuyue/tuyue_common/db/rdsex/config"
	// "gopkg.in/redis.v4"
	"github.com/go-redis/redis"
)

type RedisStore interface {
	OnRedisEncode() map[string]interface{}
	OnRedisDecode(data map[string]string) error
}

type RdsServer struct {
	client *redis.Client
}

func (r *RdsServer) Client() *redis.Client {
	return r.client
}

func (r *RdsServer) connect(host string, psd string, db int64, poolSize int64, minIdleConns int64) error {
	if r == nil {
		return errors.New("object is nil")
	}

	if len(host) == 0 {
		return errors.New("redis host is empty")
	}

	r.client = redis.NewClient(&redis.Options{
		Addr:         host,
		Password:     psd,
		DB:           int(db),
		PoolSize:     int(poolSize),
		MinIdleConns: int(minIdleConns),
		MaxConnAge:   time.Hour * time.Duration(2),
	})

	pong := r.client.Ping()
	if strings.ToUpper(pong.Val()) != "PONG" {
		return errors.New(pong.Err().Error())
	}

	return nil
}

func (r *RdsServer) OpenWithJsonConfig(cfg *tyrdsexconfig.JsonConfigStruct) error {
	if r == nil {
		return errors.New("object is nil")
	}

	if cfg == nil {
		return errors.New("config is nil")
	}

	return r.connect(cfg.Host, cfg.Psd, cfg.Db, cfg.PoolSize, cfg.MinIdleConns)
}

func (r *RdsServer) OpenWithYamlConfig(cfg *tyrdsexconfig.YamlConfigStruct) error {
	if r == nil {
		return errors.New("object is nil")
	}

	if cfg == nil {
		return errors.New("config is nil")
	}

	return r.connect(cfg.Host, cfg.Psd, cfg.Db, cfg.PoolSize, cfg.MinIdleConns)
}

func (r *RdsServer) Close() error {
	if r == nil {
		return errors.New("object is nil")
	}

	er := r.client.Close()
	if er != nil {
		return er
	}

	return nil
}

func (r *RdsServer) Pipe() *redis.Pipeline {
	pipe := r.client.Pipeline().(*redis.Pipeline)
	return pipe
}

//func (r *RdsServer) Scan(cursor uint64, match string, count int64) ([]string, uint64, error) {
//	return r.client.Scan(cursor, match, count).Result()
//}

func (r *RdsServer) ScanKeys(match string, count int64, f func(keys []string)) error {
	var (
		er     error
		cursor uint64 = 0
		keys   []string
	)
	for {
		keys, cursor, er = r.client.Scan(cursor, match, count).Result()
		if er != nil {
			log.Error(er)
			break
		}

		if len(keys) > 0 {
			f(keys)
		}

		if cursor == 0 {
			break
		}
	}
	return er
}

func (r *RdsServer) Del(keys ...string) (int64, error) {
	return r.client.Del(keys...).Result()
}

func (r *RdsServer) Keys(key string) ([]string, error) {
	return r.client.Keys(key).Result()
}

func (r *RdsServer) Exists(key ...string) bool {
	exist, er := r.client.Exists(key...).Result()
	if er != nil {
		log.Error(er)
		return false
	}
	return exist == 1
}

func (r *RdsServer) SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	return r.client.SetNX(key, value, expiration).Result()
}

func (r *RdsServer) Get(key string) (string, error) {
	return r.client.Get(key).Result()
}

func (r *RdsServer) Set(key string, v interface{}, expiration time.Duration) (string, error) {
	return r.client.Set(key, v, expiration).Result()
}

func (r *RdsServer) GetSet(key string, v interface{}) (string, error) {
	return r.client.GetSet(key, v).Result()
}

func (r *RdsServer) HGetAll(key string) (map[string]string, error) {
	return r.client.HGetAll(key).Result()
}

func (r *RdsServer) HMSet(key string, fields map[string]interface{}) (string, error) {
	return r.client.HMSet(key, fields).Result()
}

func (r *RdsServer) HMGet(key string, fields ...string) ([]interface{}, error) {
	return r.client.HMGet(key, fields...).Result()
}

//func (r *RdsServer) HScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
//	return r.client.HScan(key, cursor, match, count).Result()
//}

func (r *RdsServer) HScanValues(key string, match string, count int64, f func(values []string)) error {
	var (
		er     error
		cursor uint64 = 0
		values []string
	)
	for {
		values, cursor, er = r.client.HScan(key, cursor, match, count).Result()
		if er != nil {
			log.Error(er)
			break
		}

		if len(values) > 0 {
			f(values)
		}

		if cursor == 0 {
			break
		}
	}

	return er
}

func (r *RdsServer) SAdd(key string, members ...interface{}) (int64, error) {
	return r.client.SAdd(key, members...).Result()
}

func (r *RdsServer) SMembers(key string) ([]string, error) {
	return r.client.SMembers(key).Result()
}

//func (r *RdsServer) SScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
//	return r.client.SScan(key, cursor, match, count).Result()
//}

func (r *RdsServer) SScanValues(key string, match string, count int64, f func(values []string)) error {
	var (
		er     error
		cursor uint64 = 0
		values []string
	)
	for {
		values, cursor, er = r.client.SScan(key, cursor, match, count).Result()
		if er != nil {
			log.Error(er)
			break
		}

		if len(values) > 0 {
			f(values)
		}

		if cursor == 0 {
			break
		}
	}
	return er
}

func (r *RdsServer) ZScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return r.client.ZScan(key, cursor, match, count).Result()
}

func (r *RdsServer) ZScanValues(key string, match string, count int64, f func(values []string)) error {
	var (
		er     error
		cursor uint64 = 0
		values []string
	)
	for {
		values, cursor, er = r.client.ZScan(key, cursor, match, count).Result()
		if er != nil {
			log.Error(er)
			break
		}

		if len(values) > 0 {
			f(values)
		}

		if cursor == 0 {
			break
		}
	}
	return er
}

func (r *RdsServer) IncrBy(key string, value int64) (int64, error) {
	return r.client.IncrBy(key, value).Result()
}

func (r *RdsServer) ListenKeyExpired() {
	evnetKey := "notify-keyspace-events"
	rdscfg, er := r.client.ConfigGet(evnetKey).Result()
	if er != nil {
		log.Error(er)
		return
	}

	v := rdscfg[1]
	if v.(string) == "" {
		var s string
		s, er = r.client.ConfigSet(evnetKey, "Ex").Result()
		if er != nil {
			log.Error(er)
			return
		}
		log.Info(s)
	}

	log.Info(rdscfg)

	Pubsub := r.client.Subscribe("__keyevent@0__:expired")
	for {
		var msg *redis.Message
		msg, er = Pubsub.ReceiveMessage()
		if er != nil {
			log.Error(er)
		}
		log.Info(msg)
	}
}
