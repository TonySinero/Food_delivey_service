package service

import (
	config "food-delivery/configs"
	"food-delivery/internal/domain"
	"food-delivery/internal/repository"
)

type Order interface {
	CreateOrder(input domain.Order) (string, error)
	GetOrderByID(id string) (domain.GetOrderByID, error)
	GetCustomerOrders(id, status string) ([]domain.GetAllOrders, error)
	GetCustomerByID(id int) (domain.CustomerForRA, error)
	CreateFeedbackOnDeliveryQuality(input domain.OrderFeedback) error
	CreateFeedbackOnRestaurantQuality(input domain.OrderFeedback) error
}

type Customer interface {
	CreateCustomer(input domain.Customer) error
	GetCustomerByID(id string) (domain.CustomerInfo, error)
	UpdateCustomer(input domain.CustomerUpdate, id string) (domain.CustomerInfo, error)
}

type Service struct {
	Order
	Customer
}

func NewService(repos *repository.Repository, cfg *config.Config) *Service {
	return &Service{
		Order:    NewOrderService(repos.Order, cfg),
		Customer: NewCustomerService(repos.Customer, cfg),
	}
}
