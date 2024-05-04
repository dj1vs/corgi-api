package parser

import (
	"bytes"
	"corgi_parser/internal/app/ds"
	"corgi_parser/internal/app/parsers/codeforces_parser"
	"corgi_parser/internal/app/parsers/codewars_parser"
	"corgi_parser/internal/app/url"
	"io"
	"net/http"

	"golang.org/x/net/html"
)

func ParseProblem(problem_id ds.ProblemID) (ds.ProblemData, error) {
	problem_url, err := url.GetProblemURL(problem_id)
	if err != nil {
		return ds.ProblemData{}, err
	}

	switch problem_id.Source {
	case ds.Codeforces:
		return codeforces_parser.ParseProblem(problem_url)
	case ds.Codewars:
		return codewars_parser.ParseProblem(problem_url)
	}

	return ds.ProblemData{}, nil
}

func ParseCompetition(comp_id ds.CompetitionID) (ds.CompetitionData, error) {
	comp_url, err := url.GetCompetitionURL(comp_id)
	if err != nil {
		return ds.CompetitionData{}, err
	}

	resp, err := http.Get(comp_url)
	if err != nil {
		return ds.CompetitionData{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ds.CompetitionData{}, err
	}

	doc, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		return ds.CompetitionData{}, err
	}

	data := ds.CompetitionData{}

	switch comp_id.Source {
	case ds.Codeforces:
		data, err = codeforces_parser.ParseCompetition(doc)
	}

	data.Title = comp_id.Title

	return data, err
}
