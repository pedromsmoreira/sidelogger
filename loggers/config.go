package loggers

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Logger struct {
		Name    string `yaml:"name"`
		Shipper string `yaml:"shipper"`
	} `yaml:"logger"`
}

func ReadConfigFile(fpath string) (*Config, error) {
	f, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
