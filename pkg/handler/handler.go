package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
	}

	api := router.Group("/api")
	{
		purchase := api.Group("/buyTicket")
		{
			purchase.POST(":id_ticket", h.buyTicket)
		}
		sess := api.Group("/sessions")
		{
			sess.POST("/", h.addSession)
			sess.GET("/", h.getSessions)
			sess.GET("/:id", h.getSession)
			sess.DELETE("/:id", h.deleteSession)
		}
	}
	return router
}
