package server

import (
	"fmt"
	"github.com/yacen/gong/context"
	"net/http"
)

type serverHandler struct {
	middlewares []Middleware
}

func (h *serverHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if len(h.middlewares) == 0 {
		return
	} else {
		chain := &RealMiddlewareChain{middlewares: h.middlewares, index: 0}
		ctx := context.NewContext(res, req)
		chain.Next(ctx)
	}
}

type Middleware interface {
	Do(ctx *context.Context, chain Chain)
}

type Chain interface {
	// call next middleware
	Next(ctx *context.Context)
}

type RealMiddlewareChain struct {
	middlewares []Middleware
	index       int
}

func (c *RealMiddlewareChain) Next(ctx *context.Context) {
	if c.index >= len(c.middlewares) {
		return
	}
	fmt.Println(c.index)
	next := &RealMiddlewareChain{middlewares: c.middlewares, index: c.index + 1}
	middleware := c.middlewares[c.index]
	middleware.Do(ctx, next)
	fmt.Println(ctx)
}

func (c *RealMiddlewareChain) add(middleware Middleware) {
	c.middlewares = append(c.middlewares, middleware)
}

type MiddlewareFunc func(ctx *context.Context, chain Chain)

type FunctionMiddleware struct {
	Fn MiddlewareFunc
}

func (m *FunctionMiddleware) Do(ctx *context.Context, chain Chain) {
	m.Fn(ctx, chain)
}
