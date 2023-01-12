package main

import (
	"context"
	"log"
	"time"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:6666", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	client := NewGreeterClient(conn)
	log.Println("Starting RPC")
	resp, err := client.SayHello(context.Background(), &HelloRequest{
		Name: "Aurelien",
	})
	if err != nil {
		log.Fatal("RPC FAILED: ", err)
	}
	log.Println("Finished RPC: ", resp)

	log.Println("Waiting 30 sec to not not close tcp connection manually too fast")
	time.Sleep(30 * time.Second)
	log.Println("Shutting down the client connection")
	err = conn.Close()
	if err != nil {
		log.Fatal(err)
	}
}
