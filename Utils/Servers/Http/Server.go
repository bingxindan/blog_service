package Http

import (
	"blog_service/Internal/Router"
	"context"
	"errors"
	"fmt"
	"github.com/bingxindan/bxd_go_lib/tools/confutil"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type Server struct {
	*http.Server
	lis      net.Listener
	once     sync.Once
	endpoint *url.URL
	err      error
	network  string
	address  string
	timeout  time.Duration
	router   *gin.Engine
}

type ServerOption func(*Server)

func NewHttpServer() *Server {
	conf := confutil.GetConfStringMap("http")
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
	r := gin.New()
	// TODO-增加日志和错误处理
	r.Use()
	s := Router.RouterRegister(r)
	srv.router = s
	srv.Server = &http.Server{
		Handler: srv.router,
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
	s.BaseContext = func(net.Listener) context.Context {
		return ctx
	}
	fmt.Printf("[HTTP] server listening on: %+v\n", s.address)
	err := s.Serve(s.lis)
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	fmt.Printf("[HTTP] server stopping")
	return s.Shutdown(ctx)
}
