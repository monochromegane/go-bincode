package bincode

import "path/filepath"

func Generate(root, pkg, outputDir string, ignores []string) error {
	// generate code archive
	err := bindata{
		root:      ".",
		pkg:       pkg,
		outputDir: outputDir,
		ignores:   ignores,
	}.exec()
	if err != nil {
		return err
	}

	// generate bincoder
	return writeToFile(
		filepath.Join(outputDir, "bincoder.go"),
		codeTemplates.toString(),
		binCode{Package: pkg},
	)
}
