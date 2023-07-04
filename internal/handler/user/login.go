package user

import (
	"HackDayBackend/global"
	"HackDayBackend/internal/model"
	"HackDayBackend/internal/verify"
	"HackDayBackend/utils"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db = global.Db
var rdb = global.Rdb

func GetCode(c *gin.Context) {
	var pcode model.PhoneCode
	if err := c.ShouldBind(&pcode); err != nil {
		utils.ErrorF("GetCode bind phone err: %s", err)
		utils.Failed(c, 400, "bind phone err", nil)
		return
	}

	// get code
	code := verify.CreateVerifyCode()

	err := rdb.Set(context.Background(), pcode.Phone, code, 3*time.Minute).Err()
	if err != nil {
		utils.ErrorF("store phone code error: %s", err)
		utils.Failed(c, 400, "try again", nil)
		return
	}

	// send code
	utils.SendSms(pcode.Phone, code)

	utils.DebugF("phone: %s, code: %s", pcode.Phone, code)
	utils.Success(c, 200, "验证码已发送", nil)
}

// register and login or just login
func LoginWithCode(c *gin.Context) {
	var login model.UserLogin
	if err := c.ShouldBind(&login); err != nil {
		utils.ErrorF("LoginWithCode bind err: %s", err)
		utils.Failed(c, 400, "bind params err", nil)
		return
	}

	code, err := rdb.Get(context.Background(), login.Phone).Result()
	if err != nil {
		//utils.DebugF("phone code has expired")
		utils.Failed(c, 400, "email code has expired", nil)
		return
	}

	if code != login.Code {
		utils.DebugF("code: %s, original code: %s", login.Code, code)
		utils.Failed(c, 400, "verify code error!", nil)
		return
	}

	var user model.User
	if err := db.Model(&model.User{}).Where("phone = ?", login.Phone).First(&user).Error; err == gorm.ErrRecordNotFound {
		// register
		user = model.User{
			Phone:        login.Phone,
			PasswordHash: "",
		}

		db.Model(&model.User{}).Create(&user)

		// release token
		token, err := verify.ReleaseToken(user.ID, false)
		if err != nil {
			utils.ErrorF("get token error: %s", err)
			utils.Failed(c, 500, "get token error", nil)
			return
		}

		c.Header("Authorization", token)
		utils.Success(c, 200, "login success", token)
	} else {
		// already register
		token, err := verify.ReleaseToken(user.ID, false)
		if err != nil {
			utils.ErrorF("get token error: %s", err)
			utils.Failed(c, 500, "get token error", nil)
			return
		}

		c.Header("Authorization", token)
		utils.Success(c, 200, "login success", token)
	}
}

func LoginWithPwd(c *gin.Context) {
	var pLogin model.PwdLogin
	if err := c.ShouldBind(&pLogin); err != nil {
		utils.ErrorF("LoginWithPwd bind err: %s", err)
		utils.Failed(c, 400, "bind params err", nil)
		return
	}

	var user model.User
	if err := db.Model(&model.User{}).Where("phone = ?", pLogin.Phone).First(&user).Error; err == gorm.ErrRecordNotFound {
		utils.DebugF("user phone not found: %s", pLogin.Phone)
		utils.Failed(c, 400, "This account has not been registered", nil)
		return
	}

	if user.PasswordHash == "" {
		utils.Failed(c, 400, "you have not set password", nil)
		return
	}

	if err := utils.VerifyPassword(user.PasswordHash, pLogin.Password); err != nil {
		utils.DebugF("phone: %s, original pwd: %s, pwd: %s", user.Phone, user.PasswordHash, pLogin.Password)
		utils.Failed(c, 400, "wrong password", nil)
		return
	}

	token, err := verify.ReleaseToken(user.ID, false)
	if err != nil {
		utils.ErrorF("get token error: %s", err)
		utils.Failed(c, 500, "get token error", nil)
		return
	}

	c.Header("Authorization", token)
	utils.Success(c, 200, "login success", token)
}
