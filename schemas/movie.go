package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Title    string
	Duration int64
	Rating   int64
	Genre    string
}

type MovieResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	Title     string    `json:"title"`
	Duration  int64     `json:"duration"`
	Rating    int64     `json:"rating"`
	Genre     string    `json:"genre"`
}
