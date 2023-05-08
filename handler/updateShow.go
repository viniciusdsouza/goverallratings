package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusdsouza/goverallratings/schemas"
)

// @BasePath /api/v1
// @Summary Update show
// @Tags Shows
// @Accept json
// @Produce json
// @Param request body UpdateShowRequest true "Request Body"
// @Success 200 {object} UpdateShowResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /shows/:id [put]
func UpdateShowHandler(ctx *gin.Context) {
	cmd := UpdateShowRequest{}

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

	show := schemas.Show{}

	if err := db.First(&show, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("show with id: %s not found", id))
		return
	}

	if show.Title != "" {
		show.Title = cmd.Title
	}

	if show.Episodes > 0 {
		show.Episodes = cmd.Episodes
	}

	if show.Seasons > 0 {
		show.Seasons = cmd.Seasons
	}

	if show.Rating > 0 {
		show.Rating = cmd.Rating
	}

	if show.Genre != "" {
		show.Genre = cmd.Genre
	}

	if err := db.Save(&show).Error; err != nil {
		logger.Errorf("error updating show: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating show")
		return
	}

	sendSuccess(ctx, "Update Show", show)
}
