package mdfmt

import (
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
	lineNumsWithHorizontalRule := []int{}
	// a horizontal rule is 3 or more '-'  with only whitespace as other characters
	//
	for i, line := range lines {
		// remove all spaces
		dashCount := 0
		for _, c := range line {
			if c == '-' {
				dashCount += 1
			} else if !unicode.IsSpace(c) {
				continue
			}
		}
		if dashCount < 3 {
			continue
		}
		// we know that this is a horizontal rule
		lineNumsWithHorizontalRule = append(lineNumsWithHorizontalRule, i)
	}
	// since we called ensureOneEmptyEndLine we do not have to check for last line
	offset := 0
	for _, lineNum := range lineNumsWithHorizontalRule {
		if lines[lineNum+offset+1] != "" {
			lines = append(lines[:lineNum+offset+1], append([]string{""}, lines[lineNum+offset+1:]...)...)
			offset += 1
		}
	}
	return lines
}

func applyToLinks(lines []string, f func(string) string) []string {
	for i, line := range lines {
		openSquareBracketIndex := strings.Index(line, "[")
		if openSquareBracketIndex == -1 {
			continue
		}

		middleIndex := strings.Index(line, "](")
		if middleIndex == -1 {
			continue
		}
		left := line[:middleIndex+2]

		// stack
		closeParenthesesIndex := -1
		parenStack := 0
		for i := middleIndex + 2; i < len(line); i++ {
			switch line[i] {
			case ')':
				{
					if parenStack == 0 {
						closeParenthesesIndex = i
					} else {
						parenStack--
					}
				}
			case '(':
				parenStack++
			}
		}
		if closeParenthesesIndex == -1 {
			continue
		}
		link := line[middleIndex+2 : closeParenthesesIndex]
		right := line[closeParenthesesIndex:]
		// save text to the right of ^
		lines[i] = left + f(link) + right
	}
	return lines
}

// TODO: This needs to be fixed for multiple line links
// TODO: make this work for lines with multiple links
func removeLinkWhitespaces(lines []string) []string {
	removeWhitespace := func(link string) string {
		var builder strings.Builder
		for _, c := range link {
			if unicode.IsSpace(c) {
				builder.WriteString("%20")
			} else {
				builder.WriteRune(c)
			}
		}
		return builder.String()
	}
	return applyToLinks(lines, removeWhitespace)
}

func removeLinkParentheses(lines []string) []string {
	removeParentheses := func(link string) string {
		var builder strings.Builder
		for _, c := range link {
			if c == '(' {
				builder.WriteString("%28")
			} else if c == ')' {
				builder.WriteString("%29")
			} else {
				builder.WriteRune(c)
			}
		}
		return builder.String()
	}
	return applyToLinks(lines, removeParentheses)
}

func ensureHeaderHasEmptyLinesSurrounding(lines []string) []string {
	return lines
}

func reformat(in string) string {
	lines := strings.Split(in, "\n")
	lines = linesOfOnlyWhitespaceBecomeEmpty(lines)
	lines = removeEmptyBeginningLines(lines)
	lines = ensureOneEmptyEndLine(lines)
	lines = ensureHorizontalRuleHasEmptyLineAfter(lines)
	//lines = ensureHeaderHasEmptyLinesSurrounding(lines)
	lines = removeLinkWhitespaces(lines)
	lines = removeLinkParentheses(lines)
	out := strings.Join(lines, "\n")
	return out
}
