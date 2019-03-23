package tyrds

import (
	"time"

	log "github.com/alecthomas/log4go"
	redigo "github.com/garyburd/redigo/redis"
)

// var pool *redigo.Pool

type RdsServer struct {
	host string
	conn redigo.Conn
	pool *redigo.Pool
}

func init() {

}

func newPool(server string, password string, db int) *redigo.Pool {
	return &redigo.Pool{
		MaxIdle:     3,
		IdleTimeout: time.Duration(240) * time.Second,
		Dial: func() (redigo.Conn, error) {
			c, err := redigo.Dial("tcp", server, redigo.DialDatabase(db), redigo.DialPassword(password))
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redigo.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func ConnectRedis(host string) *RdsServer {
	if len(host) == 0 {
		log.Error("host is empty")
		return nil
	}

	// conn, err := redigo.Dial("tcp", host)
	// if err != nil {
	// 	fmt.Printf("Connect to redis error", err)
	// 	return nil
	// }

	return &RdsServer{
		host: host,
		// conn: conn,
		pool: newPool(host, "123456", 0),
	}
}

func (rs *RdsServer) GetConn() redigo.Conn {
	rs.conn = rs.pool.Get()
	return rs.conn
}

func (rs *RdsServer) SET(key, value string) int {
	if len(key) == 0 {
		return -1
	}

	if len(value) == 0 {
		return -1
	}

	reply, err := rs.conn.Do("SET", key, value)
	if err != nil {
		log.Error(err)
		return -2
	}

	switch reply.(type) {
	case string:
		log.Info(reply.(string))
	case int:
		log.Info(reply.(int))
	}

	if v2, ok := reply.(string); ok {
		log.Info(v2)
	}

	return 0
}

func (rs *RdsServer) SETEX(key, value string, ex int) int {
	if len(key) == 0 {
		return -1
	}

	if len(value) == 0 {
		return -1
	}
	reply, err := rs.conn.Do("SETEX", key, ex, value)
	if err != nil {
		log.Error(err)
		return -2
	}

	switch reply.(type) {
	case string:
		log.Info(reply.(string))
	case int:
		log.Info(reply.(int))
	}

	return 0
}

func (rs *RdsServer) GET(key string) string {
	if len(key) == 0 {
		return ""
	}

	reply, err := redigo.String(rs.conn.Do("GET", key))
	if err != nil {
		log.Error(err)
	} else {
		log.Info("Get %s: %v \n", key, reply)
	}

	return reply
}

func (rs *RdsServer) EXPIRE(key string, t int64) bool {
	if len(key) == 0 {
		return false
	}

	reply, err := rs.conn.Do("EXPIRE", "EXPIRE:"+key, t)
	if err != nil {
		log.Error(err)
		return false
	}

	if reply == int64(1) {
		log.Info("success")
	}
	return true
}

func (rs *RdsServer) INCR(key string) bool {
	if len(key) == 0 {
		return false
	}

	reply, err := rs.conn.Do("INCR", key)
	if err != nil {
		log.Error(err)
		return false
	}

	if reply == int64(1) {
		log.Info("success")
	}

	return true
}

func (rs *RdsServer) DECR(key string) bool {
	if len(key) == 0 {
		return false
	}

	reply, err := rs.conn.Do("DECR", key)
	if err != nil {
		log.Error(err)
		return false
	}

	if reply == int64(1) {
		log.Info("success")
	}

	return true
}

func (rs *RdsServer) EXISTS(key string) bool {
	if len(key) == 0 {
		return false
	}

	reply, err := rs.conn.Do("EXISTS", key)
	if err != nil {
		log.Error(err)
		return false
	}

	if reply.(int64) == 1 {
		log.Info("success")
	}

	return true
}
