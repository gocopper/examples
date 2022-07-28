//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gocopper/copper"
	"github.com/gocopper/copper/csql"
	"github.com/gocopper/examples/hackernews/migrations"
	"github.com/google/wire"

	_ "github.com/mattn/go-sqlite3"
)

func InitMigrator(*copper.App) (*csql.Migrator, error) {
	panic(
		wire.Build(WireModule),
	)
}

var WireModule = wire.NewSet(
	copper.WireModule,
	csql.WireModule,

	wire.Value(csql.Migrations(migrations.Migrations)),
)
