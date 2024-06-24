package models

type Film struct {
	ID         int      `gorm:"primary_key" json:"id,omitempty"`
	Title      string   `json:"title,omitempty"`
	Cover      string   `json:"cover,omitempty"`
	Url        string   `json:"url,omitempty"`
	Rating     float64  `json:"rating,omitempty"`
	Casts      string   `json:"casts,omitempty"`
	Star       int      `json:"star,omitempty"`
	Directors  string   `json:"directors,omitempty"`
	CoverX     int      `json:"cover_x,omitempty"`
	CoverY     int      `json:"cover_y,omitempty"`
	OtherId    int      `json:"other_id,omitempty"`
	CreateTime JsonTime `json:"create_time,omitempty"`
}

func (Film) TableName() string {
	return "a_film"
}
