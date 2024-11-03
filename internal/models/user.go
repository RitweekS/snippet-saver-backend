package models

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	Image string `json:"image"`
}