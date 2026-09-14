package dataStructures

type BST struct {
	value int
	left  *BST
	right *BST
}

func CreateBST(value int) *BST {
	return &BST{
		value: value,
		left:  nil,
		right: nil,
	}
}

func (bst *BST) AddNode(value int) *BST {
	if bst == nil {
		bst = CreateBST(value)
	} else if bst.value > value {
		if bst.left == nil {
			bst.left = CreateBST(value)
		} else {
			bst.left = bst.left.AddNode(value)
		}
	} else {
		if bst.right == nil {
			bst.right = CreateBST(value)
		} else {
			bst.right = bst.right.AddNode(value)
		}
	}
	return bst
}

func (bst *BST) DeleteNode(value int) *BST {
	if bst == nil {
		return nil
	}
	if bst.value == value {
		if bst.left == nil && bst.right == nil {
			return nil
		} else if bst.left == nil {
			return bst.right
		} else if bst.right == nil {
			return bst.left
		} else {
			replacement := bst.left
			if replacement.right == nil {
				replacement.right = bst.right
				return replacement
			}
			previous := bst.left
			for replacement != nil && replacement.right != nil {
				previous = replacement
				replacement = replacement.right
			}
			previous.right = replacement.left
			replacement.left = bst.left
			replacement.right = bst.right
			return replacement
		}
	} else if bst.value > value {
		bst.left = bst.left.DeleteNode(value)
	} else {
		bst.right = bst.right.DeleteNode(value)
	}
	return bst
}
