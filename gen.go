package bincode

func Generate(root, pkg string) error {
	// go-bindata
	err := bindata{root: ".", pkg: pkg}.exec()
	if err != nil {
		return err
	}
	return nil
}
