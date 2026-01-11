package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

type SNode struct {
	data string
	next *SNode
}

type SinglyList struct {
	head *SNode
	tail *SNode
	size int
}

func NewSinglyList() *SinglyList {
	return &SinglyList{}
}

func (sl *SinglyList) Clear() {
	sl.head = nil
	sl.tail = nil
	sl.size = 0
}

func (sl *SinglyList) PushFront(val string) {
	newNode := &SNode{data: val}
	newNode.next = sl.head
	sl.head = newNode
	if sl.tail == nil {
		sl.tail = sl.head
	}
	sl.size++
}

func (sl *SinglyList) PushBack(val string) {
	newNode := &SNode{data: val}
	if sl.head == nil {
		sl.head = newNode
		sl.tail = newNode
	} else {
		sl.tail.next = newNode
		sl.tail = newNode
	}
	sl.size++
}

func (sl *SinglyList) InsertAfter(target, val string) {
	current := sl.head
	for current != nil {
		if current.data == target {
			newNode := &SNode{data: val}
			newNode.next = current.next
			current.next = newNode
			if current == sl.tail {
				sl.tail = newNode
			}
			sl.size++
			return
		}
		current = current.next
	}
}

func (sl *SinglyList) InsertBefore(target, val string) {
	if sl.head == nil {
		return
	}

	if sl.head.data == target {
		sl.PushFront(val)
		return
	}

	current := sl.head
	for current.next != nil {
		if current.next.data == target {
			newNode := &SNode{data: val}
			newNode.next = current.next
			current.next = newNode
			sl.size++
			return
		}
		current = current.next
	}
}

func (sl *SinglyList) PopFront() {
	if sl.head == nil {
		return
	}
	sl.head = sl.head.next
	if sl.head == nil {
		sl.tail = nil
	}
	sl.size--
}

func (sl *SinglyList) PopBack() {
	if sl.head == nil {
		return
	}
	if sl.head == sl.tail {
		sl.head = nil
		sl.tail = nil
	} else {
		current := sl.head
		for current.next != sl.tail {
			current = current.next
		}
		current.next = nil
		sl.tail = current
	}
	sl.size--
}

func (sl *SinglyList) RemoveByValue(val string) {
	if sl.head == nil {
		return
	}

	if sl.head.data == val {
		sl.PopFront()
		return
	}

	current := sl.head
	for current.next != nil {
		if current.next.data == val {
			temp := current.next
			current.next = temp.next
			if temp == sl.tail {
				sl.tail = current
			}
			sl.size--
			return
		}
		current = current.next
	}
}

func (sl *SinglyList) Search(val string) bool {
	current := sl.head
	for current != nil {
		if current.data == val {
			return true
		}
		current = current.next
	}
	return false
}

func (sl *SinglyList) GetHead() string {
	if sl.head != nil {
		return sl.head.data
	}
	return ""
}

func (sl *SinglyList) GetSize() int {
	return sl.size
}

func (sl *SinglyList) Print() {
	current := sl.head
	for current != nil {
		fmt.Printf("%s -> ", current.data)
		current = current.next
	}
	fmt.Println("NULL")
}

func (sl *SinglyList) SaveToText(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	current := sl.head
	for current != nil {
		fmt.Fprintln(file, current.data)
		current = current.next
	}
	return nil
}

func (sl *SinglyList) LoadFromText(filename string) error {
	sl.Clear()

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
		sl.PushBack(line)
	}
	return nil
}

func (sl *SinglyList) SaveToBinary(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = binary.Write(file, binary.LittleEndian, int32(sl.size))
	if err != nil {
		return err
	}

	current := sl.head
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

func (sl *SinglyList) LoadFromBinary(filename string) error {
	sl.Clear()

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

		sl.PushBack(string(strBytes))
	}
	return nil
}
