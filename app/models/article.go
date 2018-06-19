package models


type Article struct {
	Id     uint64 `gorm:"primary_key" json:"id"`
	Name  string `json:"name"`
	Fetish  string `json:"fetish"`
}

