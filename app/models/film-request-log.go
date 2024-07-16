package models

type FilmRequestLog struct {
	ID      int    `gorm:"primary_key" json:"id,omitempty"`
	Keyword string `json:"keyword,omitempty"`
	Ip      string `json:"ip,omitempty"`
}

func (FilmRequestLog) TableName() string {
	return "a_film_request_log"
}
