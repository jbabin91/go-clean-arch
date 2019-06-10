package models

// Author defines the properties of the Author model
type Author struct {
	ID uint64 `json:"id"`
	Name string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}