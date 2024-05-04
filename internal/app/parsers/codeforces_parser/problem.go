package codeforces_parser

import (
	"corgi_parser/internal/app/ds"
	"corgi_parser/internal/app/html_basics"

	"golang.org/x/net/html"
)

func ParseProblem(node *html.Node) (ds.ProblemData, error) {
	data := ds.ProblemData{}

	problemStatementNode := html_basics.GetElementByAttribute(node, "class", "problem-statement")

	for statementNode := problemStatementNode.FirstChild; statementNode != nil; statementNode = statementNode.NextSibling {
		nodeClass, ok := html_basics.GetAttribute(statementNode, "class")
		if !ok {
			parseDescription(&data, statementNode)
		}
		switch nodeClass {
		case "header":
			parseHeader(&data, statementNode)
		case "input-specification":
			parseInputDescription(&data, statementNode)
		case "output-specification":
			parseOutputDescription(&data, statementNode)
		case "sample-tests":
			parseSampleTests(&data, statementNode)
		case "note":
			parseNote(&data, statementNode)
		}
	}

	return data, nil
}

func parseHeader(data *ds.ProblemData, node *html.Node) {
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		nodeClass, ok := html_basics.GetAttribute(c, "class")
		if !ok {
			continue
		}

		switch nodeClass {
		case "title":
			data.Title = html_basics.CollectText(c, 1)
		case "time-limit":
			data.TimeLimit = html_basics.CollectText(c, 1)
		case "memory-limit":
			data.MemoryLimit = html_basics.CollectText(c, 1)
		case "input-file":
			data.InputFile = html_basics.CollectText(c, 1)
		case "output-file":
			data.OutputFile = html_basics.CollectText(c, 1)
		}
	}
}

func parseDescription(data *ds.ProblemData, node *html.Node) {
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		data.Description += html_basics.CollectText(c, 3) + "\n"
	}
}

func parseInputDescription(data *ds.ProblemData, node *html.Node) {
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		nodeClass, ok := html_basics.GetAttribute(c, "class")
		if ok && nodeClass != "tex-span" {
			continue
		}
		data.InputDescription += html_basics.CollectText(c, 3) + "\n"
	}
}

func parseOutputDescription(data *ds.ProblemData, node *html.Node) {
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		_, ok := html_basics.GetAttribute(c, "class")
		if ok {
			continue
		}
		data.OutputDescription += html_basics.CollectText(c, 1) + "\n"
	}
}

func parseSampleTests(data *ds.ProblemData, node *html.Node) {
	var testsNode *html.Node

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		nodeClass, ok := html_basics.GetAttribute(c, "class")
		if !ok {
			continue
		}
		if nodeClass == "sample-test" {
			testsNode = c
			break
		}
	}

	example := ds.ProblemExample{}

	for c := testsNode.FirstChild; c != nil; c = c.NextSibling {
		nodeClass, ok := html_basics.GetAttribute(c, "class")
		if !ok {
			continue
		}

		switch nodeClass {
		case "input":
			for inputChild := c.FirstChild; inputChild != nil; inputChild = inputChild.NextSibling {
				nodeClass, ok := html_basics.GetAttribute(inputChild, "class")
				if ok && nodeClass == "title" {
					continue
				}
				example.Input += html_basics.CollectText(inputChild, 1)
			}
		case "output":
			for outputChild := c.FirstChild; outputChild != nil; outputChild = outputChild.NextSibling {
				nodeClass, ok := html_basics.GetAttribute(outputChild, "class")
				if ok && nodeClass == "title" {
					continue
				}
				example.Output += html_basics.CollectText(outputChild, 1)
			}
			data.Examples = append(data.Examples, example)
			example = ds.ProblemExample{}
		}

	}

}

func parseNote(data *ds.ProblemData, node *html.Node) {
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		nodeClass, ok := html_basics.GetAttribute(c, "class")
		if ok && nodeClass == "section-title" {
			continue
		}

		data.Note += html_basics.CollectText(c, 3) + "\n"
	}
}
