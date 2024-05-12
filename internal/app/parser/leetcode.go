package parser

import (
	"corgi-api/internal/app/ds"
)

// TODO: use chromedp; leetcode wants users to enable js
func ParseLeetcodeProblem(problemID *ds.ProblemID) (ds.ProblemData, error) {
	problemData := ds.ProblemData{}

	return problemData, nil
}

// func getLeetcodeURL(problemID *ds.ProblemID) (string, error) {
// 	return "http://leetcode.com/problems/" + problemID.Title, nil
// }
