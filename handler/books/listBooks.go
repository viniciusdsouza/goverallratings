package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusdsouza/goverallratings/schemas"
)

// @BasePath /api/v1
// @Summary List books
// @Tags Books
// @Accept json
// @Produce json
// @Success 200 {object} ListBookResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /books [get]
func ListBooksHandler(ctx *gin.Context) {
	books := []schemas.Book{}

	if err := db.Find(&books).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing books")
		return
	}

	sendSuccess(ctx, "Get List Books", books)
}
