package node

import (
	"HackDayBackend/utils"
	"github.com/gin-gonic/gin"
)

type CreateNodeRequest struct {
	Id   string `json:"id" form:"id" bind:"required"`
	Info string `json:"info" form:"info" bind:"required"`
}

func CreateNodeById(c *gin.Context) {
	var nodeReq CreateNodeRequest
	if err := c.ShouldBind(&nodeReq); err != nil {
		utils.ErrorF("Create new node request err: %s", err)
		utils.Failed(c, 400, "create node err", nil)
		return
	}
}
