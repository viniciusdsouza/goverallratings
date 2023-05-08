package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title      string
	AuthorName string
	Rating     int64
	Genre      string
}

type BookResponse struct {
	ID         uint       `json:"id"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
	DeletedAt  *time.Time `json:"deletedAt,omitempty"`
	Title      string     `json:"title"`
	AuthorName string      `json:"authorName"`
	Rating     int64      `json:"rating"`
	Genre      string     `json:"genre"`
}
