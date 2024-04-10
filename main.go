package main

import (
	"context"
	"eshop/ordering/pb"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.UnimplementedOrderingServer
}

func (s *Server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	// place a new order.

	res := &pb.CreateOrderResponse{
		OrderId: "order123",
	}

	return res, nil
}

func main() {
	listner, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	pb.RegisterOrderingServer(srv, &Server{})

	reflection.Register(srv)

	fmt.Println("starting grpc server")
	if err := srv.Serve(listner); err != nil {
		panic(err)
	}

}
