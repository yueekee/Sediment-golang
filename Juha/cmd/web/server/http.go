package server

import (
	"github.com/gin-gonic/gin"
	"net/http/pprof"
)

type HttpServer struct {

}

func (h *HttpServer) Init() error {

	return nil
}

func (h *HttpServer) pprofRouter(r gin.IRouter) {
	p := r.Group("/debug")
	{
		p.GET("/pprof/profile", pprof.Profile)
		r.Get("/pprof/cmdline", pprof.Cmdline)
		r.Get("/pprof/symbol", pprof.Symbol)
		r.Get("/pprof/trace", pprof.Trace)
		r.Get("/pprof/*", pprof.Index)
	}
}