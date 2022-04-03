package proto

import (
	"context"
	"food-delivery/internal/domain"
	"food-delivery/pkg/orderservice_fd"
	"food-delivery/pkg/paymentservice"
	"food-delivery/pkg/userservice"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type User interface {
	CreateUser(ctx context.Context, order *userservice.User) (*userservice.UserState, error)
}

type Order interface {
	UpdateOrder(ctx context.Context, in *orderservice_fd.UpdateOrderMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type Payment interface {
	ChangeStatus(ctx context.Context, status *paymentservice.PaymentResult) (*emptypb.Empty, error)
	GetCustomerByID(id int) (domain.CustomerForRA, error)
	GetOrderByID(id string) (domain.Order, error)
}

type Repository struct {
	User
	Order
	Payment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:    NewUserPostgres(db),
		Order:   NewOrderPostgres(db),
		Payment: NewPaymentPostgres(db),
	}
}
