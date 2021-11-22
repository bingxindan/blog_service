package main

import (
	"blog_service/App/Router"
	"flag"
	"github.com/bingxindan/bxd_go_lib/tools/confutil"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"time"
)

type Server struct {
	server *http.Server
	opts   ServerOptions
	exit   chan os.Signal
}

func main() {
	flag.Parse()

	confutil.InitConfig()

	defer recovery()

	s := NewServer()
	engine := s.GetGinEngine()
	engine.Use(RequestHeader())
	Router.RouterRegister(engine)

	er := s.Serve()
	if er != nil {
		log.Printf("Server stop err:%v", er)
	} else {
		log.Printf("Server exit")
	}
}

//Serve serve http request
func (s *Server) Serve() error {
	err := s.server.ListenAndServe()
	return err
}

func (s *Server) GetGinEngine() *gin.Engine {
	return s.server.Handler.(*gin.Engine)
}

func NewServer() *Server {
	opts := DefaultOptions()
	s := new(Server)
	s.opts = opts
	handler := gin.New()
	server := &http.Server{
		Addr:         opts.Addr,
		Handler:      handler,
		ReadTimeout:  opts.ReadTimeout,
		WriteTimeout: opts.WriteTimeout,
		IdleTimeout:  opts.IdleTimeout,
	}
	s.exit = make(chan os.Signal, 2)
	s.server = server
	return s
}

func RequestHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		benchmark := c.GetHeader("Bxd-Request-Type")
		if benchmark == "performance-testing" {
			c.Set("IS_BENCHMARK", "1")
		}
	}
}

//ServerOptions http server options
type ServerOptions struct {
	// run mode 可选debug/release
	Mode string `ini:"mode"`
	// TCP address to listen on, ":http" if empty
	Addr string `ini:"addr"`
	//grace mode 可选graceful/oversea 为空不使用
	Grace bool `ini:"grace"`

	// ReadTimeout is the maximum duration for reading the entire
	// request, including the body.
	//
	// Because ReadTimeout does not let Handlers make per-request
	// decisions on each request body's acceptable deadline or
	// upload rate, most users will prefer to use
	// ReadHeaderTimeout. It is valid to use them both.
	ReadTimeout time.Duration `ini:"readTimeout"`
	// WriteTimeout is the maximum duration before timing out
	// writes of the response. It is reset whenever a new
	// request's header is read. Like ReadTimeout, it does not
	// let Handlers make decisions on a per-request basis.
	WriteTimeout time.Duration `ini:"writeTimeout"`
	// IdleTimeout is the maximum amount of time to wait for the
	// next request when keep-alives are enabled. If IdleTimeout
	// is zero, the value of ReadTimeout is used. If both are
	// zero, ReadHeaderTimeout is used.
	IdleTimeout time.Duration `ini:"idelTimeout"`
}

func DefaultOptions() ServerOptions {
	return ServerOptions{
		Addr:         ":10088",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  5 * time.Second,
	}
}

func recovery() {
	if rec := recover(); rec != nil {
		log.Printf("Panic occur")
		if err, ok := rec.(error); ok {
			log.Printf("PanicRecover Unhandled error: %v\n stack:%v", err.Error(), cast.ToString(debug.Stack()))
		} else {
			log.Printf("PanicRecover Panic: %v\n stack:%v", rec, cast.ToString(debug.Stack()))
		}
	}
}
