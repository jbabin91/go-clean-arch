package models

import "time"

// Article defines the properties for the Author model
type Article struct {
	ID uint64 `json:"id"`
	Title string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
	Author Author `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}