package docs

import (
	"my-docs/model"
)

func (s *DocsController) GetDirections() {
	resp, err := model.GetAllDirection()
	if err != nil {
		s.JSON(400, err)
	}
	re := make(map[string]interface{})
	re["data"] = resp
	s.JSON(200, re)
}
