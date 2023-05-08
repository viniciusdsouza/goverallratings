package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusdsouza/goverallratings/schemas"
)

// @BasePath /api/v1
// @Summary Update book
// @Tags Books
// @Accept json
// @Produce json
// @Param request body UpdateBookRequest true "Request Body"
// @Success 200 {object} UpdateBookResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /books/:id [put]
func UpdateBookHandler(ctx *gin.Context) {
	cmd := UpdateBookRequest{}

	ctx.BindJSON(&cmd)

	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParam").Error())
	}

	err := cmd.Validate()

	if err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	book := schemas.Book{}

	if err := db.First(&book, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("book with id: %s not found", id))
		return
	}

	if book.Title != "" {
		book.Title = cmd.Title
	}

	if book.AuthorName != "" {
		book.AuthorName = cmd.AuthorName
	}

	if book.Rating > 0 {
		book.Rating = cmd.Rating
	}

	if book.Genre != "" {
		book.Genre = cmd.Genre
	}

	if err := db.Save(&book).Error; err != nil {
		logger.Errorf("error updating book: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating book")
		return
	}

	sendSuccess(ctx, "Update Book", book)
}
