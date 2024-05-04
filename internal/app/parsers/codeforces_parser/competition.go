package codeforces_parser

import (
	"corgi_parser/internal/app/ds"
	"corgi_parser/internal/app/html_basics"

	"golang.org/x/net/html"
)

func ParseCompetition(node *html.Node) (ds.CompetitionData, error) {
	data := ds.CompetitionData{}

	tableNode := html_basics.GetElementByAttribute(node, "class", "problems")
	var tableBodyNode *html.Node
	for c := tableNode.FirstChild; c != nil; c = c.NextSibling {
		if c.Data == "tbody" {
			tableBodyNode = c
			break
		}
	}

	for tableRow := tableBodyNode.FirstChild; tableRow != nil; tableRow = tableRow.NextSibling {
		if tableRow.Type == html.TextNode || tableRow.Data != "tr" {
			continue
		}
		for rowCell := tableRow.FirstChild; rowCell != nil; rowCell = rowCell.NextSibling {
			nodeClasses, ok := html_basics.GetAttributes(rowCell, "class")
			if !ok {
				continue
			}

			for _, nodeClass := range nodeClasses {
				if nodeClass == "id" {
					rawProblemTitle := html_basics.CollectText(rowCell.FirstChild.NextSibling, 1)
					problemTitle := ""
					for _, titleChar := range rawProblemTitle {
						if titleChar != '\n' && titleChar != ' ' {
							problemTitle += string(titleChar)
						}
					}
					data.Problems = append(data.Problems, problemTitle)
				}
			}

		}
	}

	return data, nil
}
