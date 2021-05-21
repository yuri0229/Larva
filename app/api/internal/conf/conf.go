package conf

import "github.com/BurntSushi/toml"

type Config struct {
	Rpc struct{
		Addr 	string
	}
	Http struct{
		Addr 	string
		Secret 	string
	}
	Db struct{
		Dsn 	string
	}
}

var (
	Conf     = &Config{}
)

func Init() (err error) {
	_, err = toml.DecodeFile("conf_api.toml", &Conf)
	return
}
