package main

import (
	"context"
	"fmt"
	"time"

	pb "simple-gRPC/model"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ReqCustomerById(client pb.CustomerServiceClient) (*pb.ResCustomerById, error) {
	// Timeout 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ping := pb.ReqCustomerById{
		SearchId: 1,
	}
	res, err := client.GetCustomerById(ctx, &ping)
	statusCode := status.Code(err)
	if statusCode != codes.OK {
		return nil, err
	}
	// fmt.Printf("Customer info: %d, statusCode: %s\n", res.Results[0], statusCode.String())
	return res, err
}

func StartCustomerClient() {
	// Disable transport security for this example only,
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial("127.0.0.1:9000", opts...)
	if err != nil {
		panic(err)
	}

	client := pb.NewCustomerServiceClient(conn)

	res, err := ReqCustomerById(client)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Results)
}

func main() {
	StartCustomerClient()
}
