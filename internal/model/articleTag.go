// 2.2.3 4：创建文章标签关联model
package model

type ArticleTag struct {
	*Model
	ArtileId uint8 `json:"article_id"`
	TagId    uint8 `json:"tag_id"`
}

func (a ArticleTag) TableName() string {
	return "blog_articel_tag"
}
