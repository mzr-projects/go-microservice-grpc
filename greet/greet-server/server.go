package main

import (
	"context"
	_ "context"
	"fmt"
	"go-microservice/greet/greetpb"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"time"
)

type server struct{}

func (*server) Greet(ctx context.Context, request *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {

	fmt.Printf("Unary Greet function was invoked with %v\n", request)

	firstName := request.GetGreeting().GetFirstName()
	lastName := request.GetGreeting().GetLastName()
	result := "Hello, " + firstName + " " + lastName

	response := &greetpb.GreetResponse{
		Result: result,
	}
	return response, nil
}

func (*server) GreetManyTimes(request *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {

	fmt.Printf("Server Stream Greet function was invoked with %v\n", request)

	firstName := request.Greeting.GetFirstName()
	lastName := request.Greeting.GetLastName()

	for i := 0; i < 10; i++ {
		result := "Hello " + firstName + " " + lastName + ", number: " + strconv.Itoa(i)
		response := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		err := stream.Send(response)
		if err != nil {
			fmt.Printf("Sending as stream has error: %v", err)
		}
		time.Sleep(1000 * time.Millisecond)
	}

	return nil
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
