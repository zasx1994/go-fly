package models

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"github.com/taoshihan1991/imaptool/common"
	"log"
	"time"
)

var DB *gorm.DB
var RDB *redis.Client

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

func init() {
	Connect()
	RDBConnect()
}

func RDBConnect() error {
	c := common.GetRedisConf()
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Server + ":" + c.Port,
		Password:     c.Password,
		WriteTimeout: time.Second,
		ReadTimeout:  time.Second,
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return err
	}
	RDB = rdb
	return nil
}
func Connect() error {
	mysql := common.GetMysqlConf()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysql.Username, mysql.Password, mysql.Server, mysql.Port, mysql.Database)
	var err error
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
		panic("数据库连接失败!")
		return err
	}
	DB.SingularTable(true)
	DB.LogMode(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	DB.DB().SetConnMaxLifetime(59 * time.Second)
	InitConfig()
	return nil
}
func Execute(sql string) error {
	return DB.Exec(sql).Error
}
func CloseDB() {
	defer DB.Close()
}
