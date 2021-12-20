package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type buyTicketInput struct {
	SessionId uuid.UUID `json:"session_id" binding:"required"`
	Place     int       `json:"place" binding:"required"`
}

func (c *Controller) buyTicket(ctx *gin.Context) {
	var input buyTicketInput

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	userId, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	c.userCallMethod(ctx, userId, "buyTicket()")

	err = c.service.Buy(userId, input.SessionId, input.Place)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResponse{Status: "Ticket was purchased"})
}
