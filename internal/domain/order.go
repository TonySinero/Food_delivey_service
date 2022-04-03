package domain

import (
	"time"
)

type CustomerForRA struct {
	FullName    string `json:"fullName" db:"full_name"`
	PhoneNumber string `json:"phoneNumber" db:"phone_number"`
}

type Order struct {
	CustomerID   int         `json:"customerID" db:"customer_id" validate:"required"`
	RestaurantID string      `json:"restaurantID" db:"restaurant_id" validate:"required,uuid"`
	Address      string      `json:"address" db:"address" validate:"required"`
	RequiredTime time.Time   `json:"requiredTime" db:"required_time" validate:"required"`
	Dishes       []OrderDish `json:"dishes" db:"order_dishes" validate:"required,min=1,dive"`
}

type OrderDish struct {
	DishID string `json:"dishID" db:"dish_id" validate:"required"`
	Amount int    `json:"amount" db:"amount" validate:"required"`
}

type GetOrderByID struct {
	ID           string      `json:"id" db:"id"`
	CustomerID   string      `json:"customerID" db:"customer_id"`
	CourierID    *string     `json:"courierID" db:"courier_id"`
	PaymentID    *string     `json:"paymentID" db:"payment_id"`
	StatusID     *int        `json:"statusID" db:"status_id"`
	Address      string      `json:"address" db:"address"`
	Cost         float64     `json:"cost" db:"cost"`
	RequiredTime time.Time   `json:"requiredTime" db:"required_time"`
	FactTime     *time.Time  `json:"factTime" db:"fact_time"`
	CreatedAt    time.Time   `json:"createdAt" db:"created_at"`
	Dishes       []OrderDish `json:"orderDishes" db:"order_dishes"`
}

type GetAllOrders struct {
	ID         string      `json:"id" db:"id"`
	CustomerID string      `json:"customerID" db:"customer_id"`
	StatusID   *int        `json:"statusID" db:"status_id"`
	Address    string      `json:"address" db:"address"`
	Cost       float64     `json:"cost" db:"cost"`
	CreatedAt  time.Time   `json:"createdAt" db:"created_at"`
	Dishes     []OrderDish `json:"orderDishes" db:"order_dishes"`
}

type OrderStatus struct {
	CustomerID string `json:"customerID" form:"customerID" validate:"required"`
	Status     string `json:"status" form:"status"`
}

type OrderFeedback struct {
	OrderID  string  `json:"orderID" validate:"required""`
	Feedback string  `json:"feedback"`
	Rating   float32 `json:"rating"`
}
