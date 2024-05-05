package parser

import (
	"bytes"
	"corgi-api/internal/app/ds"
	"corgi-api/internal/app/html_basics"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html"
)

func ParseCodeforcesProblem(problemID *ds.ProblemID) (ds.ProblemData, error) {
	data := ds.ProblemData{}

	url, err := getCodeforcesURL(problemID)
	if err != nil {
		return data, err
	}

	resp, err := http.Get(url)
	if err != nil {
		return data, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return data, err
	}

	node, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		return data, err
	}

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

	sidebarNode := html_basics.GetElementByAttribute(node, "id", "sidebar")

	for sidebarChild := sidebarNode.FirstChild; sidebarChild != nil; sidebarChild = sidebarChild.NextSibling {
		nodeClass, ok := html_basics.GetAttribute(sidebarChild, "class")
		if !ok {
			continue
		}

		if nodeClass != "roundbox sidebox borderTopRound " {
			continue
		}

		if !strings.Contains(sidebarChild.FirstChild.NextSibling.FirstChild.Data, "Problem tags") {
			continue
		}

		targetNode := sidebarChild.FirstChild.NextSibling.NextSibling.NextSibling

		for tagNode := targetNode.FirstChild; tagNode != targetNode.LastChild.PrevSibling; tagNode = tagNode.NextSibling {
			if tagNode.Type == html.TextNode {
				continue
			}
			rawTag := html_basics.CollectText(tagNode, 2)
			rawTag = strings.ReplaceAll(rawTag, "\n", "")
			rawTag = strings.TrimSpace(rawTag)

			if len(rawTag) < 1 {
				continue
			}

			if rawTag[0] == '*' && len(rawTag) > 2 {
				data.Difficulty = rawTag[1:]
			}

			data.Tags = append(data.Tags, rawTag)
		}

	}

	return data, nil
}

func ParseCodeforcesCompetition(compID *ds.CompetitionID) (ds.CompetitionData, error) {
	data := ds.CompetitionData{}

	url, err := getCodeforcesCompURL(compID)
	if err != nil {
		return data, err
	}

	resp, err := http.Get(url)
	if err != nil {
		return data, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return data, err
	}

	node, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		return data, err
	}

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
			nodeClasses, ok := html_basics.GetAttribute(rowCell, "class")
			if !ok {
				continue
			}

			for _, nodeClass := range strings.Split(nodeClasses, " ") {
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

func getCodeforcesURL(problemID *ds.ProblemID) (string, error) {
	if problemID.Competition == "" {
		return "", errors.New("no competition specified for codeforces problem")
	}
	return "https://codeforces.com/problemset/problem/" + problemID.Competition + "/" + problemID.Title, nil
}

func getCodeforcesCompURL(compID *ds.CompetitionID) (string, error) {
	return "https://codeforces.com/contest/" + compID.Title, nil
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
			rawTimeLimit := strings.Split(html_basics.CollectText(c, 1), " ")
			if len(rawTimeLimit) == 2 {
				timeUnitStr := ""
				switch rawTimeLimit[1] {
				case "seconds":
					timeUnitStr = "s"
				case "milliseconds":
					timeUnitStr = "us"
				}

				timeLimit, err := time.ParseDuration(rawTimeLimit[0] + timeUnitStr)
				if err == nil {
					data.TimeLimitMs = int(timeLimit.Milliseconds())
				}
			}
			// data.TimeLimit = html_basics.CollectText(c, 1)
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
			example.Input = example.Input[:len(example.Input)-1]
			example.Output = example.Output[:len(example.Output)-1]

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
