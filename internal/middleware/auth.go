package middleware

import (
	"HackDayBackend/global"
	"HackDayBackend/internal/model"
	"HackDayBackend/internal/verify"
	"HackDayBackend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc { //中间件
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		tokenString = tokenString[7:] //丢弃开头部分

		token, claims, err := verify.ParseToken(tokenString)
		if err != nil || !token.Valid { //返回出错或者token无效
			utils.Failed(c, 400, "Insufficient authority", nil)
			c.Abort() //抛弃
			return
		}

		Id := claims.Id

		var user model.User
		global.Db.Model(&model.User{}).Where("id = ?", Id).First(&user)

		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "用户不存在!"})
			utils.Failed(c, 400, "user does not exist", nil)
			c.Abort() //抛弃
			return
		}

		c.Set("user", Id) //写入上下文
		c.Next()

	}
}
