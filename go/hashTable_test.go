package main

import (
	"os"
	"testing"
)

func TestHashTablePutGet(t *testing.T) {
	table := NewHashTable(10)
	table.Put("key1", "value1")
	table.Put("key2", "value2")

	if table.Get("key1") != "value1" {
		t.Errorf("Expected 'value1', got '%s'", table.Get("key1"))
	}
	if table.Get("key2") != "value2" {
		t.Errorf("Expected 'value2', got '%s'", table.Get("key2"))
	}
}

func TestHashTableRemove(t *testing.T) {
	table := NewHashTable(10)
	table.Put("test", "data")
	table.Remove("test")

	if table.Get("test") != "" {
		t.Errorf("Expected empty string after removal, got '%s'", table.Get("test"))
	}
	if table.GetSize() != 0 {
		t.Errorf("Expected size 0, got %d", table.GetSize())
	}
}

func TestHashTableEmptyHash(t *testing.T) {
	table := NewHashTable(10)

	if table.Get("missing") != "" {
		t.Errorf("Expected empty string for missing key, got '%s'", table.Get("missing"))
	}
	if table.GetSize() != 0 {
		t.Errorf("Expected size 0, got %d", table.GetSize())
	}
}

func TestHashTableUpdateValue(t *testing.T) {
	table := NewHashTable(10)
	table.Put("key", "old")
	table.Put("key", "new")

	if table.Get("key") != "new" {
		t.Errorf("Expected 'new', got '%s'", table.Get("key"))
	}
	if table.GetSize() != 1 {
		t.Errorf("Expected size 1, got %d", table.GetSize())
	}
}

func TestHashTableMultipleCollisions(t *testing.T) {
	table := NewHashTable(2)
	table.Put("aa", "1")
	table.Put("bb", "2")

	if table.GetSize() != 2 {
		t.Errorf("Expected size 2, got %d", table.GetSize())
	}
}

func TestHashTableRemoveNonExisting(t *testing.T) {
	table := NewHashTable(10)
	table.Put("key1", "val1")
	table.Remove("key2")

	if table.GetSize() != 1 {
		t.Errorf("Expected size 1, got %d", table.GetSize())
	}
	if table.Get("key1") != "val1" {
		t.Errorf("Expected 'val1', got '%s'", table.Get("key1"))
	}
}

func TestHashTableSaveLoadText(t *testing.T) {
	table := NewHashTable(10)
	table.Put("k1", "v1")
	table.Put("k2", "v2")

	err := table.SaveToText("hash.txt")
	if err != nil {
		t.Fatal(err)
	}

	table2 := NewHashTable(10)
	err = table2.LoadFromText("hash.txt")
	if err != nil {
		t.Fatal(err)
	}

	if table2.GetSize() != 2 {
		t.Errorf("Expected size 2, got %d", table2.GetSize())
	}
	if table2.Get("k1") != "v1" {
		t.Errorf("Expected 'v1', got '%s'", table2.Get("k1"))
	}
	if table2.Get("k2") != "v2" {
		t.Errorf("Expected 'v2', got '%s'", table2.Get("k2"))
	}

	os.Remove("hash.txt")
}

func TestHashTableLoadFromNonExistingFile(t *testing.T) {
	table := NewHashTable(10)
	err := table.LoadFromText("non_existing.txt")
	if err == nil {
		t.Error("Expected error for non-existing file")
	}

	if table.GetSize() != 0 {
		t.Errorf("Expected size 0, got %d", table.GetSize())
	}
}

func TestHashTableChainOperations(t *testing.T) {
	table := NewHashTable(1)
	table.Put("a", "1")
	table.Put("b", "2")
	table.Put("c", "3")

	if table.GetSize() != 3 {
		t.Errorf("Expected size 3, got %d", table.GetSize())
	}
	if table.Get("a") != "1" {
		t.Errorf("Expected '1', got '%s'", table.Get("a"))
	}
	if table.Get("b") != "2" {
		t.Errorf("Expected '2', got '%s'", table.Get("b"))
	}
	if table.Get("c") != "3" {
		t.Errorf("Expected '3', got '%s'", table.Get("c"))
	}

	table.Remove("b")
	if table.GetSize() != 2 {
		t.Errorf("Expected size 2, got %d", table.GetSize())
	}
	if table.Get("b") != "" {
		t.Errorf("Expected empty string after removal, got '%s'", table.Get("b"))
	}
}

func TestHashTableClearAndReuse(t *testing.T) {
	table := NewHashTable(10)
	table.Put("x", "y")
	table.Remove("x")

	table.Put("new", "value")
	if table.GetSize() != 1 {
		t.Errorf("Expected size 1, got %d", table.GetSize())
	}
	if table.Get("new") != "value" {
		t.Errorf("Expected 'value', got '%s'", table.Get("new"))
	}
}
