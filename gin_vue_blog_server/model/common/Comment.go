package common

type Comment struct {
	UserId    int    `json:"user_id"`
	ArticleId int    `json:"article_id"`
	ParentId  int    `json:"parent_id"`
	Content   string `json:"content"`
}
