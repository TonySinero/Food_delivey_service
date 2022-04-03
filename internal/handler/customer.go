package handler

import (
	"food-delivery/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"net/http"
)

// CreateCustomer godoc
// @Tags customer
// @Description Create customer with provided data
// @Produce json
// @Param input body domain.Customer true "Customer data"
// @Success 200 {string} string "ok"
// @Failure 400 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /customer [post]
func (h *Handler) CreateCustomer(c *gin.Context) {
	var input domain.Customer
	if err := c.BindJSON(&input); err != nil {
		log.Error().Err(err).Msg("invalid json in create customer endpoint")
		newResponse(c, http.StatusBadRequest, "invalid request")
		return
	}

	var validate = validator.New()
	if err := validate.Struct(input); err != nil {
		log.Error().Err(err).Msg("invalid values of fields")
		newResponse(c, http.StatusBadRequest, "invalid values of fields")
		return
	}

	err := h.services.Customer.CreateCustomer(input)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, "failed to create customer")
		return
	}
	c.JSON(http.StatusOK, "The customer has been successfully created")
}

// UpdateCustomer godoc
// @Tags customer
// @Description Update customer data by id
// @Produce json
// @Param id path string true "Customer ID"
// @Param input body domain.CustomerUpdate true "Customer data"
// @Success 200 {object} customerResponse "ok"
// @Failure 400 {object} response
// @Failure 500 {object} response
// @Failure default {object} dataResponse
// @Router /customer/{id} [put]
func (h *Handler) UpdateCustomer(c *gin.Context) {
	var input domain.CustomerUpdate

	id := c.Param("id")

	if err := c.BindJSON(&input); err != nil {
		log.Error().Err(err).Msg("invalid json in update customer endpoint")
		newResponse(c, http.StatusBadRequest, "invalid request")
		return
	}

	customer, err := h.services.Customer.UpdateCustomer(input, id)
	if err != nil {
		switch err {
		case domain.ErrCustomerNotFound:
			newResponse(c, http.StatusBadRequest, err.Error())
			return
		default:
			newResponse(c, http.StatusInternalServerError, "failed to update customer")
			return
		}
	}
	c.JSON(http.StatusOK, customerResponse{
		Customer: customer,
	})
}

// GetCustomerByID godoc
// @Tags customer
// @Description Get customer data by id
// @Produce json
// @Param id path string true "Customer ID"
// @Success 200 {object} customerResponse "ok"
// @Failure 400 {object} response
// @Failure 500 {object} response
// @Failure default {object} dataResponse
// @Router /customer/{id} [get]
func (h *Handler) GetCustomerByID(c *gin.Context) {

	id := c.Param("id")

	customer, err := h.services.Customer.GetCustomerByID(id)
	if err != nil {
		switch err {
		case domain.ErrCustomerNotFound:
			newResponse(c, http.StatusBadRequest, err.Error())
			return
		default:
			newResponse(c, http.StatusInternalServerError, "failed to get customer")
			return
		}
	}

	c.JSON(http.StatusOK, customerResponse{
		Customer: customer,
	})
}
