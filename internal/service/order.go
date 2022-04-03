package service

import (
	config "food-delivery/configs"
	"food-delivery/internal/domain"
	"food-delivery/internal/repository"
)

type OrderService struct {
	repo repository.Order
	cfg  *config.Config
}

func NewOrderService(repo repository.Order, cfg *config.Config) *OrderService {
	return &OrderService{repo: repo, cfg: cfg}
}

func (s *OrderService) CreateOrder(input domain.Order) (string, error) {
	total, err := s.GetOrderTotal(input)
	if err != nil {
		return "", err
	}
	id, err := s.repo.CreateOrder(input, *total)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (s *OrderService) GetOrderByID(id string) (domain.GetOrderByID, error) {
	return s.repo.GetOrderByID(id)
}

func (s *OrderService) GetCustomerOrders(id, status string) ([]domain.GetAllOrders, error) {
	return s.repo.GetCustomerOrders(id, status)
}

func (s *OrderService) GetCustomerByID(id int) (domain.CustomerForRA, error) {
	return s.repo.GetCustomerByID(id)
}

func (s *OrderService) CreateFeedbackOnDeliveryQuality(input domain.OrderFeedback) error {
	return s.repo.CreateFeedbackOnDeliveryQuality(input)
}

func (s *OrderService) CreateFeedbackOnRestaurantQuality(input domain.OrderFeedback) error {
	if err := s.CreateFeedbackOnRestaurantRA(input); err != nil {
		return err
	}

	if err := s.repo.CreateFeedbackOnDeliveryQuality(input); err != nil {
		return err
	}

	return nil
}
