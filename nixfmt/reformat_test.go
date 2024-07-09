package nixfmt

import (
	"fmt"
	"os"
	"testing"
)

func TestFmt(t *testing.T) {
	files, _ := os.ReadDir("../testdata/input")
	for i := 0; i < len(files); i++ {
		in, err := os.ReadFile(fmt.Sprintf("../testdata/input/%d.nix", i))
		if err != nil {
			panic(err)
		}
		out, err := os.ReadFile(fmt.Sprintf("../testdata/output/%d.nix", i))
		if err != nil {
			panic(err)
		}

		if reformat(string(in)) != string(out) {
			t.Error()
		}
	}
}
