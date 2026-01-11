package main

import (
	"encoding/binary"
	"os"
	"strconv"
	"strings"
)

type FBNode struct {
	key   int
	left  *FBNode
	right *FBNode
}

type FullBinaryTree struct {
	root *FBNode
}

func NewFullBinaryTree() *FullBinaryTree {
	return &FullBinaryTree{}
}

func (fbt *FullBinaryTree) TINSERT(key int) {
	fbt.insert(key)
}

func (fbt *FullBinaryTree) TDEL(key int) {
	fbt.remove(key)
}

func (fbt *FullBinaryTree) ISMEMBER(key int) bool {
	return fbt.search(fbt.root, key)
}

func (fbt *FullBinaryTree) TGET(key int) string {
	if fbt.search(fbt.root, key) {
		return strconv.Itoa(key)
	}
	return ""
}

func (fbt *FullBinaryTree) PRINT_PREORDER() string {
	res := make([]int, 0)
	fbt.preorder(fbt.root, &res)
	return fbt.vecToString(res)
}

func (fbt *FullBinaryTree) PRINT_INORDER() string {
	res := make([]int, 0)
	fbt.inorder(fbt.root, &res)
	return fbt.vecToString(res)
}

func (fbt *FullBinaryTree) PRINT_POSTORDER() string {
	res := make([]int, 0)
	fbt.postorder(fbt.root, &res)
	return fbt.vecToString(res)
}

func (fbt *FullBinaryTree) PRINT_BFS() string {
	res := make([]int, 0)
	fbt.bfs(fbt.root, &res)
	return fbt.vecToString(res)
}

func (fbt *FullBinaryTree) insert(key int) {
	if fbt.root == nil {
		fbt.root = &FBNode{key: key}
		return
	}

	queue := []*FBNode{fbt.root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.left == nil {
			current.left = &FBNode{key: key}
			return
		} else if current.right == nil {
			current.right = &FBNode{key: key}
			return
		} else {
			queue = append(queue, current.left, current.right)
		}
	}
}

func (fbt *FullBinaryTree) search(node *FBNode, key int) bool {
	if node == nil {
		return false
	}
	if node.key == key {
		return true
	}
	return fbt.search(node.left, key) || fbt.search(node.right, key)
}

func (fbt *FullBinaryTree) remove(key int) {
	if fbt.root == nil {
		return
	}

	if fbt.root.key == key && fbt.root.left == nil && fbt.root.right == nil {
		fbt.root = nil
		return
	}

	var keyNode, deepest, parentOfDeepest *FBNode
	queue := []struct {
		parent *FBNode
		node   *FBNode
	}{{nil, fbt.root}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.node.key == key {
			keyNode = current.node
		}

		deepest = current.node
		parentOfDeepest = current.parent

		if current.node.left != nil {
			queue = append(queue, struct {
				parent *FBNode
				node   *FBNode
			}{current.node, current.node.left})
		}
		if current.node.right != nil {
			queue = append(queue, struct {
				parent *FBNode
				node   *FBNode
			}{current.node, current.node.right})
		}
	}

	if keyNode == nil {
		return
	}

	keyNode.key = deepest.key

	if parentOfDeepest != nil {
		if parentOfDeepest.left == deepest {
			parentOfDeepest.left = nil
		} else {
			parentOfDeepest.right = nil
		}
	}
}

func (fbt *FullBinaryTree) preorder(node *FBNode, result *[]int) {
	if node == nil {
		return
	}
	*result = append(*result, node.key)
	fbt.preorder(node.left, result)
	fbt.preorder(node.right, result)
}

func (fbt *FullBinaryTree) inorder(node *FBNode, result *[]int) {
	if node == nil {
		return
	}
	fbt.inorder(node.left, result)
	*result = append(*result, node.key)
	fbt.inorder(node.right, result)
}

func (fbt *FullBinaryTree) postorder(node *FBNode, result *[]int) {
	if node == nil {
		return
	}
	fbt.postorder(node.left, result)
	fbt.postorder(node.right, result)
	*result = append(*result, node.key)
}

func (fbt *FullBinaryTree) bfs(node *FBNode, result *[]int) {
	if node == nil {
		return
	}

	queue := []*FBNode{node}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		*result = append(*result, current.key)

		if current.left != nil {
			queue = append(queue, current.left)
		}
		if current.right != nil {
			queue = append(queue, current.right)
		}
	}
}

func (fbt *FullBinaryTree) bfsForSerialization(node *FBNode, result *[]int) {
	if node == nil {
		return
	}

	queue := []*FBNode{node}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		*result = append(*result, current.key)

		if current.left != nil {
			queue = append(queue, current.left)
		}
		if current.right != nil {
			queue = append(queue, current.right)
		}
	}
}

func (fbt *FullBinaryTree) buildCompleteTree(keys []int, index int) *FBNode {
	if index >= len(keys) {
		return nil
	}

	node := &FBNode{key: keys[index]}
	node.left = fbt.buildCompleteTree(keys, 2*index+1)
	node.right = fbt.buildCompleteTree(keys, 2*index+2)

	return node
}

func (fbt *FullBinaryTree) vecToString(vec []int) string {
	if len(vec) == 0 {
		return ""
	}

	strs := make([]string, len(vec))
	for i, v := range vec {
		strs[i] = strconv.Itoa(v)
	}
	return strings.Join(strs, " ")
}

func (fbt *FullBinaryTree) SaveToBinary(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	keys := make([]int, 0)
	fbt.bfsForSerialization(fbt.root, &keys)

	size := len(keys)
	err = binary.Write(file, binary.LittleEndian, int32(size))
	if err != nil {
		return err
	}

	for _, key := range keys {
		err = binary.Write(file, binary.LittleEndian, int32(key))
		if err != nil {
			return err
		}
	}
	return nil
}

func (fbt *FullBinaryTree) LoadFromBinary(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var size int32
	err = binary.Read(file, binary.LittleEndian, &size)
	if err != nil {
		return err
	}

	if size > 0 {
		keys := make([]int, size)
		for i := 0; i < int(size); i++ {
			var key int32
			err = binary.Read(file, binary.LittleEndian, &key)
			if err != nil {
				return err
			}
			keys[i] = int(key)
		}
		fbt.root = fbt.buildCompleteTree(keys, 0)
	} else {
		fbt.root = nil
	}
	return nil
}

func (fbt *FullBinaryTree) Clear() {
	fbt.root = nil
}
