package main

import (
	"context"
	"fmt"
	"log"

	addpb "github.com/syedamanat/testproject/src"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Error while opening a connection: %v", err)
	}

	defer cc.Close()

	c := addpb.NewAddServiceClient(cc)
	doUnary(c)
}

func doUnary(c addpb.AddServiceClient) {
	fmt.Println("Initiating Unary gRPC Operation")

	req := &addpb.AddRequest{
		Add: &addpb.Add{
			Num1: 2,
			Num2: 4,
		},
	}

	res, err := c.Add(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling Add: %v", err)
	}
	log.Printf("Response from Add: %v", res)
}
