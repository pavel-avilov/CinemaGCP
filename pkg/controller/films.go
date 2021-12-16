package controller

import (
	"CinemaGCP/pkg/src"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getAllFilmsResponse struct {
	Data []src.Film `json:"data"`
}

func (c *Controller) getAllFilms(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	c.userCallMethod(ctx, userId, "getAllFilms()")
	lists, err := c.service.Film.GetAll()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, getAllFilmsResponse{
		Data: lists,
	})

}
