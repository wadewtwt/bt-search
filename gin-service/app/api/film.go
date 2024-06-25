package api

import (
	"github.com/gin-gonic/gin"
	"github.com/putyy/ai-share/app/form"
	"github.com/putyy/ai-share/app/models"
	"github.com/putyy/ai-share/app/servers"
)

// 搜索豆瓣
func SearchFilm(c *gin.Context) {
	ip := c.ClientIP()

	var formData form.Film
	if err1 := c.ShouldBind(&formData); err1 != nil {
		ResponseError(c, "参数错误", err1.Error())
		return
	}

	insertFilmRequestLog(formData.Keyword, ip)

	data := make(map[string]interface{})
	data["list"] = servers.ListTopThree(formData.Keyword)
	ResponseSuccess(c, data)
}

// 插入log
func insertFilmRequestLog(keyword string, ip string) {
	model := models.FilmRequestLog{
		Keyword: keyword,
		Ip:      ip,
	}
	err := models.Db().Create(&model).Error
	if err != nil {
		return
	}
}

// 反馈关键词
func ReportKeyword(c *gin.Context) {
	ip := c.ClientIP()

	var formData form.Film
	if err1 := c.ShouldBind(&formData); err1 != nil {
		ResponseError(c, "参数错误", err1.Error())
		return
	}
	insertReportKeyword(formData.Keyword, ip)
}

// 插入关键词反馈
func insertReportKeyword(keyword string, ip string) {
	model := models.FilmRequestLog{
		Keyword: keyword,
		Ip:      ip,
	}
	err := models.Db().Create(&model).Error
	if err != nil {
		return
	}
}
