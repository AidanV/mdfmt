package mdfmt

import (
	"fmt"
	"strings"
	"unicode"
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

func linesOfOnlyWhitespaceBecomeEmpty(lines []string) []string {
	for i, line := range lines {
		if isOnlyWhitespace(line) {
			lines[i] = ""
		}
	}
	return lines
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
	for i, line := range lines {
		if line != "" {
			end = i + 1
		}
	}
	fmt.Println(end)
	if end == -1 {
		return []string{}
	} else {
		return append(lines[:end], "")
	}
}

func insert(a []int, c int, i int) []int {
	return append(a[:i], append([]int{c}, a[i:]...)...)
}

func ensureHorizontalRuleHasEmptyLineAfter(lines []string) []string {
	line_nums_with_horizontal_rule := []int{}
	// a horizontal rule is 3 or more '-'  with only whitespace as other characters
	//
	for i, line := range lines {
		// remove all spaces
		dash_count := 0
		for _, c := range line {
			if c == '-' {
				dash_count += 1
			} else if !unicode.IsSpace(c) {
				continue
			}
		}
		if dash_count < 3 {
			continue
		}
		// we know that this is a horizontal rule
		line_nums_with_horizontal_rule = append(line_nums_with_horizontal_rule, i)
	}
	// since we called ensureOneEmptyEndLine we do not have to check for last line
	offset := 0
	for _, line_num := range line_nums_with_horizontal_rule {
		if lines[line_num+offset+1] != "" {
			lines = append(lines[:line_num+offset+1], append([]string{""}, lines[line_num+offset+1:]...)...)
			offset += 1
		}
	}
	return lines
}

func reformat(in string) string {
	lines := strings.Split(in, "\n")
	// remove empty lines at beginning
	// convert all lines that are just white space to empty lines
	lines = linesOfOnlyWhitespaceBecomeEmpty(lines)
	lines = removeEmptyBeginningLines(lines)
	lines = ensureOneEmptyEndLine(lines)
	lines = ensureHorizontalRuleHasEmptyLineAfter(lines)
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
