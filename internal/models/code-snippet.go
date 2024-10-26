package models

import "time"

type CodeSnippet struct {
	ID          int       `gorm:"primaryKey;autoIncrement"`
	UserID      int       `gorm:"not null;column:user_id"`
	Language    string    `gorm:"not null;column:language"`
	Title       string    `gorm:"not null;column:title"`
	CodeContent string    `gorm:"type:text;not null;column:code_snippet"`
	Note 		string 	  `gorm:"type:text;column:note"`
	CreatedAt   *time.Time `gorm:"type:TIMESTAMP;column:created_at"`
	UpdatedAt   *time.Time `gorm:"type:TIMESTAMP;column:updated_at"`
}
