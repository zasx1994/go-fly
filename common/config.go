package common

import (
	"encoding/json"
	"github.com/taoshihan1991/imaptool/tools"
	"io/ioutil"
)

type Mysql struct {
	Server   string
	Port     string
	Database string
	Username string
	Password string
}

type Redis struct {
	Server   string
	Port     string
	Password string
}

func GetMysqlConf() *Mysql {
	var mysql = &Mysql{}
	isExist, _ := tools.IsFileExist(MysqlConf)
	if !isExist {
		return mysql
	}
	info, err := ioutil.ReadFile(MysqlConf)
	if err != nil {
		return mysql
	}
	err = json.Unmarshal(info, mysql)
	return mysql
}

func GetRedisConf() *Redis {
	var redis = &Redis{}
	isExist, _ := tools.IsFileExist(RedisConf)
	if !isExist {
		return redis
	}
	info, err := ioutil.ReadFile(RedisConf)
	if err != nil {
		return redis
	}
	err = json.Unmarshal(info, redis)
	return redis
}
