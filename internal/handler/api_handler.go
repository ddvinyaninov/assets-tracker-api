// internal/handler/api_handler.go
package handler

import (
	"net/http"

	"ddvinyaninov/assets-tracker-api/internal/domain"

	"github.com/gin-gonic/gin"
)

type ApiHandler struct {
	apiService domain.ApiUsecase
}

func NewApiHandler(apiService domain.ApiUsecase) *ApiHandler {
	return &ApiHandler{apiService: apiService}
}

func (h ApiHandler) List(ctx *gin.Context) {
	features, err := h.apiService.List(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"type":     "FeatureCollection",
		"features": features,
	})
}
