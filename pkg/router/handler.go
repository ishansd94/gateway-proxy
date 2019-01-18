package router

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

//Handle starts up a web server.
func (r *Router) Handle() error {

	addr := fmt.Sprintf(":%s", strconv.Itoa(r.Conf.Listener))

	log.Info("[Router-Handle]", "starting server..")

	http.HandleFunc("/", r.ProxyHandler)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Error("[Router-Handle]", "error starting server...", err.Error())
		return err

	}

	return nil
}

//ProxyHandler handle request transfer
func (r *Router) ProxyHandler(res http.ResponseWriter, req *http.Request) {
	log.Info("[Proxy]", "upstream : ", req.RequestURI)
	targetURL, err := r.GetTarget(req)
	if err != nil {
		log.Error("[Proxy]", "error getting upstream", err.Error())
	}
	log.Info("[Proxy]", "downstream : ", targetURL.String())

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	req.URL.Host = targetURL.Host
	req.URL.Scheme = targetURL.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = targetURL.Host
	log.Info("[Proxy]", "connecting downstream : ", targetURL.String())
	proxy.ServeHTTP(res, req)
}

//GetTarget returns the target for the url
func (r *Router) GetTarget(req *http.Request) (*url.URL, error) {

	uri := strings.Replace(req.RequestURI, "/", "", 1)

	var targetURL *url.URL
	var err error

	for _, route := range r.Conf.Routes {
		if r, found := route[uri]; found {
			targetURL, err = url.Parse(fmt.Sprintf("http://%s:%s%s", r.Target, strconv.Itoa(r.Port), r.Path))
			if err != nil {
				return nil, err
			}
			return targetURL, nil
		}
	}

	return url.Parse(r.Conf.Fallback)
}
