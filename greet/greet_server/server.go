package main

import (
	"../greetpb"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {}

func main() {
	fmt.Println("Greet Server!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
