package context

import (
	"github.com/yacen/gong/request"
	"github.com/yacen/gong/response"
	"net/http"
)

type Context struct {
	Res http.ResponseWriter
	Req *http.Request
	Err error
	response.Response
	request.Request
}

func NewContext(res http.ResponseWriter, req *http.Request) *Context {
	return &Context{Res: res, Req: req}
}
