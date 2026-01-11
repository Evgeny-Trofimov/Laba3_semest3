#pragma once

#include <fstream>
#include <iostream>
#include <string>

using namespace std;

class Queue {
   private:
    string* data;
    int front;
    int rear;
    int size;
    int capacity;

   public:
    Queue(int cap = 10) {
        capacity = cap;
        data = new string[capacity];
        front = 0;
        rear = -1;
        size = 0;
    }

    ~Queue() { delete[] data; }

    void resize() {
        int newCap = capacity * 2;
        string* newData = new string[newCap];
        for (int i = 0; i < size; i++) {
            newData[i] = data[(front + i) % capacity];
        }
        delete[] data;
        data = newData;
        capacity = newCap;
        front = 0;
        rear = size - 1;
    }

    void push(const string& val) {
        if (size == capacity) resize();
        rear = (rear + 1) % capacity;
        data[rear] = val;
        size++;
    }

    string pop() {
        if (size == 0) return "";
        string val = data[front];
        front = (front + 1) % capacity;
        size--;
        return val;
    }

    string peek() const {
        if (size == 0) return "";
        return data[front];
    }

    int getSize() const { return size; }

    void saveToText(const string& filename) {
        ofstream file(filename);
        file << size << endl;
        for (int i = 0; i < size; i++) {
            file << data[(front + i) % capacity] << endl;
        }
        file.close();
    }

    void loadFromText(const string& filename) {
        ifstream file(filename);
        if (!file.is_open()) return;
        int newSize;
        file >> newSize;
        string dummy;
        getline(file, dummy);

        delete[] data;
        capacity = (newSize > 0) ? newSize * 2 : 10;
        data = new string[capacity];
        front = 0;
        rear = -1;
        size = 0;

        for (int i = 0; i < newSize; i++) {
            string val;
            getline(file, val);
            push(val);
        }
        file.close();
    }

    void saveToBinary(const string& filename) {
        ofstream file(filename, ios::binary);
        file.write(reinterpret_cast<char*>(&size), sizeof(size));
        for (int i = 0; i < size; i++) {
            string val = data[(front + i) % capacity];
            size_t len = val.size();
            file.write(reinterpret_cast<char*>(&len), sizeof(len));
            file.write(val.c_str(), len);
        }
        file.close();
    }

    void loadFromBinary(const string& filename) {
        ifstream file(filename, ios::binary);
        if (!file.is_open()) return;
        int newSize;
        file.read(reinterpret_cast<char*>(&newSize), sizeof(newSize));

        delete[] data;
        capacity = (newSize > 0) ? newSize * 2 : 10;
        data = new string[capacity];
        front = 0;
        rear = -1;
        size = 0;

        for (int i = 0; i < newSize; ++i) {
            size_t len;
            file.read(reinterpret_cast<char*>(&len), sizeof(len));
            char* buf = new char[len + 1];
            file.read(buf, len);
            buf[len] = '\0';
            push(string(buf));
            delete[] buf;
        }
        file.close();
    }
};