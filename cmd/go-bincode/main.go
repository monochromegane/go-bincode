package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/monochromegane/go-bincode"
)

var pkg string

func init() {
	flag.StringVar(&pkg, "pkg", "main", "Package name to use in the generated code.")
	flag.Parse()
}
func main() {
	err := bincode.Generate(".", pkg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
