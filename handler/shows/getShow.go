package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusdsouza/goverallratings/schemas"
)

// @BasePath /api/v1
// @Summary Get show
// @Tags Shows
// @Accept json
// @Produce json
// @Success 200 {object} GetShowResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /shows/:id [get]
func GetShowHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParam").Error())
	}

	show := schemas.Show{}
	if err := db.First(&show, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("show with id: %s not found", id))
		return
	}

	sendSuccess(ctx, "Get Show", show)
}
