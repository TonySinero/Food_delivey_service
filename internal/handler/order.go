package handler

import (
	"food-delivery/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"net/http"
)

// CreateOrder godoc
// @Summary Create new order
// @Tags order
// @Description Create new order
// @Produce json
// @Param input body domain.Order true "Order info"
// @Success 200 {object} dataResponse "ok"
// @Failure 400 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /order [post]
func (h *Handler) CreateOrder(c *gin.Context) {

	var input domain.Order
	if err := c.BindJSON(&input); err != nil {
		log.Error().Err(err).Msg("invalid json in create order endpoint")
		newResponse(c, http.StatusBadRequest, "invalid request")
		return
	}

	var validate = validator.New()
	if err := validate.Struct(input); err != nil {
		log.Error().Err(err).Msg("invalid values of fields")
		newResponse(c, http.StatusBadRequest, "invalid values of fields")
		return
	}

	id, err := h.services.Order.CreateOrder(input)
	if err != nil {
		switch err {
		case domain.ErrCustomerNotFound:
			newResponse(c, http.StatusBadRequest, err.Error())
			return
		default:
			newResponse(c, http.StatusInternalServerError, "failed to create order")
			return
		}
	}
	c.JSON(http.StatusOK, dataResponse{
		Data: id,
	})
}

// GetOrderByID godoc
// @Summary Get Order by ID
// @Tags order
// @Description Get Order by order ID
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} orderResponse "ok"
// @Failure 400 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /order/{id} [get]
func (h *Handler) GetOrderByID(c *gin.Context) {

	id := c.Param("id")

	order, err := h.services.Order.GetOrderByID(id)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, orderResponse{
		Order: order,
	})
}

// GetCustomerOrders godoc
// @Tags customer
// @Description Get Orders by customer ID and order status
// @Produce json
// @Param id path string true "Customer ID"
// @Param status query string false "status"
// @Success 200 {object} ordersResponse "ok"
// @Failure 400 {object} response
// @Failure 500 {object} response
// @Failure default {object} dataResponse
// @Router /customer/orders/{id} [get]
func (h *Handler) GetCustomerOrders(c *gin.Context) {

	id := c.Param("id")

	status := c.Query("status")

	orders, err := h.services.Order.GetCustomerOrders(id, status)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, ordersResponse{
		Orders: orders,
	})
}

// FeedbackOnDeliveryQuality godoc
// @Summary Create Feedback On Delivery Quality
// @Tags order
// @Description Create feedback on delivery quality by order id and feedback
// @Produce json
// @Param orderID query string true "customer ID"
// @Param feedback query string false "feedback"
// @Success 201 {object} null "ok"
// @Failure 400 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /order/feedback/delivery-quality [post]
func (h *Handler) FeedbackOnDeliveryQuality(c *gin.Context){
	var input domain.OrderFeedback

	if err := c.Bind(&input); err != nil {
		log.Error().Err(err).Msg("invalid keys in get all orders endpoint")
		newResponse(c, http.StatusBadRequest, "invalid request")
		return
	}

	var validate = validator.New()
	if err := validate.Struct(input); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Order.CreateFeedbackOnDeliveryQuality(input)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *Handler) FeedbackOnRestaurantQuality(c *gin.Context){
	var input domain.OrderFeedback

	if err := c.Bind(&input); err != nil {
		log.Error().Err(err).Msg("invalid keys in create feedback")
		newResponse(c, http.StatusBadRequest, "invalid request")
		return
	}

	var validate = validator.New()
	if err := validate.Struct(input); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Order.CreateFeedbackOnRestaurantQuality(input)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}
