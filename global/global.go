package global

import (
	"HackDayBackend/configs"
	"HackDayBackend/utils"
	"context"
	"fmt"
	"gorm.io/driver/postgres"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
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
	//Db, err = SetMysql()
	//if err != nil {
	//	utils.ErrorF("set pgsql error: %s", err)
	//	return err
	//}
	Db, err = setupPgsql()
	if err != nil {
		utils.ErrorF("set pgsql error: %s", err)
		return err
	}
	return nil
}

//	func SetMysql() (*gorm.DB, error) {
//		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
//			configs.Mysql_config.User, configs.Mysql_config.Password, configs.Mysql_config.Ip, configs.Mysql_config.Name)
//		fmt.Println("dns", dsn)
//		database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//		if err != nil {
//			return nil, err
//		}
//
//		sqlDB, err := database.DB()
//		if err != nil {
//			return nil, err
//		}
//
//		sqlDB.SetMaxIdleConns(configs.Mysql_config.MaxIdleConns)
//		sqlDB.SetConnMaxLifetime(time.Duration(configs.Mysql_config.MaxLifeSeconds) * time.Hour)
//		sqlDB.SetMaxOpenConns(configs.Mysql_config.MaxOpenConns)
//
//		return database, nil
//	}
func setupPgsql() (*gorm.DB, error) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s sslmode=disable ",
		configs.Pgsql_config.Host, configs.Pgsql_config.User, configs.Pgsql_config.Dbname, configs.Pgsql_config.Port, configs.Pgsql_config.Password)
	//if configs.Config.Pgsql.Password != "" {
	//	dsn = dsn + fmt.Sprintf("password=%s", configs.Config.Pgsql.Password)
	//}
	log.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("connect to db error, %v", err))
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("get db error, %v", err))
	}
	sqlDB.SetMaxIdleConns(configs.Pgsql_config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(configs.Pgsql_config.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(configs.Pgsql_config.MaxLifeSeconds) * time.Second)

	return db, nil
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
