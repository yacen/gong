package server

import (
	"github.com/yacen/gong/context"
	"net/http"
)

type Server struct {
	server      *http.Server
	middlewares []Middleware
}

type MFun func(ctx *context.Context, next MFun)

func (s *Server) Use(f MiddlewareFunc) *Server {
	s.middlewares = append(s.middlewares, &FunctionMiddleware{Fn: f})
	return s
}

func (s *Server) Listen() error {
	s.server.Handler = &serverHandler{middlewares: s.middlewares}
	return s.server.ListenAndServe()
}
