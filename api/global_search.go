package api

import (
	"github.com/gin-gonic/gin"
	"my-docs/model"
	"my-docs/remodels"
)

func GlobalSearch(c *gin.Context) {
	query := c.Query("text")
	posts, err := model.GetAnyPosts(query)
	if err != nil {
		c.JSON(400, err)
	}
	var resp []remodels.SearchRe
	for _, p := range posts {
		var re remodels.SearchRe
		re.ID = p.ID
		re.Title = p.Title
		resp = append(resp, re)
	}
	c.JSON(200, resp)
}
