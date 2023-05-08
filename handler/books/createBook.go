package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusdsouza/goverallratings/schemas"
)

// @BasePath /api/v1
// @Summary Create book
// @Tags Books
// @Accept json
// @Produce json
// @Param request body CreateBookRequest true "Request Body"
// @Success 200 {object} CreateBookResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /books [post]
func CreateBookHandler(ctx *gin.Context) {
	cmd := CreateBookRequest{}

	ctx.BindJSON(&cmd)

	err := cmd.Validate()
	if err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	book := schemas.Book{
		Title:      cmd.Title,
		AuthorName: cmd.AuthorName,
		Rating:     cmd.Rating,
		Genre:      cmd.Genre,
	}

	err = db.Create(&book).Error
	if err != nil {
		logger.Errorf("error creating book: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating book")
		return
	}

	sendSuccess(ctx, "Create Book", book)
}
