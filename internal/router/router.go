package router

import (
	"HackDayBackend/internal/handler/node"
	"HackDayBackend/internal/handler/user"
	"HackDayBackend/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	ping := r.Group("/ping")
	ping.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "this is hackday",
		})
	})

	r.Use(middleware.CorsMiddleware)
	//r.Use(middleware.AuthMiddleware)

	userGroup := r.Group("/user")
	{
		userGroup.POST("/getCode", user.GetCode)
		userGroup.POST("/loginWithCode", user.LoginWithCode)
		userGroup.POST("/loginWithPwd", user.LoginWithPwd)

		userGroup.Use(middleware.AuthJWT())

		userGroup.POST("/updatePwd", user.UpdatePassword)
		userGroup.GET("/userInfo", user.GetUserInfo)
	}

	nodeGroup := r.Group("/node")
	{
		nodeGroup.GET("/")
		nodeGroup.POST("/new", node.CreateNodeById)
	}
	return r
}
