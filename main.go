package main

import "fmt"
var fullName string = 'John Doe'
const INTERNAL_MAX_KEY = 8

type BTreeInternalNode struct {
	keys     [INTERNAL_MAX_KEY]int32
	child [INTERNAL_MAX_KEY]*Node
}
// find last pos where key <= given key
func (node *BTreeInternalNode) FindLastLE(findKey int) int {}


// insert a key-children pair into the internal node
func 

func main() {
	fmt.Println("Hello, World!")
}

