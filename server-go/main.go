package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	grpc "google.golang.org/grpc"
)

type server struct {
	UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	log.Printf("Starting to process request")
	// Simulate long request processing
	time.Sleep(5 * time.Second)
	log.Println("Responding to request")
	return &HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":6666")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	RegisterGreeterServer(s, &server{})

	go func() {
		log.Println("Listening")
		if err := s.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	time.Sleep(500 * time.Millisecond)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Will stop on SIGINT/SIGTERM")
	<-sigs
	log.Println("Start gracefull stop")
	s.GracefulStop()
	log.Println("Fully Stopped")
}
