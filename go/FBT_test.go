package main

import (
	"os"
	"strconv"
	"strings"
	"testing"
)

func getElements(str string) []int {
	if str == "" {
		return []int{}
	}

	parts := strings.Fields(str)
	result := make([]int, len(parts))
	for i, part := range parts {
		val, _ := strconv.Atoi(part)
		result[i] = val
	}
	return result
}

func TestFullBinaryTreeInsertBasic(t *testing.T) {
	tree := NewFullBinaryTree()
	tree.TINSERT(10)

	if !tree.ISMEMBER(10) {
		t.Error("Expected 10 to be in tree")
	}

	if tree.PRINT_INORDER() != "10" {
		t.Errorf("Expected inorder '10', got '%s'", tree.PRINT_INORDER())
	}
}

func TestFullBinaryTreeInsertMultiple(t *testing.T) {
	tree := NewFullBinaryTree()
	tree.TINSERT(30)
	tree.TINSERT(10)
	tree.TINSERT(20)

	if !tree.ISMEMBER(30) {
		t.Error("Expected 30 to be in tree")
	}
	if !tree.ISMEMBER(10) {
		t.Error("Expected 10 to be in tree")
	}
	if !tree.ISMEMBER(20) {
		t.Error("Expected 20 to be in tree")
	}

	inorder := tree.PRINT_INORDER()
	elements := getElements(inorder)
	if len(elements) != 3 {
		t.Errorf("Expected 3 elements, got %d", len(elements))
	}
}

func TestFullBinaryTreeDeleteLeaf(t *testing.T) {
	tree := NewFullBinaryTree()
	tree.TINSERT(1)
	tree.TINSERT(2)
	tree.TINSERT(3)
	tree.TDEL(3)

	if tree.ISMEMBER(3) {
		t.Error("3 should have been deleted")
	}
}

func TestFullBinaryTreeDeleteRoot(t *testing.T) {
	tree := NewFullBinaryTree()
	tree.TINSERT(2)
	tree.TINSERT(1)
	tree.TINSERT(3)
	tree.TDEL(2)

	if tree.ISMEMBER(2) {
		t.Error("2 should have been deleted")
	}
	if !tree.ISMEMBER(1) {
		t.Error("Expected 1 to still be in tree")
	}
	if !tree.ISMEMBER(3) {
		t.Error("Expected 3 to still be in tree")
	}
}

func TestFullBinaryTreeAllTraversalsProduceSameElements(t *testing.T) {
	tree := NewFullBinaryTree()
	tree.TINSERT(3)
	tree.TINSERT(1)
	tree.TINSERT(2)

	inorder := getElements(tree.PRINT_INORDER())
	preorder := getElements(tree.PRINT_PREORDER())
	postorder := getElements(tree.PRINT_POSTORDER())
	bfs := getElements(tree.PRINT_BFS())

	if len(inorder) != 3 {
		t.Errorf("Expected 3 elements, got %d", len(inorder))
	}

	inorderMap := make(map[int]bool)
	for _, v := range inorder {
		inorderMap[v] = true
	}

	checkMap := func(name string, arr []int) {
		for _, v := range arr {
			if !inorderMap[v] {
				t.Errorf("%s: Element %d not found in inorder", name, v)
			}
		}
	}

	checkMap("preorder", preorder)
	checkMap("postorder", postorder)
	checkMap("bfs", bfs)
}

func TestFullBinaryTreeEmptyTree(t *testing.T) {
	tree := NewFullBinaryTree()

	if tree.ISMEMBER(999) {
		t.Error("Empty tree should not contain 999")
	}

	if tree.PRINT_INORDER() != "" {
		t.Errorf("Expected empty string for empty tree, got '%s'", tree.PRINT_INORDER())
	}

	if tree.TGET(999) != "" {
		t.Errorf("Expected empty string for non-existent key, got '%s'", tree.TGET(999))
	}
}

