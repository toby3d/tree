package tree

import (
	"bufio"
	"io"
	"strings"
)

// Node represent a single data structure.
type Node struct {
	Parent   *Node
	Children []*Node
	Value    string
}

const (
	// nodeBreakSymbol delimits nodes (lines).
	nodeBreakSymbol rune = '\n'

	// wordBreakSymbol delimits words (cells).
	wordBreakSymbol rune = ' '

	// edgeSymbol is used to indicate the parent/child relationship between
	// nodes.
	//
	// TODO(toby3d): allow clients to change this to '\t' for example.
	edgeSymbol rune = wordBreakSymbol
)

// Parse parses the data into a Tree Notation nodes.
//
// Note what all documents are valid Tree Notation documents. Like binary
// notation, there are no syntax errors in Tree Notation.
func Parse(data io.Reader) (root *Node) {
	stack := make([]*Node, 0)

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		node := &Node{
			Parent:   nil,
			Children: nil,
			Value:    scanner.Text(),
		}

		if root == nil {
			root = node
			stack = append(stack, root)

			continue
		}

		if lenIndent(node.Value) <= lenIndent(root.Value) && len(stack) > 0 {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

		if root.Children == nil {
			root.Children = make([]*Node, 0)
		}

		node.Parent = root
		root.Children = append(root.Children, node)
		stack = append(stack, node)
		root = stack[len(stack)-1]
	}

	if root == nil {
		return nil
	}

	for root.Parent != nil {
		root = root.Parent
	}

	return root
}

// Satisfies the fmt.Stringer interface to print node line passed in as an
// operand to any format that accepts a string, or to an unformatted printer
// such as fmt.Print.
func (n Node) String() string {
	return strings.TrimLeft(n.Value, string(edgeSymbol))
}

// GoString satisfies the fmt.GoStringer interface to print values passed as an
// operand to a %#v format.
func (n Node) GoString() (result string) {
	result += n.Value

	for i := range n.Children {
		result += string(nodeBreakSymbol)
		result += n.Children[i].GoString()
	}

	return result
}

// lenIndent count egdeSymbol prefixes in line.
//
// Returns 0 if line starts from any word character.
func lenIndent(v string) int {
	count := 0

	for i := range v {
		if rune(v[i]) != edgeSymbol {
			break
		}

		count++
	}

	return count
}
