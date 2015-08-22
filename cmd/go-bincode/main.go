//go:generate go run main.go -pkg bincode -ignore tests
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/monochromegane/go-bincode"
)

var pkg string
var outputDir string
var ignores []string

func init() {
	flag.StringVar(&pkg, "pkg", "main", "Package name to use in the generated code.")
	flag.StringVar(&outputDir, "output-dir", ".", "Output directory")
	flag.Var((*AppendSliceValue)(&ignores), "ignore", "Regex pattern to ignore")
	flag.Parse()
}
func main() {
	err := bincode.Generate(".", pkg, outputDir, ignores)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

// borrowed from https://github.com/jteeuwen/go-bindata/blob/master/go-bindata/AppendSliceValue.go
// (re-borrowed from https://github.com/hashicorp/serf/blob/master/command/agent/flag_slice_value.go)

// AppendSliceValue implements the flag.Value interface and allows multiple
// calls to the same variable to append a list.
type AppendSliceValue []string

func (s *AppendSliceValue) String() string {
	return strings.Join(*s, ",")
}

func (s *AppendSliceValue) Set(value string) error {
	if *s == nil {
		*s = make([]string, 0, 1)
	}

	*s = append(*s, value)
	return nil
}
