package ds

import (
	"time"
)

type ProblemID struct {
	Source      Source
	Title       string
	Competition string
}

type ProblemData struct {
	Title             string
	TimeLimit         time.Time
	MemoryLimit       string
	SourceSizeLimit   string
	InputFile         string
	OutputFile        string
	Description       string
	InputDescription  string
	OutputDescription string
	Examples          []ProblemExample // make [] ProblemExample
	Note              string
	Author            string
	Tags              []string
	Difficulty        string
	TotalAttempts     int
	TotalCompleted    int
	DateCreated       time.Time
	Languages         []string
}

type ProblemExample struct {
	Input  string
	Output string
}
