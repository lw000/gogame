package tydb

import (
	"log"
	"testing"
)

//表结构
type TUser struct {
	Name    string `json:"Name" form:"name"`
	Age     int    `json:"Age" form:"age"`
	Sex     int    `json:"Sex" form:"sex"`
	Address string `json:"Address" form:"address"`
}

var dbcfg *Config
var dbYamlcfg *YamlMysqlCfg

func SqlQuery(s *Mysql) {
	rows, err := s.Db().Query("SELECT * FROM user")
	defer rows.Close()

	if err != nil {
		log.Panic(err)
	}

	var us []TUser
	for rows.Next() {
		var u TUser
		if err = rows.Scan(&u.Name, &u.Age, &u.Sex, &u.Address); err == nil {
			us = append(us, u)
		}
	}
	if err = rows.Err(); err != nil {
		log.Panic(err)
	}

	for _, v := range us {
		log.Println(v)
	}
}

func SqlQueryRow(s *Mysql) {
	row := s.Db().QueryRow("SELECT * FROM user WHERE name='levi';")
	var u TUser
	err := row.Scan(&u.Name, &u.Age, &u.Sex, &u.Address)
	if err == nil {
		log.Println(u)
	}
}

func SqlQueryWhere(s *Mysql) {
	stms, err := s.Db().Prepare("SELECT * FROM user WHERE name=?;")
	if err != nil {
		log.Panic(err)
	}

	defer stms.Close()

	rows, err := stms.Query("levi")
	if err != nil {
		log.Panic(err)
	}

	var us []TUser
	for rows.Next() {
		var u TUser
		if err = rows.Scan(&u.Name, &u.Age, &u.Sex, &u.Address); err == nil {
			us = append(us, u)
		}
	}
	if err = rows.Err(); err != nil {
		log.Panic(err)
	}

	for _, v := range us {
		log.Println(v)
	}
}

func SqlDelete(s *Mysql) {
	rs, err := s.Db().Exec("DELETE FROM user WHERE name=?;", "levi")
	if err != nil {
		log.Panic(err)
	}
	id, err := rs.RowsAffected()
	if err != nil {
		log.Panic(err)
	}
	log.Printf("sqlDelete:%d\n", id)
}

func SqlUpdate(s *Mysql) {
	rs, err := s.Db().Exec("UPDATE user SET name=? WHERE name=?;", "levi1", "levi")
	if err != nil {
		log.Panic(err)
	}

	id, err := rs.RowsAffected()
	if err != nil {
		log.Panic(err)
	}

	log.Printf("sqlUpdate:%d\n", id)
}

func SqlInsert(s *Mysql) {
	rs, err := s.Db().Exec("INSERT INTO user(name, age, sex ,address) VALUES(?,?,?,?);", "hjt", 30, 1, "beijingshi")
	if err != nil {
		log.Panic(err)
	}

	id, err := rs.LastInsertId()
	if err != nil {
		log.Panic(err)
	}

	log.Println("sqlInsert: ", id)
}

func TestRunTest(t *testing.T) {
	var err error
	dbcfg, err = LoadConfigWithData([]byte(`{
	"status": 0,
	"comment": "dev=0开发环境 test=1测试环境 prod=2正式环境",
	"dev": {
		"username": "root",
		"password": "root",
		"network": "tcp",
		"host": "192.168.1.101:3306",
		"database": "mservice",
		"MaxOpenConns": 20,
		"MaxOdleConns": 5
	},
	"test": {
		"username": "root",
		"password": "root",
		"network": "tcp",
		"host": "192.168.1.101:3306",
		"database": "mservice",
		"MaxOpenConns": 20,
		"MaxOdleConns": 5
	},
	"prod": {
		"username": "root",
		"password": "Aabbcc123!@#",
		"network": "tcp",
		"host": "47.96.230.81:3306",
		"database": "mservice",
		"MaxOpenConns": 20,
		"MaxOdleConns": 5
	}
}`))
	if err != nil {
		log.Panic(err)
	}
	log.Println(dbcfg)

	dbYamlcfg, err = LoadConfigWithYamlData([]byte(`
#数据库配置
username: root
password: lwstar
host: 127.0.0.1:3306
port: 3306
database: lw
maxOpenConns: 20
maxOdleConns: 5`))
	if err != nil {
		log.Panic(err)
	}
	log.Println(dbYamlcfg)

	sql := SharedMysqlInstance()
	//_, err = sql.OpenWithConfig(&dbcfg.Cfg)
	_, err = sql.OpenWithYamlConfig(dbYamlcfg)
	if err != nil {
		log.Panic(err)
	}
	SqlQuery(sql)
	SqlQueryRow(sql)
	SqlQueryWhere(sql)
	SqlUpdate(sql)
	//	//SqlInsert(sql)
	//	// sqlDelete(sql)
}
