package controller

import (
	"CinemaGCP/pkg/src"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getAllFilmsResponse struct {
	Data []src.Film `json:"data"`
}

func (c *Controller) getFilms(ctx *gin.Context) {

	lists, err := c.service.Film.GetAll()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, getAllFilmsResponse{
		Data: lists,
	})

}
