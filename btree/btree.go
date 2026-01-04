package btree

const INTERNAL_MAX_KEY = 4
type Node any
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
	head Node
}

func NewBPTree() BPTree {
	newINode := NewINode()
	return BPTree{
		head: &newINode,
	}
}
func (tree *BPTree) insertRecursive(node Node, insertKey int, insertValue int) Node {
    if convert, ok := node.(*BTreeInternalNode); ok {
        // Check if => Internal node: navigate down
        pos := convert.FindLastLE(insertKey)
		// special process for -1 pos
		// insert in the beginning
		if pos == -1 {
			firstLeaf := NewLNode()
			firstLeaf.InsertKv(insertKey, insertValue)
			convert.InsertKv(insertKey, &firstLeaf)
		}else{
        child := convert.children[pos]
		// go down to children and keep inserting
        insertResult := tree.insertRecursive(child, insertKey, insertValue)
        
		// => if nil => no split occurred
        // Child split? Insert promoted key
        if insertResult != nil {
			// Insert the new key-child pair into this internal node
            if childConvert, ok := insertResult.(*BTreeInternalNode); ok {
                convert.InsertKv(childConvert.keys[0], childConvert)
            } else 
			// 
			{
                childConvert := insertResult.(*BTreeLeafNode)
                convert.InsertKv(childConvert.keys[0], childConvert)
            }
        }
        
        // Check if this node needs to split
        if convert.nkey == INTERNAL_MAX_KEY {
            return convert.Split()
        }
		return nil
	}
        
    } else {
        // Leaf node: insert here
        convert := node.(*BTreeLeafNode)
        convert.InsertKv(insertKey, insertValue)
        
        if convert.nkey == INTERNAL_MAX_KEY {
            return convert.Split()
        }
    }
	return nil
}

func (tree *BPTree) Insert(insertKey int, insertValue int) {
    insertResult := tree.insertRecursive(tree.head, insertKey, insertValue)
	if insertResult != nil {
		// Root split: create new root
		childConvert := insertResult.(*BTreeInternalNode)
		newHead := NewINode()
		newHead.nkey = 2
		newHead.keys[0] = tree.head.(*BTreeInternalNode).keys[0]
		newHead.keys[1] = childConvert.keys[0]
		newHead.children[0] = tree.head.(*BTreeInternalNode).children[0]
		newHead.children[1] = childConvert.children[0]
		tree.head = &newHead	
    }
}