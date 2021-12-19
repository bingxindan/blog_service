package Servers

import (
	"context"
	"net/url"
	"os"
	"time"
)

type options struct {
	id        string
	name      string
	version   string
	metadata  map[string]string
	endpoints []*url.URL

	ctx  context.Context
	sigs []os.Signal

	registerTimeout time.Duration
	stopTimeout     time.Duration
	servers         []Server
}

type Option func(o *options)

func ID(id string) Option {
	return func(o *options) {
		o.id = id
	}
}

func Name(name string) Option {
	return func(o *options) {
		o.name = name
	}
}

func Version(version string) Option {
	return func(o *options) {
		o.version = version
	}
}

func Metadata(md map[string]string) Option {
	return func(o *options) {
		o.metadata = md
	}
}

func Serve(srv ...Server) Option {
	return func(o *options) {
		o.servers = srv
	}
}
