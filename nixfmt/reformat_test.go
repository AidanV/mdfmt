package nixfmt

import (
	"fmt"
	"os"
	"testing"
)

func TestFmt(t *testing.T) {
	for i := 0; i < 1; i++ {
		in, err := os.ReadFile(fmt.Sprintf("../testdata/input/%d.nix", i))
		if err != nil {
			panic(err)
		}
		out, err := os.ReadFile(fmt.Sprintf("../testdata/output/%d.nix", i))
		if err != nil {
			panic(err)
		}
		if string(in) != string(out) {
			t.Error()
		}
	}
}
