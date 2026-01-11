package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Array struct {
	data []string
	size int
}

func NewArray(initialCapacity int) *Array {
	if initialCapacity <= 0 {
		initialCapacity = 10
	}
	return &Array{
		data: make([]string, initialCapacity),
		size: 0,
	}
}

func (a *Array) PushBack(value string) {
	a.ensureCapacity()
	a.data[a.size] = value
	a.size++
}

func (a *Array) PushFront(value string) {
	a.ensureCapacity()
	for i := a.size; i > 0; i-- {
		a.data[i] = a.data[i-1]
	}
	a.data[0] = value
	a.size++
}

func (a *Array) InsertAt(index int, value string) error {
	if index < 0 || index > a.size {
		return fmt.Errorf("index out of range")
	}
	a.ensureCapacity()
	for i := a.size; i > index; i-- {
		a.data[i] = a.data[i-1]
	}
	a.data[index] = value
	a.size++
	return nil
}

func (a *Array) PopBack() {
	if a.size > 0 {
		a.size--
	}
}

func (a *Array) PopFront() {
	if a.size == 0 {
		return
	}
	for i := 0; i < a.size-1; i++ {
		a.data[i] = a.data[i+1]
	}
	a.size--
}

func (a *Array) RemoveAt(index int) error {
	if index < 0 || index >= a.size {
		return fmt.Errorf("index out of range")
	}
	for i := index; i < a.size-1; i++ {
		a.data[i] = a.data[i+1]
	}
	a.size--
	return nil
}

func (a *Array) Find(value string) int {
	for i := 0; i < a.size; i++ {
		if a.data[i] == value {
			return i
		}
	}
	return -1
}

func (a *Array) Get(index int) (string, error) {
	if index < 0 || index >= a.size {
		return "", fmt.Errorf("index out of range")
	}
	return a.data[index], nil
}

func (a *Array) Set(index int, value string) error {
	if index < 0 || index >= a.size {
		return fmt.Errorf("index out of range")
	}
	a.data[index] = value
	return nil
}

func (a *Array) GetSize() int {
	return a.size
}

func (a *Array) Print() {
	fmt.Print("[ ")
	for i := 0; i < a.size; i++ {
		fmt.Print(a.data[i])
		if i < a.size-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println(" ]")
}

func (a *Array) ensureCapacity() {
	if a.size >= len(a.data) {
		newCapacity := len(a.data) * 2
		if newCapacity == 0 {
			newCapacity = 1
		}
		newData := make([]string, newCapacity)
		copy(newData, a.data[:a.size])
		a.data = newData
	}
}

func (a *Array) SaveToText(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	fmt.Fprintf(writer, "%d\n", a.size)
	for i := 0; i < a.size; i++ {
		fmt.Fprintln(writer, a.data[i])
	}
	return writer.Flush()
}

func (a *Array) LoadFromText(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Читаем первую строку - размер
	if !scanner.Scan() {
		return scanner.Err()
	}

	newSize, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return err
	}

	a.data = make([]string, newSize*2)
	a.size = 0

	// Читаем остальные строки - данные
	for i := 0; i < newSize && scanner.Scan(); i++ {
		val := scanner.Text()
		a.PushBack(val)
	}

	return scanner.Err()
}

func (a *Array) SaveToBinary(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = binary.Write(file, binary.LittleEndian, int32(a.size))
	if err != nil {
		return err
	}

	for i := 0; i < a.size; i++ {
		strBytes := []byte(a.data[i])
		err = binary.Write(file, binary.LittleEndian, int32(len(strBytes)))
		if err != nil {
			return err
		}
		_, err = file.Write(strBytes)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *Array) LoadFromBinary(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var newSize int32
	err = binary.Read(file, binary.LittleEndian, &newSize)
	if err != nil {
		return err
	}

	a.data = make([]string, newSize*2)
	a.size = 0

	for i := 0; i < int(newSize); i++ {
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

		a.PushBack(string(strBytes))
	}
	return nil
}
