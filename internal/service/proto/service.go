package proto

import (
	"context"
	config "food-delivery/configs"
	"food-delivery/internal/repository/proto"
	"food-delivery/pkg/orderservice_fd"
	"food-delivery/pkg/paymentservice"
	"food-delivery/pkg/userservice"
	"google.golang.org/protobuf/types/known/emptypb"
)

type User interface {
	CreateUser(ctx context.Context, order *userservice.User) (*userservice.UserState, error)
}

type Order interface {
	UpdateOrder(ctx context.Context, in *orderservice_fd.UpdateOrderMessage) (*emptypb.Empty, error)
}

type Payment interface {
	ChangeStatus(ctx context.Context, status *paymentservice.PaymentResult) (*emptypb.Empty, error)
}

type Service struct {
	User
	userservice.UnsafeUserServiceServer
	Order
	orderservice_fd.UnsafeOrderServiceFDServer
	Payment
	paymentservice.UnsafePaymentServiceServer
}

func NewService(repo *proto.Repository, cfg *config.Config) *Service {
	return &Service{
		User:    NewUserService(repo),
		Order:   NewOrderService(repo),
		Payment: NewPaymentService(repo, cfg),
	}
}
