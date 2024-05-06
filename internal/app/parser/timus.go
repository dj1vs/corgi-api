package parser

import (
	"bytes"
	"corgi-api/internal/app/ds"
	"corgi-api/internal/app/html_basics"
	"io"
	"net/http"
	"strconv"
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
			//problemData.TimeLimit = problemLimits[0]

			if strings.Contains(problemLimits[0], "second") {
				timeLimitStr := strings.Replace(problemLimits[0], " second", "", 1)[len("Time limit: "):]
				timeLimitFloat, err := strconv.ParseFloat(timeLimitStr, 32)
				if err == nil {
					problemData.TimeLimitMs = int(timeLimitFloat * 1000)
				}
			} else if strings.Contains(problemLimits[0], "millisecond") {
				timeLimitStr := strings.Replace(problemLimits[0], " millisecond", "", 1)[len("Time limit: "):]
				timeLimitFloat, err := strconv.ParseFloat(timeLimitStr, 32)
				if err == nil {
					problemData.TimeLimitMs = int(timeLimitFloat)
				}
			}

			problemData.MemoryLimit = problemLimits[1]
			continue
		}

		nodeId, ok := html_basics.GetAttribute(contentChild, "id")
		if ok && nodeId == "problem_text" {
			parseProblemText(contentChild, &problemData)
		}

	}

	tagsNode := contentNode.NextSibling
	for tagChild := tagsNode.FirstChild; tagChild != nil; tagChild = tagChild.NextSibling {
		if tagChild.Data != "a" {
			continue
		}

		problemData.Tags = append(problemData.Tags, html_basics.CollectText(tagChild, 1))
	}

	linksNode := tagsNode.NextSibling
	problemData.Difficulty = html_basics.CollectText(linksNode.FirstChild, 1)[len("Difficulty: "):]

	totalAttemptsNode := linksNode.FirstChild.NextSibling.NextSibling.
		NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling
	rawTotalAttempts := html_basics.CollectText(totalAttemptsNode, 1)
	totalAttempts, err := strconv.Atoi(rawTotalAttempts[len("All submissions (") : len(rawTotalAttempts)-1])
	if err != nil {
		return problemData, err
	}
	problemData.TotalAttempts = totalAttempts

	rawTotalCompleted := html_basics.CollectText(totalAttemptsNode.NextSibling.NextSibling, 1)
	totalCompleted, err := strconv.Atoi(rawTotalCompleted[len("All accepted submissions (") : len(rawTotalCompleted)-1])
	if err != nil {
		return problemData, err
	}
	problemData.TotalCompleted = totalCompleted

	err = getTimusLanguages(&problemData)
	if err != nil {
		return problemData, nil
	}

	problemData.SourceSizeLimit = "64 KB"

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
				problemData.Description = html_basics.CollectText(c, 3)
			case input:
				problemData.InputDescription = html_basics.CollectText(c, 3)
			case output:
				problemData.OutputDescription = html_basics.CollectText(c, 3)
			case notes:
				problemData.Note = html_basics.CollectText(c, 2)
			}
		} else if nodeClass == "problem_source" {
			if html_basics.CollectText(c.FirstChild, 1) != "Problem Author: " {
				continue
			}
			problemData.Author = strings.Split(html_basics.CollectText(c, 1), "\n")[0]
		}

		if nodeClass == "sample" {
			err := parseSample(c, problemData)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func parseSample(node *html.Node, problemData *ds.ProblemData) error {
	tbodyNode := node.FirstChild

	for trNode := tbodyNode.FirstChild.NextSibling; trNode != nil; trNode = trNode.NextSibling {
		example := ds.ProblemExample{}
		inputNode := trNode.FirstChild.FirstChild
		outputNode := trNode.FirstChild.NextSibling.FirstChild

		example.Input = html_basics.CollectText(inputNode, 1)
		example.Output = html_basics.CollectText(outputNode, 1)

		example.Input = example.Input[:len(example.Input)-1]
		example.Output = example.Output[:len(example.Output)-1]

		problemData.Examples = append(problemData.Examples, example)
	}

	return nil
}

func getTimusLanguages(problemData *ds.ProblemData) error {
	url := "https://acm.timus.ru/submit.aspx"

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	node, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		return err
	}

	languagesNode := html_basics.GetElementByAttribute(node, "name", "Language")
	for languageNode := languagesNode.FirstChild; languageNode != nil; languageNode = languageNode.NextSibling {
		problemData.Languages = append(problemData.Languages, html_basics.CollectText(languageNode, 1))
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
