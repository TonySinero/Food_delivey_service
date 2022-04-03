package proto

import (
	"context"
	"database/sql"
	"fmt"
	"food-delivery/internal/domain"
	"food-delivery/pkg/paymentservice"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PaymentPostgres struct {
	db *sqlx.DB
}

func NewPaymentPostgres(db *sqlx.DB) *PaymentPostgres {
	return &PaymentPostgres{db: db}
}

func (r *PaymentPostgres) ChangeStatus(ctx context.Context, status *paymentservice.PaymentResult) (*emptypb.Empty, error) {
	val := 1
	if status.Answer == false {
		val = 5
	}
	query := fmt.Sprintf("UPDATE orders SET status_id=$1 WHERE id = $2")
	_, err := r.db.Exec(query, val, status.IdOrder)
	if err != nil {
		log.Error().Err(err).Msg("error occurred while updating order")
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (r *PaymentPostgres) GetCustomerByID(id int) (domain.CustomerForRA, error) {
	var customer domain.CustomerForRA
	SelectCustomerQuery := fmt.Sprintf("SELECT full_name, phone_number FROM customers WHERE id = $1")

	if err := r.db.Get(&customer, SelectCustomerQuery, id); err != nil {
		log.Error().Err(err).Msg("error occurred while selecting customer")
		if err == sql.ErrNoRows {
			return customer, domain.ErrCustomerNotFound
		}
	}
	return customer, nil
}

func (r *PaymentPostgres) GetOrderByID(id string) (domain.Order, error) {
	var order domain.Order
	query := fmt.Sprintf(`SELECT customer_id, restaurant_id,
			address, required_time FROM orders WHERE id=$1`)
	err := r.db.Get(&order, query, id)
	if err != nil {
		log.Error().Err(err).Msg("error occurred while getting order")
	}

	queryDishes := fmt.Sprintf(`SELECT dish_id,
			amount FROM order_dishes WHERE order_id=$1`)
	err = r.db.Select(&order.Dishes, queryDishes, id)
	if err != nil {
		log.Error().Err(err).Msg("error occurred while getting dishes from customer order")
	}

	fmt.Printf("%+v\n", order)
	return order, err
}
