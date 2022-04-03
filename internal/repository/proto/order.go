package proto

import (
	"fmt"
	"food-delivery/pkg/orderservice_fd"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type OrderPostgres struct {
	db *sqlx.DB
}

func NewOrderPostgres(db *sqlx.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}

func (r *OrderPostgres) UpdateOrder(ctx context.Context, in *orderservice_fd.UpdateOrderMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	updateOrderQuery := fmt.Sprintf("UPDATE orders SET status_id = $1 WHERE id = $2")
	if _, err := r.db.Exec(updateOrderQuery, in.Status, in.OrderUUID); err != nil {
		log.Error().Err(err).Msg("error occurred while updating order status in FD db")

		return nil, err
	}

	return &emptypb.Empty{}, nil
}
