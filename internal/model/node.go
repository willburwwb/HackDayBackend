package model

type Node struct {
	ID    string `gorm:"column:id;type:uuid;default:gen_random_uuid();primarykey" json:"id"`
	Label string `gorm:"columns:Label;not null" json:"label"`
	Info  string `gorm:"columns:Info;not null" json:"info"`
}

func (e Node) TableName() string {
	return "node"
}
