package btree

const INTERNAL_MAX_KEY = 4
type Node interface {
	FindLastLE(findKey int) int
	InsertKv(insertKey int, insertChild Node)
	Split() Node
}
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
type BTreeLeafNode struct {
	nkey   int
	keys   [INTERNAL_MAX_KEY]int
	values [INTERNAL_MAX_KEY]int
}

func NewLNode() BTreeLeafNode {
	var nkeys = 0
	var newValues [INTERNAL_MAX_KEY]int
	var newKeys [INTERNAL_MAX_KEY]int
	return BTreeLeafNode{
		nkey:   nkeys,
		keys:   newKeys,
		values: newValues,
	}
}
// find last pos where key <= given key
func (node *BTreeLeafNode) FindLastLE(findKey int) int {
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
func (node *BTreeLeafNode) InsertKv(insertKey int, insertValue int)  {
	pos := node.FindLastLE(insertKey)
	// [1, 3, 4] insert 2
	for i := node.nkey - 1; i > pos; i-- {
		node.keys[i+1] = node.keys[i]
		node.values[i+1] = node.values[i]
	}
	node.keys[pos+1] = insertKey
	node.values[pos+1] = insertValue
	node.nkey++
}

// split a node into 2 equal parts
func (node *BTreeLeafNode) Split() BTreeLeafNode {
	var newKeys [INTERNAL_MAX_KEY]int;
	var newValues [INTERNAL_MAX_KEY]int
	pos := node.nkey / 2
	for i:= pos; i < node.nkey; i++ {
		newKeys[i - pos] = node.keys[i]
		newValues[i - pos] = node.values[i]
		node.keys[i] = 0
		node.values[i] = 0
	}
	node.nkey = pos
	return BTreeLeafNode{
		nkey:   node.nkey,
		keys:   newKeys,
		values: newValues,
	}
}

//B+Tree Structure
type BPTree struct {
	head *Node
}

func NewBPTree() BPTree {
	var head Node
	return BPTree{
		head: &head,
	}
}