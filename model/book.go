package model

type Book struct {
	ID          uint64 `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Authors     string `json:"authors,omitempty"`
	Description string `json:"description,omitempty"`
	UserID      uint64 `json:"user_id,omitempty"`
}
