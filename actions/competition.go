package actions

import (
	"corgi_parser/internal/app/ds"
	"corgi_parser/internal/app/parser"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

func CompetitionHandler(c buffalo.Context) error {
	source, title := c.Param("source"), c.Param("title")

	if source == "" {
		return c.Render(http.StatusBadRequest, r.JSON(map[string]string{"message": "You should provide problem source"}))
	}
	if title == "" {
		return c.Render(http.StatusBadRequest, r.JSON(map[string]string{"message": "You should provide problem title"}))
	}

	compData, err := parser.ParseCompetition(ds.CompetitionID{
		Title:  title,
		Source: source,
	})
	if err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(map[string]string{"message": err.Error()}))
	}

	return c.Render(http.StatusOK, r.JSON(compData))
}
