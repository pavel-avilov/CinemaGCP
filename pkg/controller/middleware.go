package controller

import (
	"CinemaGCP/pkg/logger"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (c *Controller) userCallMethod(ctx *gin.Context, userId uuid.UUID, methodName string) {
	user, err := c.service.Authorization.GetUser(userId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	logger.Infof("User: %s - call %s", user.Username, methodName)
}

func (c *Controller) userIdentity(ctx *gin.Context) {
	header := ctx.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(ctx, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newErrorResponse(ctx, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		newErrorResponse(ctx, http.StatusUnauthorized, "token is empty")
		return
	}

	userId, err := c.service.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	ctx.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (uuid.UUID, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return uuid.Nil, errors.New("user id not found")
	}

	idUUID, ok := id.(uuid.UUID)
	if !ok {
		return uuid.Nil, errors.New("user id is of invalid type")
	}

	return idUUID, nil
}
