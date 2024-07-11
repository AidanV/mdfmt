package mdfmt

import (
	"fmt"
	"strings"
)

/*
https://www.markdownguide.org/basic-syntax/
Rules to follow

# (Headers) have lines of whitespace around them

--- (horizontal rule) have a line of whitespacer after

no tabs to start paragraphs

Any spaces in url should be converted to %20 [link](https://www.example.com/test page) -> [link](https://www.example.com/test%20page)
Parentheses in a url should be replaced with () -> %28%29

i plan on always having empty line at the end
*/

func isOnlyWhitespace(in string) bool {
	return strings.TrimSpace(in) == ""
}

func removeEmptyBeginningLines(lines []string) []string {
	start := -1
	for i, line := range lines {
		if !isOnlyWhitespace(line) {
			start = i
			break
		}
	}
	if start == -1 {
		return []string{}
	} else {
		return lines[start:]
	}

}

func ensureOneEmptyEndLine(lines []string) []string {
	end := -1
	for i := range lines {
		ri := len(lines) - 1 - i
		if !isOnlyWhitespace(lines[ri]) {
			end = ri + 1
		}
	}
	if end == -1 {
		return []string{}
	} else {
		return append(lines[:end], "")
	}
}

func reformat(in string) string {
	lines := strings.Split(in, "\n")
	// remove empty lines at beginning
	lines = removeEmptyBeginningLines(lines)
	lines = ensureOneEmptyEndLine(lines)

	// ensure empty line at end
	//for i := range lines {
	//	if lines[i] == "" {
	//		continue
	//	}
	//	if lines[i][0] == '#' {
	//		if i > 0 {
	//			if lines[i-1] != "" {

	//			}
	//		}

	//	}
	//}
	line := strings.Join(lines, "\n")
	fmt.Print(line)
	return line
}
