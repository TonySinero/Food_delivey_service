package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"food-delivery/internal/domain"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"strconv"
)

type OrderPostgres struct {
	db *sqlx.DB
}

func NewOrderPostgres(db *sqlx.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}

func (r *OrderPostgres) GetCustomerByID(id int) (domain.CustomerForRA, error) {
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

func (r *OrderPostgres) CreateOrder(input domain.Order, total float64) (string, error) {
	fmt.Printf("%+v\n", input)
	var id string
	createOrderQuery := fmt.Sprintf(`INSERT INTO orders(customer_id, restaurant_id, address, cost,
			required_time, status_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`)

	row := r.db.QueryRow(createOrderQuery, input.CustomerID, input.RestaurantID, input.Address, total,
		input.RequiredTime.Format("2006-01-02 15:04:05"), 1)

	if err := row.Scan(&id); err != nil {
		log.Error().Err(err).Msg("error occurred while inserting order, order already exist")
		return "", errors.New("invalid number of value")
	}

	var insertOrderDishes string
	for key, val := range input.Dishes {
		if val.Amount <= 0 {
			log.Error().Msg("invalid number of value")
			return "", errors.New("invalid number of value")
		}
		if key == 0 {
			insertOrderDishes += "('" + id + "','" + val.DishID + "'," + strconv.Itoa(val.Amount) + ")"
		} else {
			insertOrderDishes += ",('" + id + "','" + val.DishID + "'," + strconv.Itoa(val.Amount) + ")"
		}
	}

	if insertOrderDishes != "" {
		createOrdersDishesQuery := fmt.Sprintf("INSERT INTO order_dishes (order_id, dish_id, amount) VALUES %s",
			insertOrderDishes)
		if _, err := r.db.Exec(createOrdersDishesQuery); err != nil {
			log.Error().Err(err).Msg("error occurred while creating orders dishes in food-delivery db")

			return "", err
		}
	}

	return id, nil
}

func (r *OrderPostgres) GetOrderByID(id string) (domain.GetOrderByID, error) {
	var order domain.GetOrderByID
	query := fmt.Sprintf(`SELECT id, customer_id, courier_id,
			payment_id, status_id, address, cost, required_time, fact_time, 
			created_at FROM orders WHERE id=$1`)
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

func (r *OrderPostgres) GetCustomerOrders(id, status string) ([]domain.GetAllOrders, error) {
	var orders []domain.GetAllOrders

	if status != "" {
		status = fmt.Sprintf(`AND status_id=%s`, status)
	}

	query := fmt.Sprintf(`SELECT id, customer_id, status_id, address, 
			cost, created_at FROM orders WHERE customer_id='%s' %s ORDER BY created_at desc`, id, status)
	err := r.db.Select(&orders, query)
	if err != nil {
		log.Error().Err(err).Msg("error occurred while getting orders")
	}

	OrderDishesQuery := fmt.Sprintf(`SELECT dish_id,
			amount FROM order_dishes WHERE order_id=$1`)
	for orderID := range orders {
		err = r.db.Select(&orders[orderID].Dishes, OrderDishesQuery, &orders[orderID].ID)
		if err != nil {
			log.Error().Err(err).Msg("error occurred while getting order dishes info")
		}
	}
	fmt.Printf("%+v\n", orders)
	return orders, err
}

func (r *OrderPostgres) CreateFeedbackOnDeliveryQuality(input domain.OrderFeedback) error {
	updateOrderQuery := fmt.Sprintf("INSERT INTO order_feedback(order_id, feedback, rating) VALUES ($1, $2, $3)")
	_, err := r.db.Exec(updateOrderQuery, &input.OrderID, &input.Feedback, &input.Rating)
	if err != nil {
		log.Error().Err(err).Msg("error occurred while getting order dishes info")
	}

	return err
}

func (r *OrderPostgres) CreateFeedbackOnRestaurantQuality(input domain.OrderFeedback) error {
	updateOrderQuery := fmt.Sprintf("INSERT INTO order_feedback(order_id, feedback, rating) VALUES ($1, $2, $3)")
	_, err := r.db.Exec(updateOrderQuery, &input.OrderID, &input.Feedback, &input.Rating)
	if err != nil {
		log.Error().Err(err).Msg("error occurred while getting order dishes info")
	}

	return err
}
