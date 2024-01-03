package main

import (
	"context"
	"fmt"
	"go-microservice/greet/greetpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func (*server) Greet(ctx context.Context, request *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	firstName := request.GetGreeting().GetFirstName()
	lastName := request.GetGreeting().GetLastName()
	result := "Hello, " + firstName + " " + lastName
	response := &greetpb.GreetResponse{
		Result: result,
	}
	return response, nil
}

func main() {
	fmt.Println("GRPC server is starting ...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve : %v", err)
	}
}
