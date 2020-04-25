package conf

import (
	"errors"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/ishansd94/gateway-proxy/pkg/env"
	"github.com/ishansd94/gateway-proxy/pkg/log"
)

// Read the config from the configuration file and return a Config struct
func GetConfig() (*Config, error) {

	config := Config{}

	filepath := env.Get("PROXY_CONFIG_FILE", "/etc/gateway/gateway.yaml")

	log.Info("gateway", fmt.Sprintf("reading configuration file at: %s", filepath))

	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Error("gateway" , "error reading configuration file", err)
		return nil, err
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Error("gateway" , "invalid configuration file", err)
		return nil, err
	}

	// TODO: schema validation - use this pkg https://github.com/dealancer/validate
	if config.Servers == nil {
		err = errors.New("configuration file schema validation failed")
		log.Error("gateway" , "invalid configuration file", err)
		return nil, err
	}

	return &config, nil
}



