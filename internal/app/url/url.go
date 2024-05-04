package url

import (
	"corgi_parser/internal/app/ds"
	"errors"
)

func GetProblemURL(problem_id ds.ProblemID) (string, error) {
	switch problem_id.Source {
	case "codeforces":
		return "https://codeforces.com/problemset/problem/" + problem_id.Competition + "/" + problem_id.Title, nil
	default:
		return "", errors.New("Invalid source: " + problem_id.Source)
	}
}

func GetCompetitionURL(comp_id ds.CompetitionID) (string, error) {
	switch comp_id.Source {
	case "codeforces":
		return "https://codeforces.com/contest/" + comp_id.Title, nil
	default:
		return "", errors.New("Invalid source: " + comp_id.Source)
	}
}
