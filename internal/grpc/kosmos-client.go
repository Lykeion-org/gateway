package grpc

import (
	"log/slog"
	kosmos "github.com/Lykeion-org/go-shared/pkg/grpc/generated/kosmos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


type kosmosClient struct {
	Client kosmos.KosmosServiceClient
	connection *grpc.ClientConn
}

func NewKosmosClient(target string) *kosmosClient {
	conn, err := newClientConnection(target)
	if err != nil {
		slog.Error("Failed to start Kosmos service on target", "target", target, "error",err)
		panic(err)
	}

	client := kosmos.NewKosmosServiceClient(conn)

	return &kosmosClient{
		Client: client,
		connection: conn,
	}
}

func (c *kosmosClient) StopClient() error {
	c.connection.Close()
	c.connection = nil
	return nil
}

func newClientConnection(target string) (*grpc.ClientConn, error){
	creds := insecure.NewCredentials()
	conn, err:= grpc.NewClient(target, grpc.WithTransportCredentials(creds))
	if err != nil {
		slog.Debug("Failed to open grpc client connection", "target", target, "error",err)
	}
	return conn, nil
}