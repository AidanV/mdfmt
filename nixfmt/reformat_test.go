package nixfmt

import (
	"os"
	"testing"
)

func RunTest(testName string) bool {
	in, err := os.ReadFile("../testdata/input/" + testName + ".nix")
	if err != nil {
		panic(err)
	}
	out, err := os.ReadFile("../testdata/output/" + testName + ".nix")
	if err != nil {
		panic(err)
	}
	return reformat(string(in)) == string(out)
}

func TestEmptyFunction(t *testing.T) {
	if !RunTest("EmptyFunction") { // Name of nix file
		t.Error()
	}
}

func TestSimpleRearrange(t *testing.T) {
	if !RunTest("SimpleRearrange") {
		t.Error()
	}
}
