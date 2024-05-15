package form

type SearchKeywordForm struct {
	Keyword string `form:"keyword" binding:"required"`
}
