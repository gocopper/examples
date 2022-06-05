package app

import (
	"github.com/gocopper/copper/csql"
	"github.com/gocopper/examples/hackernews/pkg/posts"
)

type ProvideMigrationsParams struct {
	Posts *posts.Migration
}

func ProvideMigrations(p ProvideMigrationsParams) []csql.Migration {
	return []csql.Migration{
		p.Posts,
	}
}
