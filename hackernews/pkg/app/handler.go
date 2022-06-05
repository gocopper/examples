package app

import (
	"net/http"

	"github.com/gocopper/copper/chttp"
	"github.com/gocopper/copper/clogger"
)

type NewHTTPHandlerParams struct {
	HTML *chttp.HTMLRouter
	App  *Router

	RequestLoggerMW *chttp.RequestLoggerMiddleware
	Logger          clogger.Logger
}

func NewHTTPHandler(p NewHTTPHandlerParams) http.Handler {
	return chttp.NewHandler(chttp.NewHandlerParams{
		GlobalMiddlewares: []chttp.Middleware{
			p.RequestLoggerMW,
		},

		Routers: []chttp.Router{
			p.HTML,

			p.App,
		},

		Logger: p.Logger,
	})
}
