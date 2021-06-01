package config

import (
	"flag"
	"github.com/BurntSushi/toml"
	xtime "gitlab.33.cn/chat/dtalk/pkg/time"
	"time"
)

var (
	confPath string

	Conf *Config
)

func init() {
	flag.StringVar(&confPath, "conf", "taotie.toml", "default config path.")
}

// Init init config.
func Init() (err error) {
	Conf = Default()
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}

func Default() *Config {
	return &Config{
		IdGenRPCClient: &RPCClient{
			Schema:  "dtalk",
			SrvName: "generator",
			Dial:    xtime.Duration(time.Second),
			Timeout: xtime.Duration(time.Second),
		},
	}
}

type Config struct {
	Env            string
	AppId          string
	IdGenRPCClient *RPCClient
	Reg            *Reg
}

// RPCClient is RPC client config.
type RPCClient struct {
	Schema  string
	SrvName string // call
	Dial    xtime.Duration
	Timeout xtime.Duration
}

// Reg is service register/discovery config
type Reg struct {
	Schema   string
	SrvName  string // call
	RegAddrs string // etcd addrs, seperate by ','
}
