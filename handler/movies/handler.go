package handler

import (
	"github.com/viniciusdsouza/goverallratings/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("movies handler")
	db = config.GetSQLite()
}
