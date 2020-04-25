package router

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/ishansd94/gateway-proxy/pkg/log"
)

type Handler struct {
	Name             string
	Server           *http.Server
	InterruptChannel chan struct{}
}

var servers sync.WaitGroup

func NewRouter(name string, httpServer *http.Server) *Handler {

	return &Handler{
		Name:             name,
		Server:           httpServer,
		InterruptChannel: make(chan struct{}),
	}
}

func (h *Handler) Start() {

	servers.Add(1)

	go h.gracefulShutdown()

	log.Info(fmt.Sprintf("router: %s", h.Name), fmt.Sprintf("listening on %s", h.Server.Addr))

	go func() {
		if err := h.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error(fmt.Sprintf("router: %s", h.Name), "error starting server...", err)
		}
	}()
}

func (h *Handler) gracefulShutdown() {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGKILL, syscall.SIGTERM)

	ctx := context.Background()

	log.Warn(fmt.Sprintf("router: %s", h.Name), fmt.Sprintf("server shutting down.. got %s", <-interrupt))
	if err := h.Server.Shutdown(ctx); err != nil {
		log.Error(fmt.Sprintf("router: %s", h.Name), "error shutting down server...", err)
	}
	servers.Done()
}

func Wait() {
	servers.Wait()
}

func (h *Handler) GetRouter() *http.Server {
	return h.Server
}

func (h *Handler) GetRouterHandler() *http.Handler {
	return &h.GetRouter().Handler
}