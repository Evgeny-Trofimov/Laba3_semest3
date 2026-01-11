package main

import (
	"os"
	"testing"
)

func TestSinglyListInsertBeforeAfterHeadTail(t *testing.T) {
	list := NewSinglyList()
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
	if list.GetHead() != "A" {
		t.Errorf("Expected head 'A', got '%s'", list.GetHead())
	}
}

func TestSinglyListPopFrontBackOnSingleAndEmpty(t *testing.T) {
	list := NewSinglyList()
	list.PopFront()
	list.PopBack()
	list.PushBack("X")
	list.PopBack()

	if list.GetSize() != 0 {
		t.Errorf("Expected size 0, got %d", list.GetSize())
	}
	if list.GetHead() != "" {
		t.Errorf("Expected empty head, got '%s'", list.GetHead())
	}
}

func TestSinglyListSaveLoadText(t *testing.T) {
	list := NewSinglyList()
	list.PushBack("1")
	list.PushBack("2")

	err := list.SaveToText("slist.txt")
	if err != nil {
		t.Fatal(err)
	}

	list2 := NewSinglyList()
	err = list2.LoadFromText("slist.txt")
	if err != nil {
		t.Fatal(err)
	}

	if list2.GetSize() != 2 {
		t.Errorf("Expected size 2, got %d", list2.GetSize())
	}
	if !list2.Search("1") {
		t.Error("Expected to find '1'")
	}
	if !list2.Search("2") {
		t.Error("Expected to find '2'")
	}

	os.Remove("slist.txt")
}

func TestSinglyListSaveLoadBinary(t *testing.T) {
	list := NewSinglyList()
	list.PushBack("a")
	list.PushBack("b")

	err := list.SaveToBinary("slist.bin")
	if err != nil {
		t.Fatal(err)
	}

	list2 := NewSinglyList()
	err = list2.LoadFromBinary("slist.bin")
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

	os.Remove("slist.bin")
}

func TestSinglyListRemoveByValueNotFoundAndClear(t *testing.T) {
	list := NewSinglyList()
	list.PushBack("1")
	list.PushBack("2")

	list.RemoveByValue("X")
	if list.GetSize() != 2 {
		t.Errorf("Expected size 2, got %d", list.GetSize())
	}
	if !list.Search("1") {
		t.Error("Expected to find '1'")
	}
	if !list.Search("2") {
		t.Error("Expected to find '2'")
	}

	list.Clear()
	if list.GetSize() != 0 {
		t.Errorf("Expected size 0, got %d", list.GetSize())
	}
	if list.GetHead() != "" {
		t.Errorf("Expected empty head, got '%s'", list.GetHead())
	}
	if list.Search("1") {
		t.Error("'1' should have been cleared")
	}
}

func TestSinglyListInsertBeforeMiddleAndRemoveMiddle(t *testing.T) {
	list := NewSinglyList()
	list.PushBack("A")
	list.PushBack("B")
	list.PushBack("C")

	list.InsertBefore("B", "X")
	if list.GetSize() != 4 {
		t.Errorf("Expected size 4, got %d", list.GetSize())
	}
	if !list.Search("X") {
		t.Error("Expected to find 'X'")
	}

	list.RemoveByValue("B")
	if list.Search("B") {
		t.Error("'B' should have been removed")
	}
	if list.GetSize() != 3 {
		t.Errorf("Expected size 3, got %d", list.GetSize())
	}
}

func TestSinglyListPopFrontAndPopBackOnLongList(t *testing.T) {
	list := NewSinglyList()
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
}

func TestSinglyListPushFrontOperations(t *testing.T) {
	list := NewSinglyList()
	list.PushFront("c")
	list.PushFront("b")
	list.PushFront("a")

	if list.GetSize() != 3 {
		t.Errorf("Expected size 3, got %d", list.GetSize())
	}
	if list.GetHead() != "a" {
		t.Errorf("Expected head 'a', got '%s'", list.GetHead())
	}
}

func TestSinglyListInsertBeforeNonExisting(t *testing.T) {
	list := NewSinglyList()
	list.PushBack("a")
	list.InsertBefore("b", "x")

	if list.GetSize() != 1 {
		t.Errorf("Expected size 1, got %d", list.GetSize())
	}
}

func TestSinglyListInsertAfterNonExisting(t *testing.T) {
	list := NewSinglyList()
	list.PushBack("a")
	list.InsertAfter("b", "x")

	if list.GetSize() != 1 {
		t.Errorf("Expected size 1, got %d", list.GetSize())
	}
}

func TestSinglyListLoadFromNonExistingFiles(t *testing.T) {
	list := NewSinglyList()
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

func TestSinglyListRemoveHeadTail(t *testing.T) {
	list := NewSinglyList()
	list.PushBack("a")
	list.PushBack("b")
	list.PushBack("c")

	list.RemoveByValue("a")
	if list.GetHead() != "b" {
		t.Errorf("Expected head 'b', got '%s'", list.GetHead())
	}

	list.RemoveByValue("c")
	if list.GetSize() != 1 {
		t.Errorf("Expected size 1, got %d", list.GetSize())
	}
	if !list.Search("b") {
		t.Error("Expected to find 'b'")
	}
}

func TestSinglyListPrintMethod(t *testing.T) {
	list := NewSinglyList()
	list.PushBack("test")
	list.Print()
}
