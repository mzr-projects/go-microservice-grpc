package main

import (
	"context"
	"fmt"
	"go-microservice/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	fmt.Println("Client is running ...")

	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Could not connect : %v", err)
	}

	defer func(cc *grpc.ClientConn) {
		err := cc.Close()
		if err != nil {

		}
	}(cc)

	c := greetpb.NewGreetServiceClient(cc)
	//fmt.Printf("Created client : %s", c)

	doUnary(err, c)
}

func doUnary(err error, c greetpb.GreetServiceClient) {

	request := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Maziar",
			LastName:  "TghPr",
		},
	}

	response, err := c.Greet(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling rpc %v", err)
	}

	log.Printf("Response from greet service: %v", response.Result)
}
