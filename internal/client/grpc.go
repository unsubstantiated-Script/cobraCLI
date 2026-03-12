package client

import (
	"context"
	pb "goDistributedSystem/pkg/pb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClient struct {
	conn   *grpc.ClientConn
	client pb.NodeServiceClient
}

func NewGRPCClient(addr string) (*GRPCClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &GRPCClient{
		conn:   conn,
		client: pb.NewNodeServiceClient(conn),
	}, nil
}

func (c *GRPCClient) Close() {
	c.conn.Close()
}

func (c *GRPCClient) ReportStatus(ctx context.Context, name string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	resp, err := c.client.ReportStatus(ctx, &pb.Request{Data: name})
	if err != nil {
		return "", err
	}

	return resp.Data, nil
}
