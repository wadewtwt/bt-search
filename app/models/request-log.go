package models

type RequestLog struct {
	ID      int    `gorm:"primary_key" json:"id,omitempty"`
	Keyword string `json:"keyword,omitempty"`
	Ip      string `json:"ip,omitempty"`
}

func (RequestLog) TableName() string {
	return "a_request_log"
}
