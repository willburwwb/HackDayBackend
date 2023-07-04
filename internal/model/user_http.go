package model

type PhoneCode struct {
	Phone string `json:"phone" form:"phone" binding:"required"`
}

type UserLogin struct {
	Phone string `json:"phone" form:"phone" binding:"required"`
	Code  string `json:"code" form:"phone" binding:"required"`
}

type PwdLogin struct {
	Phone    string `json:"phone" form:"phone" binding:"required"`
	Password string `json:"password" form:"phone" binding:"required"`
}

type SetPassword struct {
	Password string `json:"password" form:"password" binding:"required"`
}
