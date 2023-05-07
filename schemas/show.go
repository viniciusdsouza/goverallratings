package schemas

import (
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
