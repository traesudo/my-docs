package api

import (
	"golang.org/x/net/context"
	"my-docs/model"
	"my-docs/remodels"
)

type BuildProcessor struct {
	context.Context
	IP string
}

func NewBuildProcessor() *BuildProcessor {
	return &BuildProcessor{}
}

func (b *BuildProcessor) BuildPostsRe(posts []model.Post, labels []model.Label) []remodels.PostsRe {
	var resp []remodels.PostsRe
	items := make(map[int][]remodels.Children)
	for _, p := range posts {
		var c remodels.Children
		c.Key = p.ID
		c.Label = p.Title
		items[p.Label] = append(items[p.Label], c)
	}
	for k, v := range items {
		var postRe remodels.PostsRe
		postRe.Children = v
		postRe.Key = k
		postRe.Label = GetLabelName(k, labels)
		resp = append(resp, postRe)
	}
	return resp
}

func (b *BuildProcessor) BuildPostDetail(post model.Post, content string) remodels.PostDetailRe {
	var re remodels.PostDetailRe
	re.Content = content
	re.Key = post.ID
	re.Title = post.Title
	return re
}
func GetLabelName(number int, labels []model.Label) string {
	for _, l := range labels {
		if number == l.ID {
			return l.Label
		}
	}
	return ""
}
