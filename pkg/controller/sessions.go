package controller

import (
	"CinemaGCP/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type getAllSessionsResponse struct {
	Data interface{}
}

func (c *Controller) getAllSessions(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	c.userCallMethod(ctx, userId, "getAllSessions()")

	lists, err := c.service.Session.GetAll()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, getAllSessionsResponse{
		Data: lists,
	})
}

func (c *Controller) getSession(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	c.userCallMethod(ctx, userId, "getSessionById()")

	sessionId, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid session id")
	}
	logger.Info(sessionId)
	lists, err := c.service.Session.GetSessionById(sessionId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, lists)

}
