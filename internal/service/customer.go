package service

import (
	config "food-delivery/configs"
	"food-delivery/internal/domain"
	"food-delivery/internal/repository"
)

type CustomerService struct {
	repo repository.Customer
	cfg  *config.Config
}

func NewCustomerService(repo repository.Customer, cfg *config.Config) *CustomerService {
	return &CustomerService{repo: repo, cfg: cfg}
}

func (s *CustomerService) CreateCustomer(input domain.Customer) error {
	return s.repo.CreateCustomer(input)
}

func (s *CustomerService) GetCustomerByID(id string) (domain.CustomerInfo, error) {
	return s.repo.GetCustomerByID(id)
}

func (s *CustomerService) UpdateCustomer(input domain.CustomerUpdate, id string) (domain.CustomerInfo, error) {
	err := s.repo.UpdateCustomer(input, id)
	if err != nil {
		return domain.CustomerInfo{}, err
	}
	return s.repo.GetCustomerByID(id)
}
