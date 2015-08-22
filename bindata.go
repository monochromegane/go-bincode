package bincode

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type bindata struct {
	root    string
	pkg     string
	ignores []string
}

func (b bindata) exec() error {
	files, err := b.gofiles()
	if err != nil {
		return err
	}

	args := []string{"-o", "bincode.go", "-pkg", b.pkg}

	if len(b.ignores) > 0 {
		for _, i := range b.ignores {
			args = append(args, "-ignore", i)
		}
	}

	args = append(args, files...)
	return exec.Command("go-bindata", args...).Run()
}

func (b bindata) gofiles() ([]string, error) {
	var files []string
	err := filepath.Walk(b.root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			name := info.Name()
			// Skip hidden directory
			if strings.HasPrefix(name, ".") && len(name) > 1 {
				return filepath.SkipDir
			}
			return nil
		}

		// Only .go file
		if !strings.HasSuffix(info.Name(), ".go") {
			return nil
		}

		// Exclude generated code
		if info.Name() == "bincode.go" || info.Name() == "bincoder.go" {
			return nil
		}

		files = append(files, path)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return files, nil
}
