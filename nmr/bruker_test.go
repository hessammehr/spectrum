package nmr

import (
	"fmt"
	"path"
	"testing"
)

const testexpt = "../ma-catalyst/1"

func TestParseAcqu(t *testing.T) {
	acquFile := path.Join(testexpt, "acqu")
	fmt.Printf("Parsing acqu file %s\n", acquFile)

	vars := parseAcqu(acquFile)
	fmt.Printf("%v vars found\n", len(vars))
	for _, i := range []string{"SW", "O1", "NS", "TD"} {
		fmt.Printf("%v = %v\n", i, vars[i])
	}
}

func TestReadBruker(t *testing.T) {
	expt, _ := ReadBruker(testexpt)
	fmt.Printf("%v\n", expt)
}
