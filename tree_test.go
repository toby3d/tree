package tree_test

import (
	"bytes"
	"embed"
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"testing"

	"source.toby3d.me/toby3d/tree"
)

//go:embed testdata/*
var testData embed.FS

func Example() {
	file, err := testData.Open(filepath.Join(".", "testdata", "html.tree"))
	if err != nil {
		log.Fatalf("cannot open file: %v", err)
	}
	defer file.Close()

	tree := tree.Parse(file)
	fmt.Printf("%#v", tree)
	// Output:
	// html
	//  body
	//   div おはようございます
}

func TestParse(t *testing.T) {
	t.Parallel()

	//nolint: paralleltest // testName used as range value
	for testName, filePath := range map[string]string{
		"html":     filepath.Join(".", "testdata", "html.tree"),
		"math":     filepath.Join(".", "testdata", "math.tree"),
		"nursinos": filepath.Join(".", "testdata", "nursinos.tree"),
		"package":  filepath.Join(".", "testdata", "package.tree"),
		"print":    filepath.Join(".", "testdata", "print.tree"),
	} {
		testName, filePath := testName, filePath

		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			src, err := testData.ReadFile(filePath)
			if err != nil {
				t.Fatalf("cannot open testing file: %v", err)
			}

			result := tree.Parse(bytes.NewReader(src))
			if string(src) == fmt.Sprintf("%#v", result) {
				return
			}

			t.Fatalf(`'%#v' is not equal '%s'`, result, string(src))
		})
	}
}

func TestNode_String(t *testing.T) {
	t.Parallel()

	node := &tree.Node{
		Parent:   nil,
		Value:    "print Hello world",
		Children: nil,
	}

	const expResult string = "print Hello world"
	if node.String() == expResult {
		return
	}

	t.Fatalf(`'%s' is not equal '%s'`, node.String(), expResult)
}

func TestNode_GoString(t *testing.T) {
	t.Parallel()

	root := &tree.Node{
		Parent:   nil,
		Children: make([]*tree.Node, 0),
		Value:    "multiply",
	}
	root.Children = append(root.Children, []*tree.Node{{
		Parent:   root,
		Children: nil,
		Value:    " add 1 1",
	}, {
		Parent:   root,
		Children: nil,
		Value:    " add 2 2",
	}}...)

	const expResult string = "multiply\n add 1 1\n add 2 2"
	if root.GoString() == expResult {
		return
	}

	t.Fatalf(`'%s' is not equal '%s'`, root.GoString(), expResult)
}

func BenchmarkParse(b *testing.B) {
	input := strings.NewReader("title Nursinos Services Offerered\n" +
		"paragraph Here is a list of home nursing services we offer\n" +
		"table\n" +
		" Service Price\n" +
		" BloodPressureCheck $10\n" +
		" TemperatureCheck $5")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = tree.Parse(input)
	}
}
