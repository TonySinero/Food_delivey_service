package service

import (
	"context"
	"errors"
	"fmt"
	"food-delivery/internal/domain"
	"food-delivery/pkg/orderservice_ra"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"time"
)

func (s *OrderService) CreateConnectionRA() (orderservice_ra.OrderServiceClient, *grpc.ClientConn, context.Context, error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	conn, err := grpc.DialContext(ctx, fmt.Sprintf("%s:%s", s.cfg.GRPCRA.Host, s.cfg.GRPCRA.Port), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Error().Err(err).Msg("error occurred while creating conn to RA")
		return nil, nil, ctx, err
	}

	orderClient := orderservice_ra.NewOrderServiceClient(conn)

	return orderClient, conn, ctx, nil
}

func (s *OrderService) CreateFeedbackOnRestaurantRA(input domain.OrderFeedback) error {
	orderClientRA, conn, ctx, err := s.CreateConnectionRA()
	if err != nil {
		return err
	}

	orderFeedback := &orderservice_ra.OrderFeedbackOnRestaurantQuality{
		OrderID:  input.OrderID,
		Feedback: input.Feedback,
		Rating:   input.Rating,
	}

	if _, err := orderClientRA.AddRestaurantFeedback(ctx, orderFeedback); err != nil {
		log.Error().Err(err).Msg("error occurred while creating feedback in RA")

		return err
	}
	defer conn.Close()
	return nil
}

func (s *OrderService) GetOrderTotal(input domain.Order) (*float64, error) {
	orderClient, conn, ctx, err := s.CreateConnectionRA()
	if err != nil {
		return nil, err
	}

	protoDishes := make([]*orderservice_ra.Dish, 0, len(input.Dishes))
	for _, val := range input.Dishes {
		if val.Amount <= 0 {
			log.Error().Msg("invalid number of value")
			return nil, errors.New("invalid number of value")
		}
		protoDishes = append(protoDishes, &orderservice_ra.Dish{ID: val.DishID, Amount: int64(val.Amount)})
	}

	orderDishes := &orderservice_ra.OrderDishes{
		RestaurantID: input.RestaurantID,
		List:         protoDishes,
	}

	oc, err := orderClient.GetOrderTotal(ctx, orderDishes)
	if err != nil {
		log.Error().Err(err).Msg("error occurred while getting total from RA")

		return nil, err
	}

	defer conn.Close()

	return &oc.Total, nil
}
