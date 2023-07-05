package model

import (
	"HackDayBackend/global"
	"gorm.io/gorm"
)

type Node struct {
	ID    string `gorm:"column:id;type:uuid;default:gen_random_uuid();primarykey" json:"id"`
	Label string `gorm:"columns:Label;not null" json:"label"`
	Info  string `gorm:"columns:Info;not null" json:"info"`
}

func (e Node) TableName() string {
	return "node"
}
func SelectNodeByID(uid string) (*Node, error) {
	var node Node
	if err := global.Db.Model(&Node{}).Where("uid = ?", uid).First(&node).Error; err == gorm.ErrRecordNotFound {
		return nil, err
	}
	return &node, nil
}

func CreateNewNode(node *Node) error {
	if err := global.Db.Model(&Node{}).Create(node).Error; err != nil {
		return err
	}
	return nil
}
