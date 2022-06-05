package app

import (
	"github.com/gocopper/examples/hackernews/pkg/posts"
	"github.com/google/wire"
)

var WireModule = wire.NewSet(
	posts.WireModule,
)
