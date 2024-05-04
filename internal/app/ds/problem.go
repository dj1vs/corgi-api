package ds

type ProblemID struct {
	Source      Source
	Title       string
	Competition string
}

type ProblemData struct {
	Title             string
	TimeLimit         string //TODO: make it time.Time
	MemoryLimit       string
	InputFile         string
	OutputFile        string
	Description       string
	InputDescription  string
	OutputDescription string
	Examples          []ProblemExample // make [] ProblemExample
	Note              string
}

type ProblemExample struct {
	Input  string
	Output string
}
