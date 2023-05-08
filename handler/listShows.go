package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusdsouza/goverallratings/schemas"
)

// @BasePath /api/v1
// @Summary List shows
// @Tags Shows
// @Accept json
// @Produce json
// @Success 200 {object} ListShowResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /shows [get]
func ListShowsHandler(ctx *gin.Context) {
	shows := []schemas.Show{}

	if err := db.Find(&shows).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing shows")
		return
	}

	sendSuccess(ctx, "Get List Shows", shows)
}
