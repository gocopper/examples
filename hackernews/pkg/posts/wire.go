package posts

import "github.com/google/wire"

var WireModule = wire.NewSet(
	NewQueries,
)
