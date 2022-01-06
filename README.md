# tree

[![Go Reference](https://pkg.go.dev/badge/source.toby3d.me/toby3d/tree.svg)](https://pkg.go.dev/source.toby3d.me/toby3d/tree)

The fast and simple [Tree Notation](https://faq.treenotation.org/spec/) parser.
Clean and unfiltered Go with single memory allocation and without any 3rd party
dependencies.

## Usage

```go
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"source.toby3d.me/toby3d/tree"
)

func main() {
	// For example, load tree content from file.
	file, err := os.Open(filepath.Join(".", "html.example"))
	if err != nil {
		log.Fatalf("cannot open file: %v", err)
	}
	defer file.Close()

	// Parse file contents from plain Tree Notation into *tree.Node with
	// *tree.Node childrens.
	node := tree.Parse(file)

	fmt.Printf("%#v", node)
	// Output:
	// html
	//  body
	//   div おはようございます
}
```

See package documentation and another examples
[on pkg.go.dev](https://pkg.go.dev/source.toby3d.me/toby3d/tree).

## Benchmark

```bash
$ GOMAXPROCS=1 go test -bench=Parse -benchmem -benchtime=10s
goos: linux
goarch: amd64
pkg: source.toby3d.me/toby3d/tree
cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
BenchmarkParse  37139248           316.9 ns/op      4096 B/op          1 allocs/op
```

## License

See [LICENSE.md](LICENSE.md)

## Contact

Check [my website](https://toby3d.me/) or [just drop email](mailto:hey@toby3d.me).
