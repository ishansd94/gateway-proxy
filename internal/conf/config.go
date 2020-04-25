package conf

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/ishansd94/gateway-proxy/pkg/log"
)

// Config wraps all the configurations available for the loadbalancer
type Config struct {
	Logs Logs
	Servers []Server
}

// Logging configuration
type Logs struct {
	Mode string
}

// Loadbalancer server configuration
type Server struct {
	Port int
	Timeout int
	Mode string

	Backends []Backend
}

// Backend server configuration
type Backend struct {
	Match string
	Target string
	Port int
	Path string
	Scheme string
}


func (b *Backend) GetTarget() (*url.URL, error) {

	tagrgetfullpath := b.Target

	if b.Path != "" {
		tagrgetfullpath = fmt.Sprintf("%s/%s", b.Target, b.Path)
	}

	constructedURL, err := url.Parse(fmt.Sprintf("%s://%s:%s" , b.Scheme, tagrgetfullpath, strconv.Itoa(b.Port) ))
	if err != nil {
		log.Error("gateway", "error constructing target url", err)
		return nil, err
	}

	return constructedURL, nil
}

func (s *Server) GetRouteMap() map[string]*url.URL {

	rmap := map[string]*url.URL{}

	for _, backend := range s.Backends {

		target, err := backend.GetTarget()
		if err != nil {
			log.Warn("gateway", fmt.Sprintf("error creating target entry: %s", err.Error()))
			target = nil
		}

		rmap[backend.Match] = target
	}

	return rmap
}