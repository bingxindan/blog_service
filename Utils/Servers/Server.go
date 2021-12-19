package Servers

import (
	"blog_service/Utils/Servers/Http"
	"context"
)

const (
	SrvModeHttpRpcPprof = iota
	SrvModeFastRpc
	SrvModeHttp
	SrvModeFast
	SrvModeRpc
)

type Server interface {
	Start(context.Context) error
	Stop(context.Context) error
}

func InitApp(srvMode int) (*App, error) {
	switch srvMode {
	case SrvModeHttpRpcPprof:
		httpServer := Http.NewHttpServer()
		a := new(App)
		app := a.New(
			Name(AppName),
			Version(AppVersion),
			Metadata(map[string]string{}),
			Serve(httpServer),
		)
		return app, nil
	case SrvModeFastRpc:
		break
	case SrvModeHttp:
		break
	case SrvModeRpc:
		break
	case SrvModeFast:
		break
	}
	return nil, nil

}
