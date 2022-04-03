package handler

import (
	_ "food-delivery/docs"
	"food-delivery/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(
		corsMiddleware,
	)

	customer := router.Group("/customer")
	{
		customer.POST("/", h.CreateCustomer)
		customer.GET("/:id", h.GetCustomerByID)
		customer.PUT("/:id", h.UpdateCustomer)
		customer.GET("/orders/:id", h.GetCustomerOrders)
	}

	order := router.Group("/order")
	{
		order.POST("/", h.CreateOrder)
		order.GET("/:id", h.GetOrderByID)
		order.POST("/feedback/delivery-quality", h.FeedbackOnDeliveryQuality)
		order.POST("/feedback/restaurant-quality", h.FeedbackOnRestaurantQuality)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
