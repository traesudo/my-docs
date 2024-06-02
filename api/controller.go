package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type APIController struct {
	*gin.Context
}
type Doc interface {
	GetPost()
	GetPosts()
	GetDirections()
}

func (s *APIController) CJSON(code int, data interface{}) {
	if code != 200 {
		resp := data
		fmt.Println("error", data)
		s.JSON(code, resp)
	} else {
		resp := make(map[string]interface{})
		resp["data"] = data
		s.JSON(code, resp)
	}
}
