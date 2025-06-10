package main

import (
	"log"

	api "github.com/Lykeion/gateway/internal/api"
	pb "github.com/Lykeion/gateway/internal/grpc/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {



	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials())); if err != nil {
		log.Fatal("Couldn't connect to grpc service: %v", err)
	}
	defer conn.Close()

	client := pb.NewLanguageServiceClient(conn)

	a := api.NewApi(client)
	a.InitializeApi()
}