package global

import (
	"HackDayBackend/configs"
	"HackDayBackend/utils"
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	Db, err = setupPgsql()
	if err != nil {
		utils.ErrorF("set pgsql error: %s", err)
		return err
	}

	return nil
}

func setupPgsql() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s sslmode=disable TimeZone=Asia/Shanghai ",
		configs.Pgsql_config.Host, configs.Pgsql_config.User, configs.Pgsql_config.Dbname, configs.Pgsql_config.Port, configs.Pgsql_config.Password)
	//if configs.Config.Pgsql.Password != "" {
	//	dsn = dsn + fmt.Sprintf("password=%s", configs.Config.Pgsql.Password)
	//}
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		return
	}
	sqlDB.SetMaxIdleConns(configs.Pgsql_config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(configs.Pgsql_config.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(configs.Pgsql_config.MaxLifeSeconds) * time.Second)
	return
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
