package models

type Article struct {
	Id     uint64 `gorm:"primary_key" json:"id"`
	Name  string `json:"title"`
	Fetish string `json:"text"`
}

