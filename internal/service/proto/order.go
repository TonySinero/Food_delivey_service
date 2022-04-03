package proto

import (
	"context"
	"food-delivery/internal/repository/proto"
	"food-delivery/pkg/orderservice_fd"
	"google.golang.org/protobuf/types/known/emptypb"
)

type OrderService struct {
	repo *proto.Repository
}

func NewOrderService(repo *proto.Repository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) UpdateOrder(ctx context.Context, in *orderservice_fd.UpdateOrderMessage) (*emptypb.Empty, error) {
	return s.repo.UpdateOrder(ctx, in)
}
