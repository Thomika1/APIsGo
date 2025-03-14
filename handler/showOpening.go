package handler

import (
	"fmt"
	"net/http"

	"github.com/Thomika1/APIsGo.git/schemas"
	"github.com/gin-gonic/gin"
)

// Show opening
func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	// Find opening
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	sendSuccess(ctx, "show-opening", opening)

} // Func Show Opening handler
