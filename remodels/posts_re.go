package remodels

type PostsRe struct {
	Key      int        `json:"key"`
	Label    string     `json:"label"`
	Children []Children `json:"children"`
}

type Children struct {
	Key   string `json:"key"`
	Label string `json:"label"`
}

type PostDetailRe struct {
	Key       string `json:"key"`
	Title     string `json:"title"`
	UpdatedAt string `json:"updated_at"`
	Content   string `json:"content"`
}
