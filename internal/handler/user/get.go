package user

import (
	"HackDayBackend/internal/model"
	"HackDayBackend/utils"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	var user model.User
	// 获取userid
	uid := c.GetUint("user")
	if err := db.Model(&model.User{}).Where("id = ?", uid).First(&user).Error; err != nil {
		utils.ErrorF("getuserinfo db error: %s", err)
		utils.Failed(c, 500, "server internal error", nil)
		return
	}

	utils.Success(c, 200, "get user info successfully", user)
}
