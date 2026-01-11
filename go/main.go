package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	exitCode := m.Run()

	filesToRemove := []string{
		"arr.txt", "arr.bin",
		"dlist.bin", "dlist.txt",
		"queue.txt", "queue.bin",
		"stack_txt.txt", "stack_bin.dat",
		"hash.txt",
		"fulltree_test.bin",
		"slist.txt", "slist.bin",
	}

	for _, file := range filesToRemove {
		os.Remove(file)
	}

	os.Exit(exitCode)
}
