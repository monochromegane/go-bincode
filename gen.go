package bincode

import "path/filepath"

func Generate(root, pkg string, ignores []string) error {
	// generate code archive
	err := bindata{root: ".", pkg: pkg, ignores: ignores}.exec()
	if err != nil {
		return err
	}

	// generate bincoder
	return writeToFile(
		filepath.Join(root, "bincoder.go"),
		codeTemplates.toString(),
		binCode{Package: pkg},
	)
}
