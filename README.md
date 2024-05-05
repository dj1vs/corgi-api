# corgi-parser ![](https://github.com/dj1vs/corgi-parser/actions/workflows/go.yml/badge.svg)  

Simple Golang HTTP-server that returns online programming problems in a form of unified JSONs.

Corgi-parser mostly uses its own web scrappers (but, if possible, official APIs)

Take a look at a simple example:
```url
http://127.0.0.1:3000/problem?source=codeforces&problem_title=A&competition=154
```

```json
{
{
	"Title": "A. Hometask",
	"TimeLimit": "2 seconds",
	"MemoryLimit": "256 megabytes",
	"InputFile": "standard input",
	"OutputFile": "standard output",
	"Description": "Sergey attends lessons of the N-ish language. Each lesson he receives a hometask. This time the task is to translate some sentence to the N-ish language. Sentences of the N-ish language can be represented as strings consisting of lowercase Latin letters without spaces or punctuation marks.\nSergey totally forgot about the task until half an hour before the next lesson and hastily scribbled something down. But then he recollected that in the last lesson he learned the grammar of N-ish. The spelling rules state that N-ish contains some \"forbidden\" pairs of letters: such letters can never occur in a sentence next to each other. Also, the order of the letters doesn't matter (for example, if the pair of letters \"ab\" is forbidden, then any occurrences of substrings \"ab\" and \"ba\" are also forbidden). Also, each pair has different letters and each letter occurs in no more than one forbidden pair.\nNow Sergey wants to correct his sentence so that it doesn't contain any \"forbidden\" pairs of letters that stand next to each other. However, he is running out of time, so he decided to simply cross out some letters from the sentence. What smallest number of letters will he have to cross out? When a letter is crossed out, it is \"removed\" so that the letters to its left and right (if they existed), become neighboring. For example, if we cross out the first letter from the string \"aba\", we get the string \"ba\", and if we cross out the second letter, we get \"aa\".\n",
	"InputDescription": "The first line contains a non-empty string s, consisting of lowercase Latin letters — that's the initial sentence in N-ish, written by Sergey. The length of string s doesn't exceed 105.\nThe next line contains integer k (0 ≤ k ≤ 13) — the number of forbidden pairs of letters.\nNext k lines contain descriptions of forbidden pairs of letters. Each line contains exactly two different lowercase Latin letters without separators that represent the forbidden pairs. It is guaranteed that each letter is included in no more than one pair.\n",
	"OutputDescription": "Print the single number — the smallest number of letters that need to be removed to get a string without any forbidden pairs of neighboring letters. Please note that the answer always exists as it is always possible to remove all letters.\n",
	"Examples": [
		{
			"Input": "ababa\n1\nab\n",
			"Output": "2\n"
		},
		{
			"Input": "codeforces\n2\ndo\ncs\n",
			"Output": "1\n"
		}
	],
	"Note": "In the first sample you should remove two letters b.\nIn the second sample you should remove the second or the third letter. The second restriction doesn't influence the solution.\n",
	"Author": "",
	"Tags": [
		"greedy",
		"*1600"
	],
	"Difficulty": "1600",
	"TotalAttempts": "",
	"TotalCompleted": ""
}
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