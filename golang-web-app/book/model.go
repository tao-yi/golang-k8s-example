package book

import "time"

type Book struct {
	ID          int32     `json:"id"`
	Isbn        string    `json:"isbn"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
