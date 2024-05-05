package parser

import (
	"corgi_parser/internal/app/ds"
	"errors"
)

func ParseProblem(problem_id ds.ProblemID) (ds.ProblemData, error) {
	for source, problemFunc := range problemFunctions {
		if source == problem_id.Source {
			return problemFunc(&problem_id)
		}
	}

	return ds.ProblemData{}, errors.New("no such source for problem search")
}

func ParseCompetition(comp_id ds.CompetitionID) (ds.CompetitionData, error) {
	for source, compFunc := range compFunctions {
		if source == comp_id.Source {
			return compFunc(&comp_id)
		}
	}

	return ds.CompetitionData{}, errors.New("no such source for problem search")
}

var (
	problemFunctions = map[ds.Source](func(problemID *ds.ProblemID) (ds.ProblemData, error)){
		ds.Codeforces:   ParseCodeforcesProblem,
		ds.Codewars:     ParseCodewarsProblem,
		ds.ProjectEuler: ParseProjecteulerProblem,
		ds.Timus:        ParseTimusProblem,
	}

	compFunctions = map[ds.Source](func(compID *ds.CompetitionID) (ds.CompetitionData, error)){
		ds.Codeforces: ParseCodeforcesCompetition,
	}
)
