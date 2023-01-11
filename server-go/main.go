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
	"google.golang.org/grpc/keepalive"
)

type server struct {
	UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	log.Printf("Received: %v\n", in.GetName())
	// Remove comment for long request processing
	time.Sleep(3 * time.Second)
	log.Println("Responding")
	return &HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":6666")
	if err != nil {
		log.Fatal(err)
	}

	// https://lukexng.medium.com/grpc-keepalive-maxconnectionage-maxconnectionagegrace-6352909c57b8

	enforcement := keepalive.EnforcementPolicy{
		MinTime:             5 * time.Second,
		PermitWithoutStream: true,
	}

	s := grpc.NewServer(
		grpc.KeepaliveEnforcementPolicy(enforcement),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionAge:      1 * time.Second,
			MaxConnectionAgeGrace: 3 * time.Second,
		}),
	)
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
