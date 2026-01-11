package main

import (
	"os"
	"testing"
)

func TestArrayExceptions(t *testing.T) {
	arr := NewArray(10)

	_, err := arr.Get(-1)
	if err == nil {
		t.Error("Expected error for Get(-1)")
	}

	_, err = arr.Get(0)
	if err == nil {
		t.Error("Expected error for Get(0) on empty array")
	}

	err = arr.RemoveAt(0)
	if err == nil {
		t.Error("Expected error for RemoveAt(0) on empty array")
	}

	err = arr.InsertAt(1, "x")
	if err == nil {
		t.Error("Expected error for InsertAt(1, 'x') on empty array")
	}
}

func TestArraySaveLoadTextBinary(t *testing.T) {
	arr := NewArray(10)
	arr.PushBack("a")
	arr.PushBack("b")

	err := arr.SaveToText("arr.txt")
	if err != nil {
		t.Fatal(err)
	}

	arr2 := NewArray(10)
	err = arr2.LoadFromText("arr.txt")
	if err != nil {
		t.Fatal(err)
	}

	if arr2.GetSize() != 2 {
		t.Errorf("Expected size 2, got %d", arr2.GetSize())
	}

	val, err := arr2.Get(0)
	if err != nil || val != "a" {
		t.Errorf("Expected 'a', got %v", val)
	}

	err = arr.SaveToBinary("arr.bin")
	if err != nil {
		t.Fatal(err)
	}

	arr3 := NewArray(10)
	err = arr3.LoadFromBinary("arr.bin")
	if err != nil {
		t.Fatal(err)
	}

	if arr3.GetSize() != 2 {
		t.Errorf("Expected size 2, got %d", arr3.GetSize())
	}

	val, err = arr3.Get(1)
	if err != nil || val != "b" {
		t.Errorf("Expected 'b', got %v", val)
	}

	os.Remove("arr.txt")
	os.Remove("arr.bin")
}

func TestArrayBasicOperations(t *testing.T) {
	arr := NewArray(10)
	arr.PushBack("1")
	arr.PushBack("3")
	arr.PushFront("0")
	arr.InsertAt(2, "2")

	if arr.GetSize() != 4 {
		t.Errorf("Expected size 4, got %d", arr.GetSize())
	}

	val, _ := arr.Get(0)
	if val != "0" {
		t.Errorf("Expected '0', got %s", val)
	}

	val, _ = arr.Get(1)
	if val != "1" {
		t.Errorf("Expected '1', got %s", val)
	}

	val, _ = arr.Get(2)
	if val != "2" {
		t.Errorf("Expected '2', got %s", val)
	}

	val, _ = arr.Get(3)
	if val != "3" {
		t.Errorf("Expected '3', got %s", val)
	}

	arr.PopFront()
	arr.PopBack()
	if arr.GetSize() != 2 {
		t.Errorf("Expected size 2, got %d", arr.GetSize())
	}
}

func TestArrayFindAndSet(t *testing.T) {
	arr := NewArray(10)
	arr.PushBack("a")
	arr.PushBack("b")
	arr.PushBack("c")

	if arr.Find("b") != 1 {
		t.Errorf("Expected index 1 for 'b', got %d", arr.Find("b"))
	}

	if arr.Find("d") != -1 {
		t.Errorf("Expected -1 for 'd', got %d", arr.Find("d"))
	}

	arr.Set(1, "x")
	val, _ := arr.Get(1)
	if val != "x" {
		t.Errorf("Expected 'x', got %s", val)
	}
}

func TestArrayRemoveAtValid(t *testing.T) {
	arr := NewArray(10)
	arr.PushBack("a")
	arr.PushBack("b")
	arr.PushBack("c")

	arr.RemoveAt(1)
	if arr.GetSize() != 2 {
		t.Errorf("Expected size 2, got %d", arr.GetSize())
	}

	val, _ := arr.Get(0)
	if val != "a" {
		t.Errorf("Expected 'a', got %s", val)
	}

	val, _ = arr.Get(1)
	if val != "c" {
		t.Errorf("Expected 'c', got %s", val)
	}
}

func TestArrayPopOperationsOnEmpty(t *testing.T) {
	arr := NewArray(10)
	arr.PopBack()
	arr.PopFront()
	if arr.GetSize() != 0 {
		t.Errorf("Expected size 0, got %d", arr.GetSize())
	}
}

func TestArrayResizeCapacity(t *testing.T) {
	arr := NewArray(2)
	arr.PushBack("1")
	arr.PushBack("2")
	arr.PushBack("3")

	if arr.GetSize() != 3 {
		t.Errorf("Expected size 3, got %d", arr.GetSize())
	}

	val, _ := arr.Get(2)
	if val != "3" {
		t.Errorf("Expected '3', got %s", val)
	}
}

func TestArrayLoadFromNonExistingFiles(t *testing.T) {
	arr := NewArray(10)
	err := arr.LoadFromText("non_existing.txt")
	if err == nil {
		t.Error("Expected error for non-existing file")
	}

	err = arr.LoadFromBinary("non_existing.bin")
	if err == nil {
		t.Error("Expected error for non-existing binary file")
	}

	if arr.GetSize() != 0 {
		t.Errorf("Expected size 0, got %d", arr.GetSize())
	}
}

func TestArrayInsertAtBoundaries(t *testing.T) {
	arr := NewArray(10)
	arr.PushBack("a")
	arr.InsertAt(0, "first")
	arr.InsertAt(2, "last")

	if arr.GetSize() != 3 {
		t.Errorf("Expected size 3, got %d", arr.GetSize())
	}

	val, _ := arr.Get(0)
	if val != "first" {
		t.Errorf("Expected 'first', got %s", val)
	}

	val, _ = arr.Get(2)
	if val != "last" {
		t.Errorf("Expected 'last', got %s", val)
	}
}

func TestArrayEmptyArrayOperations(t *testing.T) {
	arr := NewArray(10)
	if arr.GetSize() != 0 {
		t.Errorf("Expected size 0, got %d", arr.GetSize())
	}

	err := arr.Set(0, "x")
	if err == nil {
		t.Error("Expected error for Set on empty array")
	}

	if arr.Find("x") != -1 {
		t.Errorf("Expected -1 for 'x' in empty array, got %d", arr.Find("x"))
	}
}

func TestArrayPrintMethod(t *testing.T) {
	arr := NewArray(10)
	arr.PushBack("test1")
	arr.PushBack("test2")
	arr.Print()
}
