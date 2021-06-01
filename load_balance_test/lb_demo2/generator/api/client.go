package api

//
//import (
//	"context"
//	"time"
//
//	xgrpc "gitlab.33.cn/chat/dtalk/pkg/net/grpc"
//)
//
//// AppID unique app id for service discovery
//const AppID = "identify.service.generator"
//
//type Client struct {
//	client GeneratorClient
//}
//
//func New(addr string, timeout time.Duration) *Client {
//	conn, err := xgrpc.NewGRPCConn(addr, timeout)
//	if err != nil {
//		panic(err)
//	}
//	return &Client{
//		client: NewGeneratorClient(conn),
//	}
//}
//
//func (c *Client) GetID() (int64, error) {
//	reply, err := c.client.GetID(context.Background(), &Empty{})
//	if reply == nil {
//		return 0, err
//	}
//	return reply.Id, err
//}
