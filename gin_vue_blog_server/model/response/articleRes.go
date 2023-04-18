package response

type ArticleSimpleInfo struct {
	ArticleId   uint   `json:"articleId"`
	ArticleName string `json:"articleName"`
	Description string `json:"description"`
}
