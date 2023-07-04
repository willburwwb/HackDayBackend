package global

import (
	"HackDayBackend/configs"
	"HackDayBackend/utils"
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	Db  *gorm.DB
	Rdb *redis.Client
)

func Set() error {
	var err error

	Rdb, err = setupRedis()
	if err != nil {
		utils.ErrorF("set redis error: %s", err)
		return err
	}

	time.Sleep(1 * time.Second)
	Db, err = SetMysql()
	if err != nil {
		utils.ErrorF("set pgsql error: %s", err)
		return err
	}

	return nil
}

func SetMysql() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?readTimeout=1500ms&writeTimeout=1500ms&charset=utf8&loc=Local&&parseTime=true",
		configs.Mysql_config.User, configs.Mysql_config.Password, configs.Mysql_config.Ip, configs.Mysql_config.Port, configs.Mysql_config.Name)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
		SkipDefaultTransaction: true, // 跳过默认事务, 加快数据库操作速度, 注意写入数据的时候需要手动开启事务
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := database.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(configs.Mysql_config.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(configs.Mysql_config.MaxLifeSeconds) * time.Hour)
	sqlDB.SetMaxOpenConns(configs.Mysql_config.MaxOpenConns)

	return database, nil
}

func setupRedis() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     configs.Redis_config.Addr,
		Password: configs.Redis_config.Password,
		DB:       configs.Redis_config.DB,
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return rdb, nil
}
