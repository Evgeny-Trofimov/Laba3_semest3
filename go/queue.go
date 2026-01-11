package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Queue struct {
	data     []string
	front    int
	rear     int
	size     int
	capacity int
}

func NewQueue(cap int) *Queue {
	if cap <= 0 {
		cap = 10
	}
	return &Queue{
		data:     make([]string, cap),
		front:    0,
		rear:     -1,
		size:     0,
		capacity: cap,
	}
}

func (q *Queue) resize() {
	newCap := q.capacity * 2
	newData := make([]string, newCap)

	for i := 0; i < q.size; i++ {
		newData[i] = q.data[(q.front+i)%q.capacity]
	}

	q.data = newData
	q.capacity = newCap
	q.front = 0
	q.rear = q.size - 1
}

func (q *Queue) Push(val string) {
	if q.size == q.capacity {
		q.resize()
	}
	q.rear = (q.rear + 1) % q.capacity
	q.data[q.rear] = val
	q.size++
}

func (q *Queue) Pop() string {
	if q.size == 0 {
		return ""
	}
	val := q.data[q.front]
	q.front = (q.front + 1) % q.capacity
	q.size--
	return val
}

func (q *Queue) Peek() string {
	if q.size == 0 {
		return ""
	}
	return q.data[q.front]
}

func (q *Queue) GetSize() int {
	return q.size
}

func (q *Queue) SaveToText(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	fmt.Fprintf(writer, "%d\n", q.size)
	for i := 0; i < q.size; i++ {
		fmt.Fprintln(writer, q.data[(q.front+i)%q.capacity])
	}
	return writer.Flush()
}

func (q *Queue) LoadFromText(filename string) error {
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

	q.data = make([]string, newSize*2)
	q.capacity = newSize * 2
	q.front = 0
	q.rear = -1
	q.size = 0

	// Читаем остальные строки - данные
	for i := 0; i < newSize && scanner.Scan(); i++ {
		val := scanner.Text()
		q.Push(val)
	}

	return scanner.Err()
}

func (q *Queue) SaveToBinary(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = binary.Write(file, binary.LittleEndian, int32(q.size))
	if err != nil {
		return err
	}

	for i := 0; i < q.size; i++ {
		val := q.data[(q.front+i)%q.capacity]
		strBytes := []byte(val)

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

func (q *Queue) LoadFromBinary(filename string) error {
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

	q.data = make([]string, newSize*2)
	q.capacity = int(newSize) * 2
	q.front = 0
	q.rear = -1
	q.size = 0

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

		q.Push(string(strBytes))
	}
	return nil
}
