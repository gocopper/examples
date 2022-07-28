//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gocopper/copper/csql"
	"github.com/gocopper/examples/hackernews/pkg/app"
	"github.com/gocopper/examples/hackernews/web"
	"github.com/gocopper/examples/hackernews/web/build"

	"github.com/gocopper/copper"
	"github.com/gocopper/copper/chttp"
	"github.com/google/wire"

	_ "github.com/mattn/go-sqlite3"
)

func InitServer(*copper.App) (*chttp.Server, error) {
	panic(
		wire.Build(WireModule),
	)
}

var WireModule = wire.NewSet(
	copper.WireModule,
	chttp.WireModule,
	wire.Struct(new(app.NewHTTPHandlerParams), "*"),
	app.NewHTTPHandler,
	app.WireModule,
	app.NewRouter,
	wire.Struct(new(app.NewRouterParams), "*"),

	wire.InterfaceValue(new(chttp.HTMLDir), web.HTMLDir),
	wire.InterfaceValue(new(chttp.StaticDir), build.StaticDir),
	web.HTMLRenderFuncs,
	csql.WireModule,
)
