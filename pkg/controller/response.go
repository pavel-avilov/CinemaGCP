package controller

import (
	"CinemaGCP/pkg/logger"
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(ctx *gin.Context, statusCode int, message string) {
	logger.Error(message)
	ctx.AbortWithStatusJSON(statusCode, errorResponse{message})
}
