package docs

import (
	"my-docs/api"
	"my-docs/model"
)

func (s *DocsController) GetPosts() {
	query := s.Query("direction")
	post, err := model.GetLastPost(query)
	if err != nil {
		s.JSON(400, err)
	}
	labels, err := model.GetAllLabel()
	if err != nil {
		s.JSON(400, err)
	}
	processor := api.NewBuildProcessor()
	resp := processor.BuildPostsRe(post, labels)
	re := make(map[string]interface{})
	re["data"] = resp
	s.JSON(200, re)
}
