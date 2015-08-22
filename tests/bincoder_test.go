package bincode

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"
)

func TestListCode(t *testing.T) {
	buf := &bytes.Buffer{}
	binCoder{buf}.listCode()

	actual := buf.String()
	expects := []string{
		"bindata.go",
		"cmd/go-bincode/main.go",
		"gen.go",
		"template.go",
		"writer.go",
	}

	for _, e := range expects {
		if !strings.Contains(actual, e) {
			t.Errorf("code list should contain %s, but not contains", e)
		}
	}
}

func TestListCode_Exclude(t *testing.T) {
	buf := &bytes.Buffer{}
	binCoder{buf}.listCode()

	actual := buf.String()
	excludes := []string{
		"bincode.go",
		"bincoder.go",
		".git",
	}

	for _, e := range excludes {
		if strings.Contains(actual, e) {
			t.Errorf("code list should not contain %s, but contains", e)
		}
	}
}

func TestShowCode(t *testing.T) {
	buf := &bytes.Buffer{}
	binCoder{buf}.listCode()

	list := strings.Split(buf.String(), "\n")

	for _, c := range list {
		if c == "" {
			continue
		}
		expect := &bytes.Buffer{}
		binCoder{expect}.showCode(c)

		actual, err := ioutil.ReadFile(filepath.Join("..", c))
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		if expect.String() != string(actual)+"\n" {
			t.Errorf("code should equal %s but not equal", filepath.Join("..", c))
		}
	}

}