func TestFullBinaryTreeTGETMethod(t *testing.T) {
	tree := NewFullBinaryTree()
	tree.TINSERT(42)
	tree.TINSERT(24)
	tree.TINSERT(100)

	if tree.TGET(42) != "42" {
		t.Errorf("Expected '42', got '%s'", tree.TGET(42))
	}
	if tree.TGET(24) != "24" {
		t.Errorf("Expected '24', got '%s'", tree.TGET(24))
	}
	if tree.TGET(100) != "100" {
		t.Errorf("Expected '100', got '%s'", tree.TGET(100))
	}
	if tree.TGET(999) != "" {
		t.Errorf("Expected empty string for non-existent key, got '%s'", tree.TGET(999))
	}
}

func TestFullBinaryTreeSaveLoadBinary(t *testing.T) {
	tree := NewFullBinaryTree()
	tree.TINSERT(5)
	tree.TINSERT(3)
	tree.TINSERT(7)
	tree.TINSERT(2)
	tree.TINSERT(4)

	err := tree.SaveToBinary("fulltree_test.bin")
	if err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat("fulltree_test.bin"); os.IsNotExist(err) {
		t.Fatal("File should have been created")
	}

	tree2 := NewFullBinaryTree()
	err = tree2.LoadFromBinary("fulltree_test.bin")
	if err != nil {
		t.Fatal(err)
	}

	if !tree2.ISMEMBER(5) {
		t.Error("Expected 5 to be in tree")
	}
	if !tree2.ISMEMBER(3) {
		t.Error("Expected 3 to be in tree")
	}
	if !tree2.ISMEMBER(7) {
		t.Error("Expected 7 to be in tree")
	}
	if !tree2.ISMEMBER(2) {
		t.Error("Expected 2 to be in tree")
	}
	if !tree2.ISMEMBER(4) {
		t.Error("Expected 4 to be in tree")
	}

	os.Remove("fulltree_test.bin")
}

func TestFullBinaryTreeClearTree(t *testing.T) {
	tree := NewFullBinaryTree()
	tree.TINSERT(1)
	tree.TINSERT(2)
	tree.TINSERT(3)

	if !tree.ISMEMBER(1) {
		t.Error("Expected 1 to be in tree")
	}

	tree.Clear()
	if tree.ISMEMBER(1) {
		t.Error("1 should have been cleared")
	}
	if tree.PRINT_INORDER() != "" {
		t.Errorf("Expected empty string after clear, got '%s'", tree.PRINT_INORDER())
	}
}

func TestFullBinaryTreeInsertManyElements(t *testing.T) {
	tree := NewFullBinaryTree()

	for i := 1; i <= 15; i++ {
		tree.TINSERT(i)
	}

	for i := 1; i <= 15; i++ {
		if !tree.ISMEMBER(i) {
			t.Errorf("Expected %d to be in tree", i)
		}
		if tree.TGET(i) != strconv.Itoa(i) {
			t.Errorf("Expected '%d', got '%s'", i, tree.TGET(i))
		}
	}

	bfs := tree.PRINT_BFS()
	elements := getElements(bfs)
	if len(elements) != 15 {
		t.Errorf("Expected 15 elements in BFS, got %d", len(elements))
	}
}

func TestFullBinaryTreeDeleteAndReinsert(t *testing.T) {
	tree := NewFullBinaryTree()
	tree.TINSERT(10)
	tree.TINSERT(20)
	tree.TINSERT(30)
	tree.TINSERT(40)

	tree.TDEL(20)
	if tree.ISMEMBER(20) {
		t.Error("20 should have been deleted")
	}
	if !tree.ISMEMBER(10) {
		t.Error("Expected 10 to still be in tree")
	}
	if !tree.ISMEMBER(30) {
		t.Error("Expected 30 to still be in tree")
	}
	if !tree.ISMEMBER(40) {
		t.Error("Expected 40 to still be in tree")
	}

	tree.TINSERT(50)
	if !tree.ISMEMBER(50) {
		t.Error("Expected 50 to be in tree")
	}

	inorder := tree.PRINT_INORDER()
	elements := getElements(inorder)
	if len(elements) != 4 {
		t.Errorf("Expected 4 elements after reinsertion, got %d", len(elements))
	}
}

