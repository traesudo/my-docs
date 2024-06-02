package api

import (
	"github.com/gin-gonic/gin"
	"my-docs/model"
	"my-docs/processor"
)

func GetPost(c *gin.Context) {
	id := c.Param("id")
	post, err := model.GetPostById(id)
	if err != nil {
		c.JSON(400, err)
	}
	htmlContent := processor.ConvertToHtml([]byte(post.Content))
	buildProcessor := NewBuildProcessor()
	re := make(map[string]interface{})
	re["data"] = buildProcessor.BuildPostDetail(*post, htmlContent)
	c.JSON(200, re)
}
