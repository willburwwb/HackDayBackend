package model

import (
	"HackDayBackend/global"
)

func SetModel() {
	db := global.Db

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Node{})
}
