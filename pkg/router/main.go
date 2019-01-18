package router

const DefaultConfigFile FilePath = "/var/run/proxy/conf.yaml"

type FilePath string

// Route specific config
type Route struct {
	Target string `yaml:"target"`
	Port   int    `yaml:"port"`
	Path   string `yaml:"path"`
}

// RouteMap is an alias for ma[string]Route
type RouteMap map[string]Route

// Proxy holds the YAML spec
type Proxy struct {
	Fallback string     `yaml:"fallback"`
	Listener int        `yaml:"listener"`
	Verbose  bool       `yaml:"verbose"`
	Routes   []RouteMap `yaml:"routes"`
}

type Router struct {
	Conf Proxy
}

func NewClient(proxy Proxy) *Router {
	return &Router{proxy}
}

func (r Router) Serve() error {

	if err := r.Handle(); err != nil {
		return err
	}

	return nil
}
