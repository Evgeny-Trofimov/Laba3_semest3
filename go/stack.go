package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Stack struct {
	data     []string
	size     int
	capacity int
}

func NewStack(initialCapacity int) *Stack {
	if initialCapacity <= 0 {
		initialCapacity = 10
	}
	return &Stack{
		data:     make([]string, initialCapacity),
		size:     0,
		capacity: initialCapacity,
	}
}

func (s *Stack) resize() {
	newCapacity := s.capacity * 2
	if newCapacity == 0 {
		newCapacity = 1
	}
	newData := make([]string, newCapacity)
	copy(newData, s.data[:s.size])
	s.data = newData
	s.capacity = newCapacity
}

func (s *Stack) Push(value string) {
	if s.size >= s.capacity {
		s.resize()
	}
	s.data[s.size] = value
	s.size++
}

func (s *Stack) Pop() string {
	if s.size == 0 {
		return ""
	}
	val := s.data[s.size-1]
	s.size--
	return val
}

func (s *Stack) Peek() string {
	if s.size == 0 {
		return ""
	}
	return s.data[s.size-1]
}

func (s *Stack) GetSize() int {
	return s.size
}

func (s *Stack) SaveToText(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	fmt.Fprintf(writer, "%d\n", s.size)
	for i := 0; i < s.size; i++ {
		fmt.Fprintln(writer, s.data[i])
	}
	return writer.Flush()
}

func (s *Stack) LoadFromText(filename string) error {
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

	s.data = make([]string, newSize*2)
	s.capacity = newSize * 2
	s.size = 0

	// Читаем остальные строки - данные
	for i := 0; i < newSize && scanner.Scan(); i++ {
		val := scanner.Text()
		s.Push(val)
	}

	return scanner.Err()
}

func (s *Stack) SaveToBinary(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = binary.Write(file, binary.LittleEndian, int32(s.size))
	if err != nil {
		return err
	}

	for i := 0; i < s.size; i++ {
		strBytes := []byte(s.data[i])
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

func (s *Stack) LoadFromBinary(filename string) error {
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

	s.data = make([]string, newSize*2)
	s.capacity = int(newSize) * 2
	s.size = 0

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

		s.Push(string(strBytes))
	}
	return nil
}
