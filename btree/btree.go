package btree

const INTERNAL_MAX_KEY = 4
type Node interface {}
type BTreeInternalNode struct {
	nkey   int
	keys   [INTERNAL_MAX_KEY]int
	children [INTERNAL_MAX_KEY]*Node
}

func NewINode() BTreeInternalNode {
	var nkeys = 0
	var children [INTERNAL_MAX_KEY]*Node
	var keys [INTERNAL_MAX_KEY]int
	return BTreeInternalNode{
		nkey:   nkeys,
		keys:   keys,
		children: children,
	}
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
	node.nkey++
}

// split a node into 2 equal parts
func (node *BTreeInternalNode) Split() BTreeInternalNode {
	var newKeys [INTERNAL_MAX_KEY]int;
	var newChildren [INTERNAL_MAX_KEY]*Node
	pos := node.nkey / 2
	for i:= pos; i < node.nkey; i++ {
		newKeys[i - pos] = node.keys[i]
		newChildren[i - pos] = node.children[i]
		node.keys[i] = 0
		node.children[i] = nil
	}
	node.nkey = pos
	return BTreeInternalNode{
		nkey:   node.nkey,
		keys:   newKeys,
		children: newChildren,
	}
}