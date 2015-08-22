package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/monochromegane/go-bincode"
)

var pkg string
var ignores []string

func init() {
	flag.StringVar(&pkg, "pkg", "main", "Package name to use in the generated code.")
	flag.Var((*AppendSliceValue)(&ignores), "ignore", "Regex pattern to ignore")
	flag.Parse()
}
func main() {
	err := bincode.Generate(".", pkg, ignores)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
