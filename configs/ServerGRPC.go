package configs

import (
	"context"
	"fmt"
	"mega_api/pb"
)

type ServerGRPC struct {
	pb.UnimplementedOrderQueueServer
	orders []*pb.OrderRequest
}

func (s *ServerGRPC) CreateOrder(ctx context.Context, in *pb.OrderRequest) (*pb.OrderResponse, error) {
	s.orders = append(s.orders, in)
	fmt.Printf("Novo Pedido: %v\n", in)
	return &pb.OrderResponse{
		Customer: in.GetCustomer(),
		OrderId:  in.GetOrderId(),
		Title:    in.GetTitle(),
		Price:    in.GetPrice(),
	}, nil
}

func (s *ServerGRPC) GetOrders(ctx context.Context, in *pb.Empty) (*pb.OrderResponse, error) {
	if len(s.orders) == 0 {
		return &pb.OrderResponse{}, fmt.Errorf("Nenhum Pedido Por Aqui")
	}
	order := s.orders[0]
	s.orders = s.orders[1:]
	return &pb.OrderResponse{
		Customer: order.GetCustomer(),
		OrderId:  order.GetOrderId(),
		Title:    order.GetTitle(),
		Price:    order.GetPrice(),
	}, nil
}
