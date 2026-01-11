package main

import (
	"os"
	"testing"
)

func TestDoubleListInsertBeforeHeadAndAfterTail(t *testing.T) {
	list := NewDoublyList()
	list.PushBack("B")
	list.InsertBefore("B", "A")
	list.InsertAfter("B", "C")

	if list.GetSize() != 3 {
		t.Errorf("Expected size 3, got %d", list.GetSize())
	}

	if !list.Search("A") {
		t.Error("Expected to find 'A'")
	}
	if !list.Search("B") {
		t.Error("Expected to find 'B'")
	}
	if !list.Search("C") {
		t.Error("Expected to find 'C'")
	}
}

func TestDoubleListSaveLoadBinary(t *testing.T) {
	list := NewDoublyList()
	list.PushBack("x")
	list.PushBack("y")

	err := list.SaveToBinary("dlist.bin")
	if err != nil {
		t.Fatal(err)
	}

	list2 := NewDoublyList()
	err = list2.LoadFromBinary("dlist.bin")
	if err != nil {
		t.Fatal(err)
	}

	if list2.GetSize() != 2 {
		t.Errorf("Expected size 2, got %d", list2.GetSize())
	}

	if !list2.Search("x") {
		t.Error("Expected to find 'x'")
	}
	if !list2.Search("y") {
		t.Error("Expected to find 'y'")
	}

	os.Remove("dlist.bin")
}

func TestDoubleListPopFrontBackAndRemove(t *testing.T) {
	list := NewDoublyList()
	list.PushBack("1")
	list.PushBack("2")
	list.PushBack("3")

	list.PopFront()
	if list.Search("1") {
		t.Error("'1' should have been removed")
	}
	if list.GetSize() != 2 {
		t.Errorf("Expected size 2, got %d", list.GetSize())
	}

	list.PopBack()
	if list.Search("3") {
		t.Error("'3' should have been removed")
	}
	if list.GetSize() != 1 {
		t.Errorf("Expected size 1, got %d", list.GetSize())
	}

	list.RemoveByValue("X")
	if list.GetSize() != 1 {
		t.Errorf("Expected size 1, got %d", list.GetSize())
	}

	list.RemoveByValue("2")
	if list.GetSize() != 0 {
		t.Errorf("Expected size 0, got %d", list.GetSize())
	}
}

func TestDoubleListSaveLoadText(t *testing.T) {
	list := NewDoublyList()
	list.PushBack("a")
	list.PushBack("b")

	err := list.SaveToText("dlist.txt")
	if err != nil {
		t.Fatal(err)
	}

	list2 := NewDoublyList()
	err = list2.LoadFromText("dlist.txt")
	if err != nil {
		t.Fatal(err)
	}

	if list2.GetSize() != 2 {
		t.Errorf("Expected size 2, got %d", list2.GetSize())
	}

	if !list2.Search("a") {
		t.Error("Expected to find 'a'")
	}
	if !list2.Search("b") {
		t.Error("Expected to find 'b'")
	}

	os.Remove("dlist.txt")
}

func TestDoubleListPushFrontOperations(t *testing.T) {
	list := NewDoublyList()
	list.PushFront("c")
	list.PushFront("b")
	list.PushFront("a")

	if list.GetSize() != 3 {
		t.Errorf("Expected size 3, got %d", list.GetSize())
	}

	if !list.Search("a") {
		t.Error("Expected to find 'a'")
	}
}

func TestDoubleListInsertInMiddle(t *testing.T) {
	list := NewDoublyList()
	list.PushBack("a")
	list.PushBack("c")
	list.InsertBefore("c", "b")
	list.InsertAfter("b", "b2")

	if list.GetSize() != 4 {
		t.Errorf("Expected size 4, got %d", list.GetSize())
	}

	if !list.Search("b") {
		t.Error("Expected to find 'b'")
	}
	if !list.Search("b2") {
		t.Error("Expected to find 'b2'")
	}
}

func TestDoubleListPopOperationsOnEmpty(t *testing.T) {
	list := NewDoublyList()
	list.PopFront()
	list.PopBack()
	if list.GetSize() != 0 {
		t.Errorf("Expected size 0, got %d", list.GetSize())
	}
}

func TestDoubleListRemoveNonExistingValue(t *testing.T) {
	list := NewDoublyList()
	list.PushBack("a")
	list.RemoveByValue("b")
	if list.GetSize() != 1 {
		t.Errorf("Expected size 1, got %d", list.GetSize())
	}
}

func TestDoubleListLoadFromNonExistingFiles(t *testing.T) {
	list := NewDoublyList()
	err := list.LoadFromText("non_existing.txt")
	if err == nil {
		t.Error("Expected error for non-existing file")
	}

	err = list.LoadFromBinary("non_existing.bin")
	if err == nil {
		t.Error("Expected error for non-existing binary file")
	}

	if list.GetSize() != 0 {
		t.Errorf("Expected size 0, got %d", list.GetSize())
	}
}

func TestDoubleListClearList(t *testing.T) {
	list := NewDoublyList()
	list.PushBack("a")
	list.PushBack("b")
	list.Clear()

	if list.GetSize() != 0 {
		t.Errorf("Expected size 0, got %d", list.GetSize())
	}
	if list.Search("a") {
		t.Error("'a' should have been cleared")
	}
}

func TestDoubleListGetTailOnEmpty(t *testing.T) {
	list := NewDoublyList()
	if list.GetTail() != "" {
		t.Errorf("Expected empty tail, got %s", list.GetTail())
	}
}

func TestDoubleListPrintMethods(t *testing.T) {
	list := NewDoublyList()
	list.PushBack("a")
	list.PushBack("b")
	list.PrintForward()
	list.PrintBackward()
}
