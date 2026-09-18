package dataStructures

type AVL struct {
	value  int
	left   *AVL
	right  *AVL
	height int
}

func CreateAVL(value int) *AVL {
	return &AVL{
		value:  value,
		left:   nil,
		right:  nil,
		height: 0,
	}
}

func (avl *AVL) AddNode(value int) *AVL {
	if avl.value > value {
		if avl.left == nil {
			avl.left = CreateAVL(value)
		} else {
			avl.left = avl.left.AddNode(value)
		}

	} else {
		if avl.right == nil {
			avl.right = CreateAVL(value)
		} else {
			avl.right = avl.right.AddNode(value)
		}

	}
	avl.height = avl.getHeight()
	balance := getBalance(avl)
	if balance < -1 {
		if getBalance(avl.left) < 0 {
			//Left Left
			return rightRotate(avl)
		} else {
			//Left Right
			avl.left = leftRotate(avl.left)
			return rightRotate(avl)
		}
	} else if balance > 1 {
		if getBalance(avl.right) > 0 {
			//Right Right
			return leftRotate(avl)
		} else {
			//Right left
			avl.right = rightRotate(avl.right)
			return leftRotate(avl)
		}
	}
	return avl
}

func (avl *AVL) DeleteNode(value int) *AVL {
	if avl == nil {
		return nil
	}
	if avl.value == value {
		if avl.left == nil && avl.right == nil {
			return nil
		} else if avl.left == nil {
			return avl.right
		} else if avl.right == nil {
			return avl.left
		} else {
			successor := avl.right.minValueNode()
			avl.value = successor.value
			avl.right = avl.right.DeleteNode(successor.value)
		}
	} else if avl.value > value {
		avl.left = avl.left.DeleteNode(value)
	} else {
		avl.right = avl.right.DeleteNode(value)
	}
	avl.height = avl.getHeight()
	balance := getBalance(avl)

	if balance < -1 {
		if getBalance(avl.left) < 0 {
			//Left Left
			return rightRotate(avl)
		} else {
			//Left Right
			avl.left = leftRotate(avl.left)
			return rightRotate(avl)
		}
	} else if balance > 1 {
		if getBalance(avl.right) > 0 {
			//Right Right
			return leftRotate(avl)
		} else {
			//Right left
			avl.right = rightRotate(avl.right)
			return leftRotate(avl)
		}
	}

	return avl
}

func getBalance(avl *AVL) int {
	if avl.left == nil && avl.right == nil {
		return 0
	}
	if avl.left == nil {
		return avl.right.height
	} else if avl.right == nil {
		return -avl.left.height
	}
	return avl.right.height - avl.left.height
}

func (avl *AVL) getHeight() int {
	if avl.left == nil && avl.right == nil {
		return 0
	}
	if avl.left == nil {
		return 1 + avl.right.height
	}
	if avl.right == nil {
		return 1 + avl.left.height
	}
	return 1 + max(avl.left.height, avl.right.height)
}

func (avl *AVL) minValueNode() *AVL {
	for avl.left != nil {
		avl = avl.left
	}
	return avl
}

func leftRotate(avl *AVL) *AVL {
	newRoot := avl.right
	levelTwo := newRoot.left
	newRoot.left = avl
	avl.right = levelTwo

	avl.height = avl.getHeight()
	newRoot.height = 1 + max(newRoot.left.height, newRoot.right.height)
	return newRoot
}

func rightRotate(avl *AVL) *AVL {
	newRoot := avl.left
	levelTwo := newRoot.right
	avl.left = levelTwo
	newRoot.right = avl

	avl.height = avl.getHeight()
	newRoot.height = 1 + max(newRoot.left.height, newRoot.right.height)
	return newRoot
}
