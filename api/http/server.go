// Package http
// @Description:

package http

import (
	"context"
	"github.com/frank-cloud-tech/class-oa/api/http/handler/probe"
	"github.com/frank-cloud-tech/class-oa/internal"
	"net/http"
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	server *http.Server
}

func NewHTTPServer() (*HTTPServer, error) {
	var s HTTPServer
	e := gin.New()
	//中间件
	//e.Use(middleware.GinAccessLogger(), middleware.MetricsMiddleware(),
	//	middleware.ErrorHandleMiddleWare(), middleware.PanicHandleMiddleWare)
	if err := internal.Init(); err != nil {
		return nil, err
	}
	router := e.Group("/api/v1")
	//中间件
	//router.Use(middleware.LoginRequired(), middleware.TreeRequired(utils.HeaderTreeIdGetter))

	// controller层对象初始化, 路由注册
	probe.Register(router)
	probe.Register(router.Group("metrics"))

	s.server = &http.Server{Addr: "0.0.0.0:8080", Handler: e}
	return &s, nil
}

func (s *HTTPServer) Run() error {
	return s.server.ListenAndServe()
}

func (s *HTTPServer) Close() error {
	return s.server.Shutdown(context.Background())
}
