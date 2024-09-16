package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"todo/pkg/service"
)

type Handler struct {
	service *service.Service
	log     zap.Logger
}

func NewHandler(service *service.Service, log zap.Logger) *Handler {
	return &Handler{
		service: service,
		log:     log,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.singUp)
		auth.POST("/sign-in", h.singIn)
	}
	api := router.Group("/api")
	{
		lists := api.Group("/list")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)
			items := lists.Group("/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
				items.GET("/:id", h.getItemById)
				items.PUT("/:id", h.updateItem)
				items.DELETE("/:id", h.deleteItem)
			}
		}
	}
	return router
}
