package config

import yaml "gopkg.in/yaml.v2"

type server struct {
	Port int
}

type bookservice struct {
	Host string
}

type ConfigType struct {
	Server      server
	Bookservice bookservice
}

var (
	Config  ConfigType
	Version string
)

func LoadConfig(in []byte) {

	err := yaml.Unmarshal(in, &Config)
	if err != nil {
		panic(err)
	}
}
