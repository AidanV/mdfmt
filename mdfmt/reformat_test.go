package mdfmt

import (
	"fmt"
	"os"
	"testing"
)

func RunTest(testName string) bool {
	in, err := os.ReadFile("../testdata/input/" + testName + ".md")
	if err != nil {
		panic(err)
	}
	out, err := os.ReadFile("../testdata/output/" + testName + ".md")
	if err != nil {
		panic(err)
	}
	return Reformat(string(in)) == string(out)
}

func TestRemoveEmptyBeginningLines(t *testing.T) {
	lines := []string{"", "non empty"}
	lines = removeEmptyBeginningLines(lines)
	if len(lines) != 1 || lines[0] != "non empty" {
		t.Error()
	}
}

func TestEnsureOneEmptyEndLine(t *testing.T) {
	lines := []string{"non empty", "", ""}
	lines = ensureOneEmptyEndLine(lines)
	fmt.Print(lines)
	if len(lines) != 2 || lines[0] != "non empty" || lines[1] != "" {
		t.Error()
	}
}

func TestEmptyLineEnd(t *testing.T) {
	if !RunTest("EmptyLineEnd") { // Name of md file
		t.Error()
	}
}

func TestHeaderSpacing(t *testing.T) {
	if !RunTest("HeaderSpacing") {
		t.Error()
	}
}

func TestLineBreakSpacing(t *testing.T) {
	if !RunTest("LineBreakSpacing") {
		t.Error()
	}
}

func TestLinkNoWhitespace(t *testing.T) {
	if !RunTest("LinkNoWhitespace") {
		t.Error()
	}
}

func TestLinkParentheses(t *testing.T) {
	if !RunTest("LinkParentheses") {
		t.Error()
	}
}

func TestNoEmptyLineBeginning(t *testing.T) {
	if !RunTest("NoEmptyLineBeginning") {
		t.Error()
	}
}

func TestNoTabParagraph(t *testing.T) {
	if !RunTest("NoTabParagraph") {
		t.Error()
	}
}

func TestAll(t *testing.T) {
	if !RunTest("All") {
		t.Error()
	}
}
