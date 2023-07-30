package main

type Products struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"type:varchar(100)"`
	Description string  `gorm:"type:text"`
	Category    string  `gorm:"type:varchar(50)"`
	Price       float64 `gorm:"type:decimal(7,2)"`
	Stock       int     `gorm:"type:int"`
	AvgRating   float64 `gorm:"type:decimal(3,1);default:0"`
	RatingCount int     `gorm:"type:int"`
	Slug        string  `gorm:"type:varchar(200)"`
}
