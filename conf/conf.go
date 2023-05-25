package conf

import "github.com/jasontconnell/conf"

type Config struct {
	// config values
}

func LoadConfig(filename string) (Config, error) {
	cfg := Config{}
	err := conf.LoadConfig(filename, &cfg)
	return cfg, err
}
