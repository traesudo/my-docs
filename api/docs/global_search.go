package docs

import (
	"my-docs/model"
	"my-docs/remodels"
)

func (s *DocsController) GlobalSearch() {
	query := s.Query("text")
	posts, err := model.GetAnyPosts(query)
	if err != nil {
		s.JSON(400, err)
	}
	var resp []remodels.SearchRe
	for _, p := range posts {
		var re remodels.SearchRe
		re.ID = p.ID
		re.Title = p.Title
		resp = append(resp, re)
	}
	s.JSON(200, resp)
}
