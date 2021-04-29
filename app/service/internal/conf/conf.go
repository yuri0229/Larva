package conf

import "github.com/BurntSushi/toml"

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
}

var (
	Conf     = &Config{}
)

func Init() (err error) {
	_, err = toml.DecodeFile("conf_service.toml", &Conf)
	return
}
