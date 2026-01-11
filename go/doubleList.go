package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

type DNode struct {
	data string
	next *DNode
	prev *DNode
}

type DoublyList struct {
	head *DNode
	tail *DNode
	size int
}

func NewDoublyList() *DoublyList {
	return &DoublyList{}
}

func (dl *DoublyList) Clear() {
	dl.head = nil
	dl.tail = nil
	dl.size = 0
}

func (dl *DoublyList) PushFront(val string) {
	newNode := &DNode{data: val}
	if dl.head == nil {
		dl.head = newNode
		dl.tail = newNode
	} else {
		newNode.next = dl.head
		dl.head.prev = newNode
		dl.head = newNode
	}
	dl.size++
}

func (dl *DoublyList) PushBack(val string) {
	newNode := &DNode{data: val}
	if dl.tail == nil {
		dl.head = newNode
		dl.tail = newNode
	} else {
		dl.tail.next = newNode
		newNode.prev = dl.tail
		dl.tail = newNode
	}
	dl.size++
}

func (dl *DoublyList) InsertAfter(target, val string) {
	current := dl.head
	for current != nil {
		if current.data == target {
			newNode := &DNode{data: val}
			newNode.next = current.next
			newNode.prev = current

			if current.next != nil {
				current.next.prev = newNode
			} else {
				dl.tail = newNode
			}
			current.next = newNode
			dl.size++
			return
		}
		current = current.next
	}
}

func (dl *DoublyList) InsertBefore(target, val string) {
	current := dl.head
	for current != nil {
		if current.data == target {
			newNode := &DNode{data: val}
			newNode.prev = current.prev
			newNode.next = current

			if current.prev != nil {
				current.prev.next = newNode
			} else {
				dl.head = newNode
			}
			current.prev = newNode
			dl.size++
			return
		}
		current = current.next
	}
}

func (dl *DoublyList) PopFront() {
	if dl.head == nil {
		return
	}
	dl.head = dl.head.next
	if dl.head != nil {
		dl.head.prev = nil
	} else {
		dl.tail = nil
	}
	dl.size--
}

func (dl *DoublyList) PopBack() {
	if dl.tail == nil {
		return
	}
	dl.tail = dl.tail.prev
	if dl.tail != nil {
		dl.tail.next = nil
	} else {
		dl.head = nil
	}
	dl.size--
}

func (dl *DoublyList) RemoveByValue(val string) {
	current := dl.head
	for current != nil {
		if current.data == val {
			if current.prev != nil {
				current.prev.next = current.next
			} else {
				dl.head = current.next
			}

			if current.next != nil {
				current.next.prev = current.prev
			} else {
				dl.tail = current.prev
			}
			current = nil
			dl.size--
			return
		}
		current = current.next
	}
}

func (dl *DoublyList) Search(val string) bool {
	current := dl.head
	for current != nil {
		if current.data == val {
			return true
		}
		current = current.next
	}
	return false
}

func (dl *DoublyList) GetTail() string {
	if dl.tail != nil {
		return dl.tail.data
	}
	return ""
}

func (dl *DoublyList) GetSize() int {
	return dl.size
}

func (dl *DoublyList) PrintForward() {
	current := dl.head
	for current != nil {
		fmt.Printf("%s <-> ", current.data)
		current = current.next
	}
	fmt.Println("NULL")
}

func (dl *DoublyList) PrintBackward() {
	current := dl.tail
	for current != nil {
		fmt.Printf("%s <-> ", current.data)
		current = current.prev
	}
	fmt.Println("NULL")
}

func (dl *DoublyList) SaveToText(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	current := dl.head
	for current != nil {
		fmt.Fprintln(file, current.data)
		current = current.next
	}
	return nil
}

func (dl *DoublyList) LoadFromText(filename string) error {
	dl.Clear()

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var line string
	for {
		_, err := fmt.Fscanln(file, &line)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		dl.PushBack(line)
	}
	return nil
}

func (dl *DoublyList) SaveToBinary(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = binary.Write(file, binary.LittleEndian, int32(dl.size))
	if err != nil {
		return err
	}

	current := dl.head
	for current != nil {
		strBytes := []byte(current.data)
		err = binary.Write(file, binary.LittleEndian, int32(len(strBytes)))
		if err != nil {
			return err
		}
		_, err = file.Write(strBytes)
		if err != nil {
			return err
		}
		current = current.next
	}
	return nil
}

func (dl *DoublyList) LoadFromBinary(filename string) error {
	dl.Clear()

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

	for i := 0; i < int(size); i++ {
		var strLen int32
		err = binary.Read(file, binary.LittleEndian, &strLen)
		if err != nil {
			return err
		}

		strBytes := make([]byte, strLen)
		_, err = io.ReadFull(file, strBytes)
		if err != nil {
			return err
		}

		dl.PushBack(string(strBytes))
	}
	return nil
}
