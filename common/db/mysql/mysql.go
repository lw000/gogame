package ggmysql

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
	tydbconfig "tuyue/tuyue_common/db/mysql/config"

	_ "github.com/go-sql-driver/mysql"
	// "github.com/jinzhu/gorm"
)

type Mysql struct {
	db *sql.DB
}

func (m *Mysql) Db() *sql.DB {
	return m.db
}

func (m *Mysql) OpenWithYamlConfig(cfg *tydbconfig.YamlMysqlCfg) error {
	c := make(map[string]interface{})
	c["username"] = cfg.Username
	c["password"] = cfg.Password
	c["host"] = cfg.Host
	c["database"] = cfg.Database
	c["MaxOpenConns"] = cfg.MaxOpenConns
	c["MaxOdleConns"] = cfg.MaxOdleConns
	return m.open(c)
}

func (m *Mysql) OpenWithJsonConfig(cfg *tydbconfig.JsonConfigStruct) error {
	c := make(map[string]interface{})
	c["username"] = cfg.Username
	c["password"] = cfg.Password
	c["host"] = cfg.Host
	c["database"] = cfg.Database
	c["MaxOpenConns"] = cfg.MaxOpenConns
	c["MaxOdleConns"] = cfg.MaxOdleConns
	return m.open(c)
}

func (m *Mysql) open(cfg map[string]interface{}) error {
	var (
		username     string
		password     string
		host         string
		database     string
		maxOpenConns int64
		maxOdleConns int64
	)

	if v, ok := cfg["database"]; ok {
		database = v.(string)
	} else {
		return errors.New("database is empty")
	}

	if v, ok := cfg["username"]; ok {
		username = v.(string)
	} else {
		return errors.New("username is empty")
	}

	if v, ok := cfg["password"]; ok {
		password = v.(string)
	} else {
		return errors.New("password is empty")
	}

	if v, ok := cfg["host"]; ok {
		host = v.(string)
	} else {
		return errors.New("host is empty")
	}

	if v, ok := cfg["MaxOpenConns"]; ok {
		maxOpenConns = v.(int64)
	} else {
		maxOpenConns = 10
	}

	if v, ok := cfg["MaxOdleConns"]; ok {
		maxOdleConns = v.(int64)
	} else {
		maxOdleConns = 0
	}

	var err error
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", username, password, host, database)
	m.db, err = sql.Open("mysql", dns)
	if err != nil {
		return err
	}

	m.db.SetConnMaxLifetime(time.Hour * time.Duration(3))
	m.db.SetMaxOpenConns(int(maxOpenConns))
	m.db.SetMaxIdleConns(int(maxOdleConns))

	if err = m.db.Ping(); err != nil {
		return err
	}

	return nil
}

func (m *Mysql) Close() {
	err := m.db.Close()
	if err != nil {

	}
}
