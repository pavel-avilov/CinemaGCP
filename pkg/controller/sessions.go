package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *Controller) buyTicket(ctx *gin.Context) {

}

type getAllSessionsResponse struct {
	Data []map[string]string
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

}
