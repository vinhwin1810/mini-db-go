package btree

import (
	"fmt"
	"testing"
)

// func Test_Insert(t *testing.T) {
// 	node := NewINode()
// 	node.InsertKv(1, nil)
// 	if node.nkey != 1 {
// 		t.Errorf("Got nkey= %v, expect %v", node.nkey, 1)
// 	}
// 	if node.keys[0] != 1 {
// 		t.Errorf("Got key= %v, expect %v", node.keys[0], 1)
// 	}
// 	node.InsertKv(3, nil)
// 	node.InsertKv(-2, nil)
// 	node.InsertKv(2, nil)
// 	fmt.Print(node)
// }
func Test_Split(t *testing.T) {
	node := NewINode()
	node.InsertKv(1, nil)
	node.InsertKv(3, nil)
	node.InsertKv(5, nil)
	node.InsertKv(7, nil)
	newNode := node.Split()
	fmt.Println("Old Node: ", node.keys)
	fmt.Println("New Node: ", newNode.keys)
	if node.nkey != 2 {
		t.Errorf("Got nkey= %v, expect %v", node.nkey, 2)
	}
	if newNode.nkey != 2 {
		t.Errorf("Got nkey= %v, expect %v", newNode.nkey, 2)
	}
	if node.keys[0] != 1 || node.keys[1] != 3 {
		t.Errorf("Got keys= %v, expect %v", node.keys, [INTERNAL_MAX_KEY]int{1, 3})
	}
	if newNode.keys[0] != 5 || newNode.keys[1] != 7 {
		t.Errorf("Got keys= %v, expect %v", newNode.keys, [INTERNAL_MAX_KEY]int{5, 7})
	}
}