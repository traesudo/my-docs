package api

import (
	"github.com/gin-gonic/gin"
	"my-docs/model"
)

func GetPosts(c *gin.Context) {
	query := c.Query("direction")
	post, err := model.GetLastPost(query)
	if err != nil {
		c.JSON(400, err)
	}
	labels, err := model.GetAllLabel()
	if err != nil {
		c.JSON(400, err)
	}
	processor := NewBuildProcessor()
	resp := processor.BuildPostsRe(post, labels)
	re := make(map[string]interface{})
	re["data"] = resp
	c.JSON(200, re)
}
