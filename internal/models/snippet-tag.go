package models

import "time"

type SnippetTag struct {
	ID            int        `gorm:"primaryKey"`
	Name          string     `gorm:"not null;column:name"`
	CodeSnippetID int        `gorm:"not null;column:code_snippets"`
	CreatedAt     *time.Time `gorm:"type:TIMESTAMP;column:created_at"`
	UpdatedAt     *time.Time `gorm:"type:TIMESTAMP;column:updated_at"`
}