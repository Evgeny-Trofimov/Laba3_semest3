package main

import (
	"os"
	"strconv"
	"testing"
)

func TestQueuePushPop(t *testing.T) {
	q := NewQueue(10)
	q.Push("1")
	q.Push("2")
	q.Push("3")

	if q.Pop() != "1" {
		t.Errorf("Expected '1', got '%s'", q.Pop())
	}
	if q.Pop() != "2" {
		t.Errorf("Expected '2', got '%s'", q.Pop())
	}
	if q.GetSize() != 1 {
		t.Errorf("Expected size 1, got %d", q.GetSize())
	}
}

func TestQueueEmptyQueue(t *testing.T) {
	q := NewQueue(10)

	if q.Pop() != "" {
		t.Errorf("Expected empty string, got '%s'", q.Pop())
	}
	if q.Peek() != "" {
		t.Errorf("Expected empty string, got '%s'", q.Peek())
	}
	if q.GetSize() != 0 {
		t.Errorf("Expected size 0, got %d", q.GetSize())
	}
}

func TestQueueMultipleOperations(t *testing.T) {
	q := NewQueue(10)
	q.Push("a")
	q.Push("b")
	q.Pop()
	q.Push("c")

	if q.Pop() != "b" {
		t.Errorf("Expected 'b', got '%s'", q.Pop())
	}
	if q.Peek() != "c" {
		t.Errorf("Expected 'c', got '%s'", q.Peek())
	}
}

func TestQueueResize(t *testing.T) {
	q := NewQueue(1)
	q.Push("1")
	q.Push("2")

	if q.GetSize() != 2 {
		t.Errorf("Expected size 2, got %d", q.GetSize())
	}
}

func TestQueuePeek(t *testing.T) {
	q := NewQueue(10)
	q.Push("test")

	if q.Peek() != "test" {
		t.Errorf("Expected 'test', got '%s'", q.Peek())
	}
	if q.GetSize() != 1 {
		t.Errorf("Expected size 1, got %d", q.GetSize())
	}
}

func TestQueueSaveLoadText(t *testing.T) {
	q := NewQueue(10)
	q.Push("a")
	q.Push("b")
	q.Push("c")

	err := q.SaveToText("queue.txt")
	if err != nil {
		t.Fatal(err)
	}

	q2 := NewQueue(10)
	err = q2.LoadFromText("queue.txt")
	if err != nil {
		t.Fatal(err)
	}

	if q2.GetSize() != 3 {
		t.Errorf("Expected size 3, got %d", q2.GetSize())
	}
	if q2.Pop() != "a" {
		t.Errorf("Expected 'a', got '%s'", q2.Pop())
	}
	if q2.Pop() != "b" {
		t.Errorf("Expected 'b', got '%s'", q2.Pop())
	}
	if q2.Pop() != "c" {
		t.Errorf("Expected 'c', got '%s'", q2.Pop())
	}

	os.Remove("queue.txt")
}

func TestQueueSaveLoadBinary(t *testing.T) {
	q := NewQueue(10)
	q.Push("x")
	q.Push("y")

	err := q.SaveToBinary("queue.bin")
	if err != nil {
		t.Fatal(err)
	}

	q2 := NewQueue(10)
	err = q2.LoadFromBinary("queue.bin")
	if err != nil {
		t.Fatal(err)
	}

	if q2.GetSize() != 2 {
		t.Errorf("Expected size 2, got %d", q2.GetSize())
	}
	if q2.Pop() != "x" {
		t.Errorf("Expected 'x', got '%s'", q2.Pop())
	}
	if q2.Pop() != "y" {
		t.Errorf("Expected 'y', got '%s'", q2.Pop())
	}

	os.Remove("queue.bin")
}

func TestQueueLoadFromNonExistingFiles(t *testing.T) {
	q := NewQueue(10)
	err := q.LoadFromText("non_existing.txt")
	if err == nil {
		t.Error("Expected error for non-existing file")
	}

	err = q.LoadFromBinary("non_existing.bin")
	if err == nil {
		t.Error("Expected error for non-existing binary file")
	}

	if q.GetSize() != 0 {
		t.Errorf("Expected size 0, got %d", q.GetSize())
	}
}

func TestQueueCircularBufferWrap(t *testing.T) {
	q := NewQueue(3)
	q.Push("1")
	q.Push("2")
	q.Push("3")
	q.Pop()
	q.Push("4")

	if q.GetSize() != 3 {
		t.Errorf("Expected size 3, got %d", q.GetSize())
	}
	if q.Pop() != "2" {
		t.Errorf("Expected '2', got '%s'", q.Pop())
	}
	if q.Pop() != "3" {
		t.Errorf("Expected '3', got '%s'", q.Pop())
	}
	if q.Pop() != "4" {
		t.Errorf("Expected '4', got '%s'", q.Pop())
	}
}

func TestQueueMultipleResizeOperations(t *testing.T) {
	q := NewQueue(2)
	for i := 0; i < 10; i++ {
		q.Push(strconv.Itoa(i))
	}

	if q.GetSize() != 10 {
		t.Errorf("Expected size 10, got %d", q.GetSize())
	}
	for i := 0; i < 10; i++ {
		if q.Pop() != strconv.Itoa(i) {
			t.Errorf("Expected '%d', got '%s'", i, q.Pop())
		}
	}
}
