package config

import (
	"example/toml_test/pkg/config"
	"flag"
	"github.com/rs/zerolog/log"
)

var (
	confPath string

	Conf *Config
)

func init() {
	flag.StringVar(&confPath, "conf", "test.toml", "default config path.")
}

// Init init config.
func Init() (err error) {
	Conf = Default()
	log.Debug().Str("confPath", confPath).Msg("")
	return config.LoadConfig(confPath, &Conf, config.UseEnv())
}

func Default() *Config {
	return &Config{
		MySQL: &MySQL{
			Host: "",
			Port: 0,
			User: "",
			Pwd:  "",
			Db:   "",
		},
	}
}

type MySQL struct {
	Host string
	Port int32
	User string
	Pwd  string
	Db   string
}

type Config struct {
	MySQL *MySQL
}
