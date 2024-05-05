package parser

import (
	"bytes"
	"corgi_parser/internal/app/ds"
	"corgi_parser/internal/app/html_basics"
	"io"
	"net/http"

	"golang.org/x/net/html"
)

func ParseProjecteulerProblem(problemID *ds.ProblemID) (ds.ProblemData, error) {
	data := ds.ProblemData{}

	url, err := getProjecteulerURL(problemID)
	if err != nil {
		return data, err
	}

	httpClient := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "corgi-parser")
	resp, err := httpClient.Do(req)
	if err != nil {
		return ds.ProblemData{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ds.ProblemData{}, err
	}

	node, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		return ds.ProblemData{}, err
	}

	contentNode := html_basics.GetElementByAttribute(node, "id", "content")

	for contentChild := contentNode.FirstChild; contentChild != nil; contentChild = contentChild.NextSibling {
		if contentChild.Data == "h2" {
			data.Title = contentChild.FirstChild.Data
			continue
		}

		nodeClass, ok := html_basics.GetAttribute(contentChild, "class")

		if ok && nodeClass == "problem_content" {
			for c := contentChild.FirstChild; c != nil; c = c.NextSibling {
				data.Description += html_basics.CollectText(c, 10)
			}
		}
	}

	return data, nil
}

func getProjecteulerURL(problemID *ds.ProblemID) (string, error) {
	return "http://projecteuler.net/problem=" + problemID.Title, nil
}
