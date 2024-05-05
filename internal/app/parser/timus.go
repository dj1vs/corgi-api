package parser

import (
	"bytes"
	"corgi-api/internal/app/ds"
	"corgi-api/internal/app/html_basics"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func ParseTimusProblem(problemID *ds.ProblemID) (ds.ProblemData, error) {
	problemData := ds.ProblemData{}

	url, err := getTimusURL(problemID)
	if err != nil {
		return problemData, err
	}

	resp, err := http.Get(url)
	if err != nil {
		return problemData, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return problemData, err
	}

	node, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		return problemData, err
	}

	contentNode := html_basics.GetElementByAttribute(node, "class", "problem_content")

	for contentChild := contentNode.FirstChild; contentChild != nil; contentChild = contentChild.NextSibling {
		if contentChild.Data == "h2" {
			problemData.Title = contentChild.FirstChild.Data
			continue
		}

		nodeClass, ok := html_basics.GetAttribute(contentChild, "class")

		if ok && nodeClass == "problem_limits" {
			problemLimits := strings.Split(html_basics.CollectText(contentChild, 1), "\n")
			// problemData.TimeLimit = problemLimits[0]
			problemData.MemoryLimit = problemLimits[1]
			continue
		}

		nodeId, ok := html_basics.GetAttribute(contentChild, "id")
		if ok && nodeId == "problem_text" {
			parseProblemText(contentChild, &problemData)
		}

	}

	return problemData, nil
}

func getTimusURL(problemID *ds.ProblemID) (string, error) {
	return "http://acm.timus.ru/problem.aspx?num=" + problemID.Title, nil
}

func parseProblemText(node *html.Node, problemData *ds.ProblemData) error {
	section := description

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if c.Data == "h3" {
			sectionContent := c.FirstChild.Data
			switch sectionContent {
			case "Input":
				section = input
			case "Output":
				section = output
			case "Sample":
				section = sample
			case "Notes":
				section = notes
			}
			continue
		}

		nodeClass, ok := html_basics.GetAttribute(c, "class")

		if !ok {
			continue
		}

		if nodeClass == "problem_par" {
			switch section {
			case description:
				problemData.Description = html_basics.CollectText(c, 2)
			case input:
				problemData.InputDescription = html_basics.CollectText(c, 2)
			case output:
				problemData.OutputDescription = html_basics.CollectText(c, 2)
			case sample:
				// TODO
			case notes:
				problemData.Note = html_basics.CollectText(c, 2)
			}
		}
	}
	return nil
}

type timusSection int

const (
	noSection timusSection = iota
	description
	input
	output
	sample
	notes
)
