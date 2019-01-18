package app

import (
	"io/ioutil"
	"os"

	"github.com/ishansd94/go-reverse-proxy/pkg/router"
	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

//Bootstrap loads the configs and starts a router client
func Bootstrap() error {

	var proxyConf router.Proxy

	log.Info("[Bootstrap]", "starting..")

	f := fileLoc()

	log.Info("[Bootstrap]", "read config file at : ", f)

	configFile, err := ioutil.ReadFile(os.Getenv("CONFIG_FILE"))
	if err != nil {
		return err
	}

	log.Info("[Bootstrap]", "loading configuration..")

	err = yaml.Unmarshal(configFile, &proxyConf)
	if err != nil {
		return err
	}

	log.Info("[Bootstrap]", "configuration loaded..")

	log.Info("[Bootstrap]", "starting router..")
	client := router.NewClient(proxyConf)

	log.Info("[Bootstrap]", "router configured to listen on ..: ", proxyConf.Listener)
	if err := client.Serve(); err != nil {
		log.Error("[Bootstrap]", "router listen failed..: ", proxyConf.Listener)
		return err
	}

	return nil
}

func fileLoc() router.FilePath {
	fileLocation := router.DefaultConfigFile

	if f := os.Getenv("CONFIG_FILE"); f != "" {
		fileLocation = router.FilePath(f)
	}

	return fileLocation
}
