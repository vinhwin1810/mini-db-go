package main

import "fmt"

const INTERNAL_MAX_KEY = 8
type Node struct {}
type BTreeInternalNode struct {
	nkey   int
	keys   [INTERNAL_MAX_KEY]int
	children [INTERNAL_MAX_KEY]*Node
}
// find last pos where key <= given key
func (node *BTreeInternalNode) FindLastLE(findKey int) int {
	pos := -1
	for i := 0; i < node.nkey; i++ {
		if node.keys[i] <= findKey {
			pos = i
		} else {
			break
		}
	}
	return pos
}


// insert a key-children pair into the internal node
func (node *BTreeInternalNode) InsertKv(insertKey int, insertChild Node)  {
	pos := node.FindLastLE(insertKey)
	// [1, 3, 4] insert 2
	for i := node.nkey - 1; i > pos; i-- {
		node.keys[i+1] = node.keys[i]
		node.children[i+1] = node.children[i]
	}
	node.keys[pos+1] = insertKey
	node.children[pos+1] = &insertChild
}

func main() {
	fmt.Println("Hello, World!")
}

