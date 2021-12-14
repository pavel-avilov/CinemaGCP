package controller

import (
	"github.com/gin-gonic/gin"
	"log"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(ctx *gin.Context, statusCode int, message string) {
	log.Printf("ERROR %v \n", message)
	ctx.AbortWithStatusJSON(statusCode, errorResponse{message})
}