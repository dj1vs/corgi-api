package actions

import (
	"corgi_parser/internal/app/ds"
	"corgi_parser/internal/app/parser"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

func ProblemHandler(c buffalo.Context) error {
	source, title, competition := c.Param("source"), c.Param("problem_title"), c.Param("competition")
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

	problemData, err := parser.ParseProblem(ds.ProblemID{
		Source:      sourceParsed,
		Title:       title,
		Competition: competition,
	})

	if err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(map[string]string{"message": err.Error()}))
	}

	return c.Render(http.StatusOK, r.JSON(problemData))
}
