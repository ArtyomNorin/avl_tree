package main

import "fmt"

type AvlTree struct {
	root *avlTreeNode
}

type avlTreeNode struct {
	key    int
	height int
	parent *avlTreeNode
	left   *avlTreeNode
	right  *avlTreeNode
}

func (at *AvlTree) Insert(key int) {
	at.root = at.insert(at.root, key)
}

func (at *AvlTree) insert(node *avlTreeNode, key int) *avlTreeNode {
	if node == nil {
		return &avlTreeNode{key: key, height: 1}
	}

	if key >= node.key {
		node.right = at.insert(node.right, key)
		node.right.parent = node
	} else {
		node.left = at.insert(node.left, key)
		node.left.parent = node
	}

	at.calcHeight(node)

	return at.balance(node)
}

func (at *AvlTree) calcBalanceFactor(node *avlTreeNode) int {
	if node.left == nil && node.right == nil {
		return 0
	}

	if node.left == nil && node.right != nil {
		return node.right.height
	}

	if node.right == nil && node.left != nil {
		return 0 - node.left.height
	}

	return node.right.height - node.left.height
}

func (at *AvlTree) balance(node *avlTreeNode) *avlTreeNode {
	balanceFactor := at.calcBalanceFactor(node)

	if balanceFactor == 2 {
		if at.calcBalanceFactor(node.right) < 0 {
			node.right = at.rightRotate(node.right)
		}

		return at.leftRotate(node)
	} else if balanceFactor == -2 {
		if at.calcBalanceFactor(node.left) > 0 {
			node.left = at.leftRotate(node.left)
		}

		return at.rightRotate(node)
	}

	return node
}

func (at *AvlTree) leftRotate(node *avlTreeNode) *avlTreeNode {
	newNode := node.right

	newNode.parent = node.parent
	node.parent = newNode

	if newNode.left != nil {
		node.right = newNode.left
		node.right.parent = node
	} else {
		node.right = nil
	}

	newNode.left = node

	at.calcHeight(node)
	at.calcHeight(newNode)

	return newNode
}

func (at *AvlTree) rightRotate(node *avlTreeNode) *avlTreeNode {
	newNode := node.left

	newNode.parent = node.parent
	node.parent = newNode

	if newNode.right != nil {
		node.left = newNode.right
		node.left.parent = node
	} else {
		node.left = nil
	}

	newNode.right = node

	at.calcHeight(node)
	at.calcHeight(newNode)

	return newNode
}

func (at *AvlTree) calcHeight(node *avlTreeNode) {
	if node.right == nil && node.left == nil {
		node.height = 1
		return
	}

	if node.left == nil && node.right != nil {
		node.height = node.right.height + 1
		return
	}

	if node.right == nil && node.left != nil {
		node.height = node.left.height + 1
		return
	}

	if node.left.height >= node.right.height {
		node.height = node.left.height + 1
		return
	}

	node.height = node.right.height + 1
}

func (at *AvlTree) FindMin() int {
	min := at.findMin(at.root)

	if min == nil {
		return -1
	}

	return min.key
}

func (at *AvlTree) findMin(node *avlTreeNode) *avlTreeNode {
	if node == nil {
		return nil
	}

	if node.left == nil {
		return node
	}

	return at.findMin(node.left)
}

func (at *AvlTree) FindMax() int {
	max := at.findMax(at.root)

	if max == nil {
		return -1
	}

	return max.key
}

func (at *AvlTree) findMax(node *avlTreeNode) *avlTreeNode {
	if node == nil {
		return nil
	}

	if node.right == nil {
		return node
	}

	return at.findMax(node.right)
}

func (at *AvlTree) Delete(key int) {
	at.root = at.delete(at.root, key)
}

func (at *AvlTree) deleteLeaf(node *avlTreeNode) *avlTreeNode {
	if node.left != nil || node.right != nil {
		return node
	}

	if node.parent.left == node {
		node.parent.left = nil
	} else {
		node.parent.right = nil
	}

	node.parent = nil

	return nil
}

