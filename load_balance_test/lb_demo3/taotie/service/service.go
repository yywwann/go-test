package service

import (
	"context"
	"example/load_balance_test/lb_demo3/common"
	idgen "example/load_balance_test/lb_demo3/generator/api"
	"example/load_balance_test/lb_demo3/naming"
	"example/load_balance_test/lb_demo3/taotie/config"
	"fmt"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
	"google.golang.org/grpc/resolver"
	"os"
	"time"
)

type Service struct {
	log            zerolog.Logger
	cfg            *config.Config
	idGenRPCClient idgen.GeneratorClient
}

var srvName = "taotie/srv"

func New(cfg *config.Config) *Service {
	s := &Service{
		log:            zlog.Logger.With().Str("service", srvName).Logger(),
		cfg:            cfg,
		idGenRPCClient: newIdGenClient(cfg),
	}

	//log init
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if config.Conf.Env == "debug" {
		s.log = zlog.Output(zerolog.ConsoleWriter{Out: os.Stdout}).With().Str("service", srvName).Logger()
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else if config.Conf.Env == "benchmark" {
		zerolog.SetGlobalLevel(zerolog.Disabled)
	}

	return s
}

func newIdGenClient(cfg *config.Config) idgen.GeneratorClient {
	rb := naming.NewResolver(cfg.Reg.RegAddrs, cfg.IdGenRPCClient.Schema)
	resolver.Register(rb)

	addr := fmt.Sprintf("%s:///%s", cfg.IdGenRPCClient.Schema, cfg.IdGenRPCClient.SrvName) // "schema://[authority]/service"
	fmt.Println("generator rpc client call addr:", addr)

	conn, err := common.NewGRPCConn(addr, time.Duration(cfg.IdGenRPCClient.Dial))
	if err != nil {
		panic(err)
	}
	return idgen.NewGeneratorClient(conn)
}

//GetLogId 由 generator 服务生成唯一 id
func (s *Service) GetLogId(ctx context.Context, index int64) (id int64, err error) {
	var (
		req   idgen.Empty
		reply *idgen.GetIDReply
	)
	req.Index = index
	reply, err = s.idGenRPCClient.GetID(ctx, &req)
	if err != nil {
		return 0, errors.WithMessagef(err, "GetLogId")
	}

	return reply.Id, nil
}
