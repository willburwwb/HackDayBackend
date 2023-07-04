package user

import (
	"HackDayBackend/internal/model"
	"HackDayBackend/utils"

	"github.com/gin-gonic/gin"
)

func UpdatePassword(c *gin.Context) {
	var newPassword model.SetPassword
	if err := c.ShouldBind(&newPassword); err != nil {
		utils.ErrorF("change password parse error: %s", err)
		utils.Failed(c, 400, "get new password, code error", nil)
		return
	}

	// 获取userid
	uid := c.GetUint("user")

	// 更新密码
	hashPwd, err := utils.HashPassword(newPassword.Password)
	if err != nil {
		utils.ErrorF("change password error: %s", err)
		utils.Failed(c, 400, "server internal error", nil)
		return
	}

	if err := db.Model(&model.User{}).Where("id = ?", uid).Update("passwordHash", hashPwd).Error; err != nil {
		utils.ErrorF("UpdatePassword db err: %s", err)
		utils.Failed(c, 500, "server internal err", nil)
		return
	}

	utils.Success(c, 200, "change password sucessfully", nil)
}
