package api

import (
	"github.com/gin-gonic/gin"
	"my-docs/model"
)

func GetLives(c *gin.Context) {
	resp, err := model.GetAllLive()
	if err != nil {
		c.JSON(400, err)
	}
	re := make(map[string]interface{})
	re["data"] = resp
	c.JSON(200, re)
}
