package codewars_parser

import (
	"corgi_parser/internal/app/ds"
	"encoding/json"
	"io"
	"net/http"
)

func ParseProblem(problem_url string) (ds.ProblemData, error) {
	data := ds.ProblemData{}

	resp, err := http.Get(problem_url)
	if err != nil {
		return ds.ProblemData{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ds.ProblemData{}, err
	}

	var apiProblem ds.CodewarsApiProblem
	err = json.Unmarshal(body, &apiProblem)
	if err != nil {
		return ds.ProblemData{}, err
	}
	data.Title = apiProblem.Name

	data.Description = apiProblem.Description

	return data, nil
}
