package main

import (
	"context"
	"log"
	"net"
	"time"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	log.Printf("Received: %v\n", in.GetName())
	// time.Sleep(5 * time.Second)
	// log.Println("Responding")
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
	time.Sleep(time.Second)

	go func() {
		conn, err := grpc.Dial("localhost:6666", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatal(err)
		}
		client := NewGreeterClient(conn)
		resp, err := client.SayHello(context.Background(), &HelloRequest{
			Name: "Aurelien DEROIDE",
		})
		_ = resp
		if err != nil {
			log.Fatal(err)
		}
	}()
	time.Sleep(2 * time.Second)

	log.Println("Start stop")
	s.GracefulStop()
	log.Println("Stopped")
	time.Sleep(5 * time.Second)

}
