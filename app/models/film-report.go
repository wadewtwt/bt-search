package models

type FilmReport struct {
	ID         int    `gorm:"primary_key" json:"id,omitempty"`
	Keyword    string `json:"keyword,omitempty"`
	Ip         string `json:"ip,omitempty"`
	ReportType int    `json:"report_type,omitempty"`
}

func (FilmReport) TableName() string {
	return "a_film_report"
}
