# corgi_parser
Simple Golang HTTP-server that returns online programming problems in a form of unified JSONs.

Take a look at a simple example:
```url
http://127.0.0.1:3000/problem?source=codeforces&problem_title=A&competition=1
```

```json
{
	"Title": "A. Theatre Square",
	"TimeLimit": "1 second",
	"MemoryLimit": "256 megabytes",
	"InputFile": "standard input",
	"OutputFile": "standard output",
	"Description": "Theatre Square in the capital city of Berland has a rectangular shape with the size n × m meters. On the occasion of the city's anniversary, a decision was taken to pave the Square with square granite flagstones. Each flagstone is of the size a × a.\nWhat is the least number of flagstones needed to pave the Square? It's allowed to cover the surface larger than the Theatre Square, but the Square has to be covered. It's not allowed to break the flagstones. The sides of flagstones should be parallel to the sides of the Square.\n",
	"InputDescription": "The input contains three positive integer numbers in the first line: n,  m and a (1 ≤  n, m, a ≤ 109).\n",
	"OutputDescription": "Write the needed number of flagstones.\n",
	"Examples": [
		{
			"Input": "6 6 4\n",
			"Output": "4\n"
		}
	],
	"Note": ""
}
```

Sometimes it uses official API, sometimes it has to scrap information from problems html page.

## Handles:
- `/problem`
- `/competition` - returns competitions title and a list of competitions problems

## Supported platforms:
- Codeforces
- Codewars (only title and description)
- Project Euler (only title and description)