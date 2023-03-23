package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"tinderData/internal/controller"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Use(ghttp.MiddlewareHandlerResponse)
			s.Group("records/v1/zmq", func(group *ghttp.RouterGroup) {
				group.Bind(
					controller.Zmq,
				)
			})

			s.Group("/info", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.GET("/version", func(r *ghttp.Request) {
					r.Response.Writefln("v1.0_20221223")
				})

			})

			s.Group("/iotdb", func(group *ghttp.RouterGroup) {
				group.Bind(
					controller.IotDb,
				)
			})

			s.Run()
			return nil
		},
	}
)
