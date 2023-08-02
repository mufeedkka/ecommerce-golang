package models

type Products struct {
	// gorm.Model
	short_name string `json:"shortName" gorm:"not null" validate:"required,min=3"`
	long_name  string `json:"longName" gorm:"not null" validate:"required,min=3"`
}
