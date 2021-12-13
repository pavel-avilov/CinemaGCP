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

func (con *Controller) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", con.signIn)
		auth.POST("/sign-up", con.signUp)
	}

	api := router.Group("/api")
	{
		purchase := api.Group("/buyTicket")
		{
			purchase.POST(":id_ticket", con.buyTicket)
		}
		sess := api.Group("/sessions")
		{
			sess.POST("/", con.addSession)
			sess.GET("/", con.getSessions)
			sess.GET("/:id", con.getSession)
			sess.DELETE("/:id", con.deleteSession)
		}
	}
	return router
}
