package content

import "github.com/gobuffalo/helpers/hctx"

func New() hctx.Map {
	return hctx.Map{
		"contentOf":  ContentOf,
		"contentFor": ContentFor,
	}
}