package proto

import (
	"context"
	"errors"
	"fmt"
	config "food-delivery/configs"
	"food-delivery/internal/domain"
	"food-delivery/internal/repository/proto"
	"food-delivery/pkg/orderservice_ra"
	"food-delivery/pkg/paymentservice"
	"github.com/golang/protobuf/ptypes"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

type PaymentService struct {
	repo *proto.Repository
	cfg  *config.Config
}

func NewPaymentService(repo *proto.Repository, cfg *config.Config) *PaymentService {
	return &PaymentService{repo: repo, cfg: cfg}
}

func (s *PaymentService) ChangeStatus(ctx context.Context, status *paymentservice.PaymentResult) (*emptypb.Empty, error) {
	if status.Answer == true {
		input, err := s.repo.GetOrderByID(status.IdOrder)
		if err != nil {
			log.Error().Err(err).Msg("error occurred while getting order from db")
			return &emptypb.Empty{}, err
		}

		err = s.CreateOrderRA(status, input)
		if err != nil {
			log.Error().Err(err).Msg("error occurred while creating conn to RA")
			return &emptypb.Empty{}, err
		}
	}
	return s.repo.ChangeStatus(ctx, status)
}

func (s *PaymentService) CreateConnectionRA() (orderservice_ra.OrderServiceClient, *grpc.ClientConn, context.Context, error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	conn, err := grpc.DialContext(ctx, fmt.Sprintf("%s:%s", s.cfg.GRPCRA.Host, s.cfg.GRPCRA.Port), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Error().Err(err).Msg("error occurred while creating conn to RA")
		return nil, nil, ctx, err
	}

	orderClient := orderservice_ra.NewOrderServiceClient(conn)

	return orderClient, conn, ctx, nil
}

func (s *PaymentService) CreateOrderRA(status *paymentservice.PaymentResult, input domain.Order) error {
	orderClient, conn, ctx, err := s.CreateConnectionRA()
	if err != nil {
		return err
	}

	protoDate, err := ptypes.TimestampProto(input.RequiredTime)
	if err != nil {
		log.Error().Err(err).Msg("error occurred while updating RequiredTime")
		return err
	}

	protoDishes := make([]*orderservice_ra.Dish, 0, len(input.Dishes))
	for _, val := range input.Dishes {
		if val.Amount <= 0 {
			log.Error().Msg("invalid number of value")
			return errors.New("invalid number of value")
		}
		protoDishes = append(protoDishes, &orderservice_ra.Dish{ID: val.DishID, Amount: int64(val.Amount)})
	}

	customer, err := s.repo.GetCustomerByID(input.CustomerID)
	if err != nil {
		return err
	}

	order := &orderservice_ra.Order{
		OrderID:           status.IdOrder,
		RestaurantID:      input.RestaurantID,
		DeliveryTime:      protoDate,
		ClientFullName:    customer.FullName,
		ClientPhoneNumber: customer.PhoneNumber,
		Address:           input.Address,
		PaymentType:       status.PaymentType,
		List:              protoDishes,
	}

	_, err = orderClient.CreateOrder(ctx, order)
	if err != nil {
		log.Error().Err(err).Msg("error occurred while creating order in RA")

		return err
	}

	defer conn.Close()

	return nil
}
