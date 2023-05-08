package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusdsouza/goverallratings/schemas"
)

// @BasePath /api/v1
// @Summary Delete show
// @Tags Shows
// @Accept json
// @Produce json
// @Success 200 {object} DeleteShowResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /shows/:id [delete]
func DeleteShowHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParam").Error())
	}

	show := schemas.Show{}

	if err := db.First(&show, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("show with id: %s not found", id))
		return
	}

	if err := db.Delete(&show).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting show with id: %s", id))
		return
	}

	sendSuccess(ctx, "Delete Show", show)
}
