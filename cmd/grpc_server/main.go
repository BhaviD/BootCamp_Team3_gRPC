package main

import (
	"fmt"
	"github.com/BhaviD/BootCamp_Team3_gRPC/pkg/dynamoDB"
	"github.com/BhaviD/BootCamp_Team3_gRPC/pkg/services/grpcPb"
	"github.com/BhaviD/BootCamp_Team3_gRPC/pkg/services/grpc_server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main()  {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Sorry failed to load server %v: ", err)
	}

	s := grpc.NewServer()

	grpcPb.RegisterGRPCServiceServer(s, &grpc_server.GrpcServer{})

	dynamoDB.AddCustomerTable()
	dynamoDB.LoadCustomerData()
	dynamoDB.AddRestaurantTable()
	dynamoDB.LoadRestaurantData()
	dynamoDB.AddOrderTable()
	dynamoDB.LoadOrderData()

	fmt.Println("Orders Server starting...")
	if s.Serve(lis); err != nil {
		log.Fatalf("failed to Serve %v", err)
	}
}

