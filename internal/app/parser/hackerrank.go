package parser

import (
	"bytes"
	"corgi-api/internal/app/ds"
	"corgi-api/internal/app/html_basics"
	"io"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func ParseHackerrankProblem(problemID *ds.ProblemID) (ds.ProblemData, error) {
	problemData := ds.ProblemData{}

	url, err := getHackerrankURL(problemID)
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

	doc, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		return problemData, err
	}

	titleNode := html_basics.GetElementByAttribute(doc, "class", "ui-icon-label page-label")
	problemData.Title = html_basics.CollectText(titleNode, 1)

	challengeBodyNode := html_basics.GetElementByAttribute(doc, "class", "challenge-body-html")

	for c := challengeBodyNode.FirstChild; c != nil; c = c.NextSibling {
		nodeClass, _ := html_basics.GetAttribute(c, "class")
		switch nodeClass {
		case "challenge_problem_statement":
			statement, err := parseHackerrankStatement(c)
			if err == nil {
				problemData.Description += statement
			}
		case "challenge_constraints":
			continue
		case "challenge_sample_input":
			continue
		case "challenge_sample_output":
			continue
		case "challenge_explanation":
			continue
		}
	}

	return problemData, nil
}

func getHackerrankURL(problemID *ds.ProblemID) (string, error) {
	return "http://hackerrank.com/challenges/" + problemID.Title + "/problem", nil
}

// hackerrank draws math equations using svg;
// this is the only site i've found that uses this kind of retarded behaviour
// so right now i don't know how to deal with it
func parseHackerrankStatement(node *html.Node) (string, error) {
	var statementBeginNode *html.Node

	for c := node.FirstChild.FirstChild.FirstChild; c != nil; c = c.NextSibling {
		if c.Data == "p" {
			statementBeginNode = c
			break
		}
	}

	statement := ""

	for c := statementBeginNode; c != nil; c = c.NextSibling {
		log.Println(c.Data)
		if c.Type == html.ElementNode {
			statement += html_basics.CollectText(c, 3)
		}
	}

	return statement, nil
}
