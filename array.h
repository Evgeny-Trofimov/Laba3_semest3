#pragma once

#include <fstream>
#include <iostream>
#include <stdexcept>
#include <string>

using namespace std;

class Array {
   public:
    Array(int initialCapacity = 10) {
        capacity = initialCapacity;
        size = 0;
        data = new string[capacity];
    }

    ~Array() { delete[] data; }

    void pushBack(const string& value) {
        resizeCapacity();
        data[size++] = value;
    }

    void pushFront(const string& value) {
        resizeCapacity();
        for (int i = size; i > 0; i--) {
            data[i] = data[i - 1];
        }
        data[0] = value;
        size++;
    }

    void insertAt(int index, const string& value) {
        if (index < 0 || index > size) throw out_of_range("Index out of range");
        resizeCapacity();
        for (int i = size; i > index; i--) {
            data[i] = data[i - 1];
        }
        data[index] = value;
        size++;
    }

    void popBack() {
        if (size > 0) size--;
    }

    void popFront() {
        if (size == 0) return;
        for (int i = 0; i < size - 1; i++) {
            data[i] = data[i + 1];
        }
        size--;
    }

    void removeAt(int index) {
        if (index < 0 || index >= size) throw out_of_range("Index out of range");
        for (int i = index; i < size - 1; i++) {
            data[i] = data[i + 1];
        }
        size--;
    }

    int find(const string& value) const {
        for (int i = 0; i < size; ++i) {
            if (data[i] == value) return i;
        }
        return -1;
    }

    string get(int index) const {
        if (index < 0 || index >= size) throw out_of_range("Index out of range");
        return data[index];
    }

    void set(int index, const string& value) {
        if (index < 0 || index >= size) throw out_of_range("Index out of range");
        data[index] = value;
    }

    int getSize() const { return size; }

    void print() const {
        cout << "[ ";
        for (int i = 0; i < size; i++) {
            cout << data[i] << (i < size - 1 ? ", " : "");
        }
        cout << " ]" << endl;
    }

    void saveToText(const string& filename) {
        ofstream file(filename);
        file << size << endl;
        for (int i = 0; i < size; i++) {
            file << data[i] << endl;
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
        size = 0;

        for (int i = 0; i < newSize; i++) {
            string val;
            getline(file, val);
            pushBack(val);
        }
        file.close();
    }

    void saveToBinary(const string& filename) {
        ofstream file(filename, ios::binary);
        file.write(reinterpret_cast<char*>(&size), sizeof(size));
        for (int i = 0; i < size; ++i) {
            size_t len = data[i].size();
            file.write(reinterpret_cast<char*>(&len), sizeof(len));
            file.write(data[i].c_str(), len);
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
        size = 0;
        for (int i = 0; i < newSize; ++i) {
            size_t len;
            file.read(reinterpret_cast<char*>(&len), sizeof(len));
            char* buf = new char[len + 1];
            file.read(buf, len);
            buf[len] = '\0';
            pushBack(string(buf));
            delete[] buf;
        }
        file.close();
    }

   private:
    string* data;
    int size;
    int capacity;

    void resizeCapacity() {
        if (size >= capacity) {
            int newCapacity = (capacity == 0) ? 1 : capacity * 2;
            string* newData = new string[newCapacity];
            for (int i = 0; i < size; i++) {
                newData[i] = data[i];
            }
            delete[] data;
            data = newData;
            capacity = newCapacity;
        }
    }
};