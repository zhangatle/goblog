package article

import (
	"goblog/route"
	"strconv"
)

// Article 文章模型
type Article struct {
	ID    uint64
	Title string
	Body  string
}

func (a *Article) Link() string {
	return route.Name2URL("articles.show", "id", strconv.FormatUint(a.ID, 10))
}
