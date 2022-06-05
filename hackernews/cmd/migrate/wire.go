//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gocopper/examples/hackernews/pkg/app"

	"github.com/gocopper/copper"
	"github.com/gocopper/copper/csql"
	"github.com/google/wire"
)

func InitMigrator(*copper.App) (*csql.Migrator, error) {
	panic(
		wire.Build(WireModule),
	)
}

var WireModule = wire.NewSet(
	copper.WireModule,
	csql.WireModule,

	app.WireModule,

	wire.Struct(new(app.ProvideMigrationsParams), "*"),
	app.ProvideMigrations,
)
