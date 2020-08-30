package server

import (
	"github.com/gin-gonic/gin"
)

type HttpServer struct {

}

func (h *HttpServer) Init() *gin.Engine {
	router := gin.Default()


	if mode := gin.Mode(); mode == gin.TestMode {
		router.LoadHTMLGlob("./../templates/*")
	} else {
		router.LoadHTMLGlob("templates/*")
	}
	router.Static("/statics", "./statics")
	router.StaticFile("/favicon.ico", "./favicon.ico")

	return router
}

//func (h *HttpServer) pprofRouter(r gin.IRouter) {
//	p := r.Group("/debug")
//	{
//		p.GET("/pprof/profile", pprof.Profile)
//		r.Get("/pprof/cmdline", pprof.Cmdline)
//		r.Get("/pprof/symbol", pprof.Symbol)
//		r.Get("/pprof/trace", pprof.Trace)
//		r.Get("/pprof/*", pprof.Index)
//	}
//}