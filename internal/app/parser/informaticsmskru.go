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

func ParseInformaticsMskRuProblem(problemID *ds.ProblemID) (ds.ProblemData, error) {
	problemData := ds.ProblemData{}

	url, err := getInformaticsMskRuURL(problemID)
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

	pageNode := html_basics.GetElementByAttribute(doc, "id", "page")
	pageContentNode := html_basics.GetElementByAttribute(pageNode, "id", "region-main-box")

	mainContentBoxNode := html_basics.GetElementByAttribute(pageContentNode, "id", "region-main")
	mainContentContainerNode := html_basics.GetElementByAttribute(mainContentBoxNode, "role", "main")

	var mainContentNode *html.Node = &html.Node{}
	for mainContentContainerChild := mainContentContainerNode.FirstChild; mainContentContainerChild != nil; mainContentContainerChild = mainContentContainerChild.NextSibling {
		switch mainContentContainerChild.Data {
		case "h2":
			problemData.Title = html_basics.CollectText(mainContentContainerChild, 1)
		case "div":
			_, ok := html_basics.GetAttribute(mainContentContainerChild, "class")
			if ok {
				mainContentNode = mainContentContainerChild
			}
		}
	}

	var statementNode *html.Node = &html.Node{}
	var examplesNode *html.Node = &html.Node{}
	for mainContentChild := mainContentNode.FirstChild; mainContentChild != nil; mainContentChild = mainContentChild.NextSibling {
		nodeClass, classOk := html_basics.GetAttribute(mainContentChild, "class")

		if !classOk {
			continue
		}

		if nodeClass != "problem-statement" {
			continue
		}

		firstChildClass, ok := html_basics.GetAttribute(mainContentChild.FirstChild, "class")
		log.Println(firstChildClass, ok)

		if firstChildClass == "sample-tests" {
			examplesNode = mainContentChild
		} else {
			statementNode = mainContentChild
		}

	}

	log.Println(examplesNode)

	for statementChild := statementNode.FirstChild; statementChild != nil; statementChild = statementChild.NextSibling {
		nodeClass, ok := html_basics.GetAttribute(statementChild, "class")

		if !ok {
			continue
		}

		switch nodeClass {
		case "legend":
			problemData.Description = html_basics.CollectText(statementChild, 10)
		case "input-specification":
			problemData.InputDescription = parseInformaticsMskRuStreamDesc(statementChild)
		case "output-specification":
			problemData.OutputDescription = parseInformaticsMskRuStreamDesc(statementChild)
		}
	}

	return problemData, nil
}

func getInformaticsMskRuURL(problemID *ds.ProblemID) (string, error) {
	return "https://informatics.msk.ru/mod/statements/view.php?chapterid=" + problemID.Title, nil
}

func parseInformaticsMskRuStreamDesc(node *html.Node) string {
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if c.Data == "p" {
			return html_basics.CollectText(c, 2)
		}
	}

	return ""
}