func TestFullBinaryTreeDeleteNonExisting(t *testing.T) {
	tree := NewFullBinaryTree()
	tree.TINSERT(1)
	tree.TINSERT(2)
	tree.TDEL(999)

	if !tree.ISMEMBER(1) {
		t.Error("Expected 1 to still be in tree")
	}
	if !tree.ISMEMBER(2) {
		t.Error("Expected 2 to still be in tree")
	}
}

func TestFullBinaryTreeDeleteNodeWithTwoChildren(t *testing.T) {
	tree := NewFullBinaryTree()
	tree.TINSERT(10)
	tree.TINSERT(5)
	tree.TINSERT(15)
	tree.TINSERT(3)
	tree.TINSERT(7)
	tree.TINSERT(12)
	tree.TINSERT(20)

	tree.TDEL(5)

	if tree.ISMEMBER(5) {
		t.Error("5 should have been deleted")
	}
	if !tree.ISMEMBER(10) {
		t.Error("Expected 10 to still be in tree")
	}
	if !tree.ISMEMBER(15) {
		t.Error("Expected 15 to still be in tree")
	}
	if !tree.ISMEMBER(3) {
		t.Error("Expected 3 to still be in tree")
	}
	if !tree.ISMEMBER(7) {
		t.Error("Expected 7 to still be in tree")
	}
	if !tree.ISMEMBER(12) {
		t.Error("Expected 12 to still be in tree")
	}
	if !tree.ISMEMBER(20) {
		t.Error("Expected 20 to still be in tree")
	}
}

func TestFullBinaryTreeLoadFromNonExistingFile(t *testing.T) {
	tree := NewFullBinaryTree()
	err := tree.LoadFromBinary("non_existing.bin")
	if err == nil {
		t.Error("Expected error for non-existing file")
	}

	if tree.ISMEMBER(1) {
		t.Error("Tree should be empty after failed load")
	}
	if tree.PRINT_INORDER() != "" {
		t.Errorf("Expected empty string after failed load, got '%s'", tree.PRINT_INORDER())
	}
}

func TestFullBinaryTreeSingleNodeTree(t *testing.T) {
	tree := NewFullBinaryTree()
	tree.TINSERT(100)

	if !tree.ISMEMBER(100) {
		t.Error("Expected 100 to be in tree")
	}
	if tree.TGET(100) != "100" {
		t.Errorf("Expected '100', got '%s'", tree.TGET(100))
	}
	if tree.PRINT_INORDER() != "100" {
		t.Errorf("Expected '100', got '%s'", tree.PRINT_INORDER())
	}

	tree.TDEL(100)
	if tree.ISMEMBER(100) {
		t.Error("100 should have been deleted")
	}
	if tree.PRINT_INORDER() != "" {
		t.Errorf("Expected empty string after deletion, got '%s'", tree.PRINT_INORDER())
	}
}

func TestFullBinaryTreeMultipleDeleteOperations(t *testing.T) {
	tree := NewFullBinaryTree()
	for i := 1; i <= 7; i++ {
		tree.TINSERT(i)
	}

	tree.TDEL(4)
	tree.TDEL(2)
	tree.TDEL(6)

	if tree.ISMEMBER(4) {
		t.Error("4 should have been deleted")
	}
	if tree.ISMEMBER(2) {
		t.Error("2 should have been deleted")
	}
	if tree.ISMEMBER(6) {
		t.Error("6 should have been deleted")
	}

	for _, i := range []int{1, 3, 5, 7} {
		if !tree.ISMEMBER(i) {
			t.Errorf("Expected %d to still be in tree", i)
		}
	}
}
