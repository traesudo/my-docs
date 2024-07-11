package docs

import (
	"my-docs/api"
	"my-docs/model"
	"my-docs/processor"
)

func (s *DocsController) GetPost() {
	id := s.Param("id")
	post, err := model.GetPostById(id)
	s.Logger()
	if err != nil {
		s.JSON(400, err)
	}
	htmlContent := processor.ConvertToHtml([]byte(post.ContentMarkdown))
	buildProcessor := api.NewBuildProcessor()
	re := make(map[string]interface{})
	re["data"] = buildProcessor.BuildPostDetail(*post, htmlContent)
	s.JSON(200, re)
}
