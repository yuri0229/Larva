package conf

import (
	"gf/pkg/log"
	"github.com/BurntSushi/toml"
)

type Config struct {
	Rpc struct{
		Addr 	string
	}
	Http struct{
		Addr 	string
	}
	Db struct{
		Dsn 	string
	}
	Log *log.Config
}

var (
	Conf     = &Config{}
)

func Init() (err error) {
	_, err = toml.DecodeFile("conf_service.toml", &Conf)
	return
}
