package form

type SearchKeywordForm struct {
	Keyword string `form:"keyword" binding:"required"`
	PageNo  int    `form:"pageNo" binding:"required"`
}
