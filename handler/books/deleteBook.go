package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusdsouza/goverallratings/schemas"
)

// @BasePath /api/v1
// @Summary Delete book
// @Tags Books
// @Accept json
// @Produce json
// @Success 200 {object} DeleteBookResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /books/:id [delete]
func DeleteBookHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParam").Error())
	}

	book := schemas.Book{}

	if err := db.First(&book, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("book with id: %s not found", id))
		return
	}

	if err := db.Delete(&book).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting book with id: %s", id))
		return
	}

	sendSuccess(ctx, "Delete Book", book)
}
