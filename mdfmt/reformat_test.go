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

