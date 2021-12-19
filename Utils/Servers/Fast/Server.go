package Fast

import (
	"blog_service/Internal/Router"
	"context"
	"errors"
	"fmt"
	"github.com/bingxindan/bxd_go_lib/tools/confutil"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type Server struct {
	*fasthttp.Server
	lis      net.Listener
	once     sync.Once
	endpoint *url.URL
	err      error
	network  string
	address  string
	timeout  time.Duration
	router   *fasthttprouter.Router
}

type ServerOption func(*Server)

func NewFastServer() *Server {
	conf := confutil.GetConfStringMap("fast")
	var opts []ServerOption
	opts = append(opts, Network(conf["Network"]))
	opts = append(opts, Address(conf["Addr"]))
	srv := NewServer(opts...)
	return srv
}

func NewServer(opts ...ServerOption) *Server {
	srv := &Server{
		network: "tcp",
		address: ":0",
		timeout: 1 * time.Second,
	}
	for _, o := range opts {
		o(srv)
	}
	routerRegister := Router.RegisterFast()
	srv.router = routerRegister

	srv.Server = &fasthttp.Server{
		Handler: srv.router.Handler,
	}
	return srv
}

func Network(network string) ServerOption {
	return func(s *Server) { s.network = network }
}

func Address(address string) ServerOption {
	return func(s *Server) { s.address = address }
}

func Timeout(timeout time.Duration) ServerOption {
	return func(s *Server) { s.timeout = timeout }
}

func (s *Server) Start(ctx context.Context) error {
	if _, err := s.Endpoint(); err != nil {
		return err
	}
	fmt.Printf("[FAST] server listening on: %+v\n", s.address)
	err := s.Serve(s.lis)
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	fmt.Printf("[HTTP] server stopping")
	return s.Shutdown()
}
