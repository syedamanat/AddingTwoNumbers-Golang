package main

import (
	"context"
	"fmt"
	"log"
	"net"

	addpb "github.com/syedamanat/testproject/src"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Add(ctx context.Context, req *addpb.AddRequest) (*addpb.AddResponse, error) {
	fmt.Println("Add Function was invoked with %v ", req)

	num1 := req.GetAdd().GetNum1()
	num2 := req.GetAdd().GetNum2()

	result := num1 + num2

	res := &addpb.AddResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	fmt.Println("Server hosted on 50051")

	if err != nil {
		log.Fatalf("Error while opening a connection: %v", err)
	}

	s := grpc.NewServer()
	addpb.RegisterAddServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

}
