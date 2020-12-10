package config

import "github.com/BurntSushi/toml"

var config *Config

type Config struct {
	Server Server
	Mysql Mysql
}

type Server struct {
	Port int64
}

type Mysql struct {
	Host,User,Pass,DataName string
	Port int64

}

//读取配置文件

func ReadConf() *Config {

	if config != nil {
		return config
	}
	var c *Config

	if _,err := toml.DecodeFile("./config/config.toml",&c); err != nil {
		panic(err)
	}

	return config
}