package schemas

import (
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Title    string
	Duration int64
	Rating   int64
	Genre    string
}
