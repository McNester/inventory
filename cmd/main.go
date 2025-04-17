package main

import (
	"fmt"
	"inventory/db"
	"inventory/handlers"
	"net"

	pb "cloud_commons/inventory"

	"google.golang.org/grpc"
)

var (
	product_handler *handlers.ProductHandler
)

func main() {

	defer db.CloseConnection()

	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic("Failed to listen tcp on 50051")
	}

	server := grpc.NewServer()
	product_handler = handlers.NewProductHandler()

	pb.RegisterProductServiceServer(server, product_handler)

	fmt.Println("Starting inventory grpc on 50051")

	err = server.Serve(lis)

	if err != nil {
		panic("Failed to serve the inventory: " + err.Error())
	}
}
