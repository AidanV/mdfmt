package mdfmt

import (
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
	return reformat(string(in)) == string(out)
}

func TestHeaderSpacing(t *testing.T) {
	if !RunTest("HeaderSpacing") { // Name of md file
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

func TestNoTabParagraph(t *testing.T) {
	if !RunTest("NoTabParagraph") {
		t.Error()
	}
}
