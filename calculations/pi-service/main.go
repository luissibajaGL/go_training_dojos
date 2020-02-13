package main

import (
	pb "calculations/pb/pi/v1"
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Failed to listen:  %v", err)
	}

	fmt.Println("Pi service is running on :3000")

	s := grpc.NewServer()
	pb.RegisterPiServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) Compute(cxt context.Context, r *pb.PiRequest) (*pb.PiResponse, error) {
	result := &pb.PiResponse{}
	result.Result = piNum()
	logMessage, _ := fmt.Println("Result: ", result.Result)
	log.Println(logMessage)
	return result, nil
}

func piNum() float64 {
	return 3.141592653589793238
}
