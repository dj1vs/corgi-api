package actions

import (
	"corgi-api/internal/app/ds"
	"corgi-api/internal/app/parser"
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

	sourceParsed, ok := ds.ParseSourceString(source)
	if !ok {
		return c.Render(http.StatusBadRequest, r.JSON(map[string]string{"message": "Source  " + source + " is not supported."}))
	}

	compData, err := parser.ParseCompetition(ds.CompetitionID{
		Title:  title,
		Source: sourceParsed,
	})
	if err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(map[string]string{"message": err.Error()}))
	}

	return c.Render(http.StatusOK, r.JSON(compData))
}
