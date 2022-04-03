package handler

import (
	"food-delivery/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type dataResponse struct {
	Data interface{} `json:"data"`
}

type orderResponse struct {
	Order domain.GetOrderByID `json:"data"`
}

type ordersResponse struct {
	Orders []domain.GetAllOrders `json:"data"`
}

type customerResponse struct {
	Customer domain.CustomerInfo `json:"data"`
}

type response struct {
	Message string `json:"message"`
}

func newResponse(c *gin.Context, statusCode int, message string) {
	log.Error().Msg(message)
	c.AbortWithStatusJSON(statusCode, response{message})
}
