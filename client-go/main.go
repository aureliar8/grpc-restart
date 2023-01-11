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
	log.Println("Saying hello")
	resp, err := client.SayHello(context.Background(), &HelloRequest{
		Name: "Aurelien DEROIDE",
	})
	if err != nil {
		log.Fatal("RPC FAILED: ", err)
	}
	log.Println("Got an hello response:", resp)

	log.Println("Waiting 30 sec to not force the tcp conn drop")
	time.Sleep(30 * time.Second)

	err = conn.Close()
	if err != nil {
		log.Fatal(err)
	}
}
