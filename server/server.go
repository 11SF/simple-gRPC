package main

import (
	"context"
	"fmt"
	"net"

	pb "simple-gRPC/model"

	"google.golang.org/grpc"
)

type CustomerServiceServer struct {
}

func (s *CustomerServiceServer) GetCustomerById(ctx context.Context, req *pb.ReqCustomerById) (*pb.ResCustomerById, error) {

	resp := pb.ResCustomerById{
		Results: []*pb.ResCustomerById_Result{
			{
				Id:       "1",
				Username: "11SF",
				Email:    "test1@gmail.com",
			},
			{
				Id:       "2",
				Username: "Folk",
				Email:    "test2@gmail.com",
			},
		},
	}

	return &resp, nil
}

func StartCustomerServer() {
	server := CustomerServiceServer{}

	lis, err := net.Listen("tcp", "localhost:9000")

	grpcServer := grpc.NewServer()
	pb.RegisterCustomerServiceServer(grpcServer, &server)

	// Start grpcServer
	if err = grpcServer.Serve(lis); err != nil {
		panic(err)
	}
	fmt.Println("server started")
}

func main() {
	StartCustomerServer()
}
