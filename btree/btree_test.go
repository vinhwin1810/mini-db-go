package btree

import (
	"fmt"
	"testing"
)

func Test_Insert(t *testing.T) {
	fmt.Println("Test Insert Key-Value into Internal Node")
	node := NewINode()
	node.InsertKv(1, nil)
	if node.nkey != 1 {
		t.Errorf("Got nkey= %v, expect %v", node.nkey, 1)
	}
	if node.keys[0] != 1 {
		t.Errorf("Got key= %v, expect %v", node.keys[0], 1)
	}
	node.InsertKv(3, nil)
	node.InsertKv(-2, nil)
	node.InsertKv(2, nil)
	fmt.Print(node)
}
func Test_Split(t *testing.T) {
	fmt.Println("Test Split Node")
	node := NewINode()
	node.InsertKv(3, nil)
	node.InsertKv(10, nil)
	node.InsertKv(5, nil)
	node.InsertKv(12, nil)
	// node.InsertKv(9, nil)
	newNode := node.Split()
	fmt.Println("Old Node: ", node.keys)
	fmt.Println("New Node: ", newNode.keys)
	if node.nkey != 2 {
		t.Errorf("Got nkey= %v, expect %v", node.nkey, 2)
	}
	if newNode.nkey != 2 {
		t.Errorf("Got nkey= %v, expect %v", newNode.nkey, 2)
	}
	if node.keys[0] != 3 || node.keys[1] != 5 {
		t.Errorf("Got keys= %v, expect %v", node.keys, [INTERNAL_MAX_KEY]int{3, 5})
	}
	if newNode.keys[0] != 10 || newNode.keys[1] != 12 {
		t.Errorf("Got keys= %v, expect %v", newNode.keys, [INTERNAL_MAX_KEY]int{10, 12})
	}
}
func TestLNode(t *testing.T) {
	fmt.Println("Test Leaf Node")
	leaf := NewLNode()
	leaf.InsertKv(1, 100)
	leaf.InsertKv(3, 300)
	leaf.InsertKv(2, 200)
	if leaf.nkey != 3 {
		t.Errorf("Got nkey= %v, expect %v", leaf.nkey, 3)
	}
	if leaf.keys[0] != 1 || leaf.values[0] != 100 {
		t.Errorf("Got key-value= %v-%v, expect %v-%v", leaf.keys[0], leaf.values[0], 1, 100)
	}
	if leaf.keys[1] != 2 || leaf.values[1] != 200 {
		t.Errorf("Got key-value= %v-%v, expect %v-%v", leaf.keys[1], leaf.values[1], 2, 200)
	}
	if leaf.keys[2] != 3 || leaf.values[2] != 300 {
		t.Errorf("Got key-value= %v-%v, expect %v-%v", leaf.keys[2], leaf.values[2], 3, 300)
	}
}
func TestBTree(t *testing.T) {
	fmt.Println("Test B+ Tree")
	tree := NewBPTree()
	//Insert
	// head = []
	tree.Insert(1, 1)
	fmt.Println("Tree Head: ", tree.head.(*BTreeInternalNode).children[0])
	if tree.head.(*BTreeInternalNode).nkey != 1 {
		t.Errorf("Got nkey= %v, expect %v", tree.head.(*BTreeInternalNode).nkey, 1)
	}
	child := tree.head.(*BTreeInternalNode).children[0]
	if (*child).(*BTreeLeafNode).nkey != 1 {
		t.Errorf("Got key= %v, expect %v", (*child).(*BTreeLeafNode).values[0], 1)
	}
}