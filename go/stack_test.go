package main

import (
	"os"
	"strconv"
	"testing"
)

func TestStackEmptyStack(t *testing.T) {
	s := NewStack(10)

	if s.GetSize() != 0 {
		t.Errorf("Expected size 0, got %d", s.GetSize())
	}
	if s.Pop() != "" {
		t.Errorf("Expected empty string, got '%s'", s.Pop())
	}
	if s.Peek() != "" {
		t.Errorf("Expected empty string, got '%s'", s.Peek())
	}
}

func TestStackResizeInternalBuffer(t *testing.T) {
	s := NewStack(1)
	s.Push("a")
	s.Push("b")
	s.Push("c")

	if s.GetSize() != 3 {
		t.Errorf("Expected size 3, got %d", s.GetSize())
	}
	if s.Pop() != "c" {
		t.Errorf("Expected 'c', got '%s'", s.Pop())
	}
	if s.Pop() != "b" {
		t.Errorf("Expected 'b', got '%s'", s.Pop())
	}
	if s.Pop() != "a" {
		t.Errorf("Expected 'a', got '%s'", s.Pop())
	}
}

func TestStackSaveLoadText(t *testing.T) {
	s := NewStack(10)
	s.Push("one")
	s.Push("two")

	err := s.SaveToText("stack_txt.txt")
	if err != nil {
		t.Fatal(err)
	}

	s2 := NewStack(10)
	err = s2.LoadFromText("stack_txt.txt")
	if err != nil {
		t.Fatal(err)
	}

	if s2.GetSize() != 2 {
		t.Errorf("Expected size 2, got %d", s2.GetSize())
	}
	if s2.Pop() != "two" {
		t.Errorf("Expected 'two', got '%s'", s2.Pop())
	}
	if s2.Pop() != "one" {
		t.Errorf("Expected 'one', got '%s'", s2.Pop())
	}

	os.Remove("stack_txt.txt")
}

func TestStackSaveLoadBinary(t *testing.T) {
	s := NewStack(10)
	s.Push("x")
	s.Push("y")

	err := s.SaveToBinary("stack_bin.dat")
	if err != nil {
		t.Fatal(err)
	}

	s2 := NewStack(10)
	err = s2.LoadFromBinary("stack_bin.dat")
	if err != nil {
		t.Fatal(err)
	}

	if s2.GetSize() != 2 {
		t.Errorf("Expected size 2, got %d", s2.GetSize())
	}
	if s2.Pop() != "y" {
		t.Errorf("Expected 'y', got '%s'", s2.Pop())
	}
	if s2.Pop() != "x" {
		t.Errorf("Expected 'x', got '%s'", s2.Pop())
	}

	os.Remove("stack_bin.dat")
}

func TestStackLoadFromNonExistingFiles(t *testing.T) {
	s := NewStack(10)
	err := s.LoadFromText("no_such_file.txt")
	if err == nil {
		t.Error("Expected error for non-existing file")
	}

	err = s.LoadFromBinary("no_such_file.bin")
	if err == nil {
		t.Error("Expected error for non-existing binary file")
	}

	if s.GetSize() != 0 {
		t.Errorf("Expected size 0, got %d", s.GetSize())
	}
}

func TestStackZeroInitialCapacity(t *testing.T) {
	s := NewStack(0)
	s.Push("a")

	if s.GetSize() != 1 {
		t.Errorf("Expected size 1, got %d", s.GetSize())
	}
}

func TestStackPushWithoutResize(t *testing.T) {
	s := NewStack(10)
	s.Push("a")
	s.Push("b")

	if s.GetSize() != 2 {
		t.Errorf("Expected size 2, got %d", s.GetSize())
	}
	if s.Peek() != "b" {
		t.Errorf("Expected 'b', got '%s'", s.Peek())
	}
}

func TestStackMultiplePushPop(t *testing.T) {
	s := NewStack(10)
	for i := 0; i < 100; i++ {
		s.Push(strconv.Itoa(i))
	}

	if s.GetSize() != 100 {
		t.Errorf("Expected size 100, got %d", s.GetSize())
	}

	for i := 99; i >= 0; i-- {
		if s.Pop() != strconv.Itoa(i) {
			t.Errorf("Expected '%d', got '%s'", i, s.Pop())
		}
	}

	if s.GetSize() != 0 {
		t.Errorf("Expected size 0, got %d", s.GetSize())
	}
}

func TestStackPeekOnNonEmpty(t *testing.T) {
	s := NewStack(10)
	s.Push("first")
	s.Push("second")

	if s.Peek() != "second" {
		t.Errorf("Expected 'second', got '%s'", s.Peek())
	}
	s.Pop()
	if s.Peek() != "first" {
		t.Errorf("Expected 'first', got '%s'", s.Peek())
	}
}

func TestStackConstructWithDifferentCapacities(t *testing.T) {
	s1 := NewStack(5)
	if s1.GetSize() != 0 {
		t.Errorf("Expected size 0, got %d", s1.GetSize())
	}
	s1.Push("test")
	if s1.Pop() != "test" {
		t.Errorf("Expected 'test', got '%s'", s1.Pop())
	}

	s2 := NewStack(100)
	for i := 0; i < 50; i++ {
		s2.Push(strconv.Itoa(i))
	}
	if s2.GetSize() != 50 {
		t.Errorf("Expected size 50, got %d", s2.GetSize())
	}
}
