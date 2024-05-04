package ds

import "strings"

type Source int

const (
	NoSource Source = iota
	Codeforces
)

var (
	sourcesMap = map[string]Source{
		"codeforces": Codeforces,
	}
)

func ParseSourceString(str string) (Source, bool) {
	c, ok := sourcesMap[strings.ToLower(str)]
	return c, ok
}
