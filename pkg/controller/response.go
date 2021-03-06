package controller

import (
	"CinemaGCP/pkg/logger"
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(ctx *gin.Context, statusCode int, message string) {
	logger.Error(message)
	ctx.AbortWithStatusJSON(statusCode, errorResponse{message})
}
