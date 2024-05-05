package parser

import (
	"corgi_parser/internal/app/ds"
	"encoding/json"
	"io"
	"net/http"
)

func ParseCodewarsProblem(problemID *ds.ProblemID) (ds.ProblemData, error) {
	data := ds.ProblemData{}

	url, err := getCodewarsURL(problemID)
	if err != nil {
		return data, err
	}

	resp, err := http.Get(url)
	if err != nil {
		return ds.ProblemData{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ds.ProblemData{}, err
	}

	var apiProblem codewarsApiProblem
	err = json.Unmarshal(body, &apiProblem)
	if err != nil {
		return ds.ProblemData{}, err
	}
	data.Title = apiProblem.Name

	data.Description = apiProblem.Description

	return data, nil
}

func getCodewarsURL(problemID *ds.ProblemID) (string, error) {
	return "http://codewars.com/api/v1/code-challenges/" + problemID.Title, nil
}

type codewarsApiProblem struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
