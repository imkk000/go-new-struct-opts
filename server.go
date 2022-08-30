package server

import (
	"net/http"
)

type Server struct {
	cfg Config
	sv  *http.Server
}

func NewServer(opts ...OptionFunc) *Server {
	return &Server{
		cfg: NewConfig(opts...),
	}
}

func NewConfig(opts ...OptionFunc) Config {
	cfg := defaultConfig
	if opts != nil {
		for _, opt := range opts {
			opt(&cfg)
		}
	}
	return cfg
}
