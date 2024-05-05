package actions

import (
	"corgi_parser/internal/app/ds"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

func SourcesHandler(c buffalo.Context) error {
	sources := []string{}

	for source := range ds.SourcesMap {
		sources = append(sources, source)
	}

	return c.Render(http.StatusOK, r.JSON(map[string]([]string){"available_sources": sources}))
}
