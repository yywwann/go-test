package grpc

import (
	"context"
	"github.com/rs/zerolog/log"
	"time"

	pb "example/load_balance_test/lb_demo1/generator/api"
	"example/load_balance_test/lb_demo1/generator/service"
	xgrpc "gitlab.33.cn/chat/dtalk/pkg/net/grpc"
	"google.golang.org/grpc"
)

func New(c *xgrpc.ServerConfig, svr *service.Service) *xgrpc.Server {
	connectionTimeout := grpc.ConnectionTimeout(time.Duration(c.Timeout))
	ws := xgrpc.NewServer(c, connectionTimeout)
	pb.RegisterGeneratorServer(ws.Server(), &server{svr})
	ws, err := ws.Start()
	if err != nil {
		panic(err)
	}
	return ws
}

type server struct {
	//pb.UnimplementedGeneratorServer
	svr *service.Service
}

type key = struct{}

func (s *server) GetID(ctx context.Context, req *pb.Empty) (*pb.GetIDReply, error) {
	log.Info().Int64("index", req.Index).Msg("")
	return &pb.GetIDReply{
		Id: s.svr.GetID(),
	}, nil
}