func (at *AvlTree) deleteNodeOnlyWithRightChild(node *avlTreeNode) *avlTreeNode {
	if node.left != nil || node.right == nil {
		return node
	}

	newNode := node.right
	newNode.parent = node.parent

	if node.parent == nil {
		node.right = nil
		return newNode
	}

	if node.parent.right == node {
		node.parent.right = newNode
	} else {
		node.parent.left = newNode
	}

	node.right = nil
	node.parent = nil

	return newNode
}

func (at *AvlTree) deleteNodeOnlyWithLeftChild(node *avlTreeNode) *avlTreeNode {
	if node.left == nil || node.right != nil {
		return node
	}

	newNode := node.left
	newNode.parent = node.parent

	if node.parent == nil {
		node.left = nil
		return newNode
	}

	if node.parent.right == node {
		node.parent.right = newNode
	} else {
		node.parent.left = newNode
	}

	node.left = nil
	node.parent = nil

	return newNode
}

func (at *AvlTree) deleteMin(node *avlTreeNode) *avlTreeNode {
	if node == nil {
		return nil
	}

	if node.left != nil {
		node.left = at.deleteMin(node.left)
		at.calcHeight(node)
		return at.balance(node)
	}

	if node.right != nil {
		return at.deleteNodeOnlyWithRightChild(node)
	}

	return at.deleteLeaf(node)
}

func (at *AvlTree) deleteMax(node *avlTreeNode) *avlTreeNode {
	if node == nil {
		return nil
	}

	if node.right != nil {
		node.right = at.deleteMax(node.right)
		at.calcHeight(node)
		return at.balance(node)
	}

	if node.left != nil {
		return at.deleteNodeOnlyWithLeftChild(node)
	}

	return at.deleteLeaf(node)
}

func (at *AvlTree) deleteNodeWithTwoChild(node *avlTreeNode) *avlTreeNode {
	if node.left == nil || node.right == nil {
		return node
	}

	var newNode *avlTreeNode

	if node.parent == nil || node.parent.right == node {
		newNode = at.findMin(node.right)
		at.deleteMin(node.right)
	} else {
		newNode = at.findMax(node.left)
		at.deleteMax(node.left)
	}

	newNode.parent = node.parent
	newNode.left = node.left
	newNode.right = node.right

	at.calcHeight(newNode)
	return at.balance(newNode)
}

func (at *AvlTree) delete(node *avlTreeNode, key int) *avlTreeNode {
	if node == nil {
		return nil
	}

	if key > node.key {
		node.right = at.delete(node.right, key)
	} else if key < node.key {
		node.left = at.delete(node.left, key)
	} else {
		if node.left == nil && node.right == nil {
			return at.deleteLeaf(node)
		}

		if node.left == nil && node.right != nil {
			return at.deleteNodeOnlyWithRightChild(node)
		}

		if node.left != nil && node.right == nil {
			return at.deleteNodeOnlyWithLeftChild(node)
		}

		return at.deleteNodeWithTwoChild(node)
	}

	at.calcHeight(node)
	return at.balance(node)
}

func (at *AvlTree) Search(key int) int {
	node := at.search(at.root, key)

	if node == nil {
		return -1
	}

	return node.key
}

func (at *AvlTree) search(node *avlTreeNode, key int) *avlTreeNode {
	if node == nil {
		return nil
	}

	if node.key == key {
		return node
	}

	if key >= node.key {
		return at.search(node.right, key)
	}

	return at.search(node.left, key)
}

func (at *AvlTree) Height() int {
	return at.height(at.root)
}

func (at *AvlTree) height(node *avlTreeNode) int {
	if node == nil {
		return 0
	}

	leftHeight := at.height(node.left)
	rightHeight := at.height(node.right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}

	return rightHeight + 1
}

func (at *AvlTree) Print() {
	at.print(at.root)
}

func (at *AvlTree) print(node *avlTreeNode) {
	if node != nil {
		at.print(node.left)
		fmt.Print(node.key, " ")
		at.print(node.right)
	}
}
