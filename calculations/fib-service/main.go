package main

import (
	pb "calculations/pb/fib/v1"
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

var cache = map[uint64]uint64{}

func main() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Failed to listen:  %v", err)
	}

	fmt.Println("Fibnoacci service is running on :3000")

	s := grpc.NewServer()
	pb.RegisterFibServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) Compute(cxt context.Context, r *pb.FibRequest) (*pb.FibResponse, error) {
	result := &pb.FibResponse{}
	result.Result = Fibonacci(r.FibNum)
	logMessage := fmt.Sprintf("Request: %d, Result: %d", r.FibNum, result.Result)
	log.Println(logMessage)
	return result, nil
}

func Fibonacci(n uint64) uint64 {
	var a uint64 = 0
	var b uint64 = 1
	for i := 0; a <= n; i++ {
		temp := a
		a = b
		b = temp + a
		fmt.Println(a)
	}
	return a
}
