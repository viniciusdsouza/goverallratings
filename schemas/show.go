package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Show struct {
	gorm.Model
	Title    string
	Episodes int64
	Seasons  int64
	Rating   int64
	Genre    string
}

type ShowResponse struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	Title     string     `json:"title"`
	Episodes  int64      `json:"episodes"`
	Seasons   int64      `json:"seasons"`
	Rating    int64      `json:"rating"`
	Genre     string     `json:"genre"`
}
