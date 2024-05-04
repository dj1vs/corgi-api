package ds

import "strings"

type Source int

const (
	NoSource Source = iota
	Codeforces
	Codewars
)

var (
	sourcesMap = map[string]Source{
		"codeforces": Codeforces,
		"codewars":   Codewars,
	}
)

func ParseSourceString(str string) (Source, bool) {
	c, ok := sourcesMap[strings.ToLower(str)]
	return c, ok
}
