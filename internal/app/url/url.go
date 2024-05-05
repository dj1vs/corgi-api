package url

import (
	"corgi_parser/internal/app/ds"
	"errors"
)

func GetProblemURL(problem_id ds.ProblemID) (string, error) {
	switch problem_id.Source {
	case ds.Codeforces:
		if problem_id.Competition == "" {
			return "", errors.New("no competition specified for codeforces problem")
		}
		return "https://codeforces.com/problemset/problem/" + problem_id.Competition + "/" + problem_id.Title, nil
	case ds.Codewars:
		return "http://codewars.com/api/v1/code-challenges/" + problem_id.Title, nil
	case ds.ProjectEuler:
		return "http://projecteuler.net/problem=" + problem_id.Title, nil
	case ds.Timus:
		return "http://acm.timus.ru/problem.aspx?num=" + problem_id.Title, nil
	default:
		return "", errors.New("invalid source")
	}
}

func GetCompetitionURL(comp_id ds.CompetitionID) (string, error) {
	switch comp_id.Source {
	case ds.Codeforces:
		return "https://codeforces.com/contest/" + comp_id.Title, nil
	default:
		return "", errors.New("invalid source")
	}
}
