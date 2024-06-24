package servers

import (
	"github.com/putyy/ai-share/app/models"
)

func ListTopThree(keyword string) (resList []models.Film) {
	models.Db().Where("title like ?", "%"+keyword+"%").Select("*").Limit(3).Order("id desc").Find(&resList)
	return
}
