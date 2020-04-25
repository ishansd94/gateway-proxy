package gateway

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sanity-io/litter"

	"github.com/ishansd94/gateway-proxy/internal/conf"
	"github.com/ishansd94/gateway-proxy/pkg/router"
)

type Handler struct {
	RouteMap map[string]*url.URL
}

func Run() error {
	config, err := conf.GetConfig()
	if err != nil {
		return err
	}

	for _, s := range config.Servers {

		port := fmt.Sprintf(":%s", strconv.Itoa(s.Port))

		r := gin.Default()

		rmap := s.GetRouteMap()

		for _, backend := range s.Backends {

			target := rmap[backend.Match]

			r.Any(backend.Match, gin.WrapF(NewProxyHandler(target)))
		}

		serverConfig := &http.Server{
			Addr:         port,
			Handler:      r,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		}

		server := router.NewRouter("app-server", serverConfig)

		server.Start()

	}

	router.Wait()

	return nil
}

func NewProxyHandler(target *url.URL) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		proxy := httputil.NewSingleHostReverseProxy(target)

		litter.Dump(r)

		r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))

		r.URL.Host = target.Host
		r.URL.Scheme = target.Scheme
		r.Host = target.Host

		proxy.ServeHTTP(w, r)
	}
}
