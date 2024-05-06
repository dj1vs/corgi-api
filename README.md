# corgi-api ![](https://github.com/dj1vs/corgi-parser/actions/workflows/go.yml/badge.svg)  

Simple Golang/Buffalo HTTP-server that returns online programming problems from different platforms in a form of unified JSON.

**corgi-api** mostly uses its own web scrappers (due to the lack of the official APIs).

Simple example of a parsed problem:
```url
http://127.0.0.1:3000/problem?source=timus&problem_title=2035
```

```json
{
	"Title": "2035. Another Dress Rehearsal",
	"TimeLimitMs": 500,
	"MemoryLimit": "Memory limit: 64 MB",
	"SourceSizeLimit": "64 KB",
	"InputFile": "",
	"OutputFile": "",
	"Description": "Now Kirill has to recover the lost tests as soon as possible. He has answers to the tests, and he remembers that the summands A and B were integers such that 0 ≤ A ≤ X and 0 ≤ B ≤ Y. Help Kirill recover the tests!",
	"InputDescription": "The only input line contains integers X, Y, and C separated with a space (0 ≤ X, Y, C ≤ 109).",
	"OutputDescription": "If Kirill is wrong and there are no such integers, output “Impossible” (without quotation marks). Otherwise, output the integers A and B separated with a space. If there are several pairs satisfying the conditions, output any of them.",
	"Examples": [
		{
			"Input": "2 7 5",
			"Output": "2 3"
		},
		{
			"Input": "9 15 100",
			"Output": "Impossible"
		}
	],
	"Note": "",
	"Author": "Kirill Borozdin",
	"Tags": null,
	"Difficulty": "72",
	"TotalAttempts": 14976,
	"TotalCompleted": 4315,
	"DateCreated": "0001-01-01T00:00:00Z",
	"Languages": [
		"FreePascal 2.6",
		"Visual C 2022",
		"Visual C++ 2022",
		"Visual C 2022 x64",
		"Visual C++ 2022 x64",
		"GCC 13.2 x64",
		"G++ 13.2 x64",
		"Clang++ 17 x64",
		"Java 1.8",
		"C# .NET 8",
		"Python 3.12 x64",
		"PyPy 3.10 x64",
		"Go 1.14 x64",
		"Haskell 7.6",
		"Scala 2.11",
		"Rust 1.75 x64",
		"Kotlin 1.9.22"
	]
}
```

## Handles:
- `/sources` - Returns all available problem sources
- `/problem` - Returns parsed problem

## Supported platforms:
- Codeforces
- Codewars 
- Project Euler
- Timus

## Planned platforms:
- Code chef
- Leetcode
- SPOJ
- https://informatics.msk.ru//
- https://acmp.ru/
- hackerrank

## Adding a new platform:
Adding a new platform is extremly easy. You just have to follow these 3 steps:

1. Add your new platform in the `ds/sources.go` file
2. Go to `parsers` package folder, and create here new file called `<your platform>.go`. Here, you should write a function `func Parse<Your platform>Problem(problemID *ds.ProblemID) (ds.ProblemData, error)`. You can use `parser/codeforces.go` as an example.
3. Go to `parser/parser.go`. Add your new platform function to the `problemFunctions` map.