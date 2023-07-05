package node

import (
	"HackDayBackend/internal/model"
	"HackDayBackend/internal/prompt"
	"HackDayBackend/utils"
	"github.com/gin-gonic/gin"
	"log"
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
	node, err := model.SelectNodeByID(nodeReq.Id)
	if err != nil {
		log.Println("db select error", err)
		utils.Failed(c, 400, "db select error", nil)
	}
	newinfo, err := prompt.GetInfoByLableAndInfoPrompt(node.Label, node.Info)
	if err != nil {
		log.Println("prompt error", err)
		return
	}
	log.Println("newinfo ", newinfo)

	newlabel, err := prompt.GetLableByInfoPrompt(newinfo)
	if err != nil {
		log.Println("prompt error", err)
		return
	}
	log.Println("newinfo ", newlabel)

	err = model.CreateNewNode(&model.Node{
		Label: newlabel,
		Info:  newinfo,
	})
}
