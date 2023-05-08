package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusdsouza/goverallratings/schemas"
)

// @BasePath /api/v1
// @Summary Create show
// @Tags Shows
// @Accept json
// @Produce json
// @Param request body CreateShowRequest true "Request Body"
// @Success 200 {object} CreateShowResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /shows [post]
func CreateShowHandler(ctx *gin.Context) {
	cmd := CreateShowRequest{}

	ctx.BindJSON(&cmd)

	err := cmd.Validate()
	if err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	show := schemas.Show{
		Title:    cmd.Title,
		Episodes: cmd.Episodes,
		Seasons:  cmd.Seasons,
		Rating:   cmd.Rating,
		Genre:    cmd.Genre,
	}

	err = db.Create(&show).Error
	if err != nil {
		logger.Errorf("error creating show: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating show")
		return
	}

	sendSuccess(ctx, "Create Show", show)
}
