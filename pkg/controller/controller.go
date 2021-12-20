package controller

import (
	"CinemaGCP/pkg/service"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service *service.Service
}

func NewController(s *service.Service) *Controller {
	return &Controller{service: s}
}

func (c *Controller) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", c.signIn)
		auth.POST("/sign-up", c.signUp)
	}

	api := router.Group("/api", c.userIdentity)
	{
		purchase := api.Group("/tickets")
		{
			purchase.POST("/buy", c.buyTicket)
		}

		film := api.Group("/films")
		{
			film.GET("/", c.getAllFilms)
		}

		sess := api.Group("/sessions")
		{
			sess.GET("/", c.getAllSessions)
			sess.GET("/:id", c.getSession)
		}
	}
	return router
}
