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
		purchase := api.Group("/buyTicket")
		{
			purchase.POST(":id_ticket", c.buyTicket)
		}
		sess := api.Group("/sessions")
		{
			sess.POST("/", c.addSession)
			sess.GET("/", c.getSessions)
			sess.GET("/:id", c.getSession)
			sess.DELETE("/:id", c.deleteSession)
		}
	}
	return router
}
