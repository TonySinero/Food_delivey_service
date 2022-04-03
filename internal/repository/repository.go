package repository

import (
	"food-delivery/internal/domain"
	"github.com/jmoiron/sqlx"
)

type Order interface {
	GetCustomerOrders(id, status string) ([]domain.GetAllOrders, error)
	CreateOrder(input domain.Order, total float64) (string, error)
	GetOrderByID(id string) (domain.GetOrderByID, error)
	GetCustomerByID(id int) (domain.CustomerForRA, error)
	CreateFeedbackOnDeliveryQuality(input domain.OrderFeedback) error
	CreateFeedbackOnRestaurantQuality(input domain.OrderFeedback) error
}

type Customer interface {
	CreateCustomer(input domain.Customer) error
	GetCustomerByID(id string) (domain.CustomerInfo, error)
	UpdateCustomer(input domain.CustomerUpdate, id string) error
}

type Repository struct {
	Order
	Customer
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Order:    NewOrderPostgres(db),
		Customer: NewCustomerPostgres(db),
	}
}
