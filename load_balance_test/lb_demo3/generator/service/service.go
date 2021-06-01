package service

import (
	"example/load_balance_test/lb_demo3/generator/config"
	"gitlab.33.cn/chat/dtalk/pkg/util"
)

type Service struct {
	idGenerator *util.Snowflake
}

func New(c *config.Config) *Service {
	g, err := util.NewSnowflake(c.Node)
	if err != nil {
		panic(err)
	}
	s := &Service{
		idGenerator: g,
	}
	return s
}

func (s *Service) GetID() int64 {
	return s.idGenerator.NextId()
}
