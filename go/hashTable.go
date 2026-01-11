package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type HashNode struct {
	key   string
	value string
}

type HashTable struct {
	table    [][]HashNode
	capacity int
	size     int
}

func NewHashTable(cap int) *HashTable {
	if cap <= 0 {
		cap = 10
	}
	return &HashTable{
		table:    make([][]HashNode, cap),
		capacity: cap,
		size:     0,
	}
}

func (ht *HashTable) hashFunction(key string) int {
	hash := 0
	for _, c := range key {
		hash = (hash*31 + int(c)) % ht.capacity
	}
	return hash
}

func (ht *HashTable) Put(key, value string) {
	index := ht.hashFunction(key)
	for i, node := range ht.table[index] {
		if node.key == key {
			ht.table[index][i].value = value
			return
		}
	}
	ht.table[index] = append(ht.table[index], HashNode{key: key, value: value})
	ht.size++
}

func (ht *HashTable) Get(key string) string {
	index := ht.hashFunction(key)
	for _, node := range ht.table[index] {
		if node.key == key {
			return node.value
		}
	}
	return ""
}

func (ht *HashTable) Remove(key string) {
	index := ht.hashFunction(key)
	for i, node := range ht.table[index] {
		if node.key == key {
			ht.table[index] = append(ht.table[index][:i], ht.table[index][i+1:]...)
			ht.size--
			return
		}
	}
}

func (ht *HashTable) GetSize() int {
	return ht.size
}

func (ht *HashTable) SaveToText(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	fmt.Fprintf(writer, "%d\n", ht.size)

	for _, chain := range ht.table {
		for _, node := range chain {
			fmt.Fprintf(writer, "%s %s\n", node.key, node.value)
		}
	}

	return writer.Flush()
}

func (ht *HashTable) LoadFromText(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	ht.table = make([][]HashNode, ht.capacity)
	ht.size = 0

	scanner := bufio.NewScanner(file)

	if !scanner.Scan() {
		return scanner.Err()
	}

	var newSize int
	fmt.Sscanf(scanner.Text(), "%d", &newSize)

	for i := 0; i < newSize; i++ {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) >= 2 {
			ht.Put(parts[0], parts[1])
		}
	}

	return scanner.Err()
}
