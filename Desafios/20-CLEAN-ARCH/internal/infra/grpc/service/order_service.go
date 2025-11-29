package service

import (
	"context"

	"github.com/Frank-Macedo/20-cleanArch/internal/infra/grpc/pb"
	"github.com/Frank-Macedo/20-cleanArch/internal/usecase"
	"google.golang.org/protobuf/types/known/emptypb"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) GetOrders(ctx context.Context, _ *emptypb.Empty) (*pb.OrderList, error) {
	getOrdersUseCase := usecase.NewGetOrdersUseCase(s.CreateOrderUseCase.OrderRepository)
	orders, err := getOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}
	var pbOrders []*pb.Order
	for _, order := range orders {
		pbOrder := &pb.Order{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		}
		pbOrders = append(pbOrders, pbOrder)
	}
	return &pb.OrderList{
		Orders: pbOrders,
	}, nil
}
