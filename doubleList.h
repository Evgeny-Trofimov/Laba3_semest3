#pragma once

#include <fstream>
#include <iostream>
#include <string>

using namespace std;

struct DNode {
    string data;
    DNode* next;
    DNode* prev;
    DNode(string val) : data(val), next(nullptr), prev(nullptr) {}
};

class DoublyList {
   public:
    DoublyList() : head(nullptr), tail(nullptr), size(0) {}

    ~DoublyList() { clear(); }

    void clear() {
        while (head) {
            DNode* temp = head;
            head = head->next;
            delete temp;
        }
        tail = nullptr;
        size = 0;
    }

    void pushFront(const string& val) {
        DNode* newNode = new DNode(val);
        if (!head) {
            head = tail = newNode;
        } else {
            newNode->next = head;
            head->prev = newNode;
            head = newNode;
        }
        size++;
    }

    void pushBack(const string& val) {
        DNode* newNode = new DNode(val);
        if (!tail) {
            head = tail = newNode;
        } else {
            tail->next = newNode;
            newNode->prev = tail;
            tail = newNode;
        }
        size++;
    }

    void insertAfter(const string& target, const string& val) {
        DNode* current = head;
        while (current) {
            if (current->data == target) {
                DNode* newNode = new DNode(val);
                newNode->next = current->next;
                newNode->prev = current;

                if (current->next) {
                    current->next->prev = newNode;
                } else {
                    tail = newNode;
                }
                current->next = newNode;
                size++;
                return;
            }
            current = current->next;
        }
    }

    void insertBefore(const string& target, const string& val) {
        DNode* current = head;
        while (current) {
            if (current->data == target) {
                DNode* newNode = new DNode(val);
                newNode->prev = current->prev;
                newNode->next = current;

                if (current->prev) {
                    current->prev->next = newNode;
                } else {
                    head = newNode;
                }
                current->prev = newNode;
                size++;
                return;
            }
            current = current->next;
        }
    }

    void popFront() {
        if (!head) return;
        DNode* temp = head;
        head = head->next;
        if (head)
            head->prev = nullptr;
        else
            tail = nullptr;
        delete temp;
        size--;
    }

    void popBack() {
        if (!tail) return;
        DNode* temp = tail;
        tail = tail->prev;
        if (tail)
            tail->next = nullptr;
        else
            head = nullptr;
        delete temp;
        size--;
    }

    void removeByValue(const string& val) {
        DNode* current = head;
        while (current) {
            if (current->data == val) {
                if (current->prev)
                    current->prev->next = current->next;
                else
                    head = current->next;

                if (current->next)
                    current->next->prev = current->prev;
                else
                    tail = current->prev;

                delete current;
                size--;
                return;
            }
            current = current->next;
        }
    }

    bool search(const string& val) const {
        DNode* current = head;
        while (current) {
            if (current->data == val) return true;
            current = current->next;
        }
        return false;
    }

    string getTail() const { return tail ? tail->data : ""; }
    int getSize() const { return size; }

    void printForward() const {
        DNode* curr = head;
        while (curr) {
            cout << curr->data << " <-> ";
            curr = curr->next;
        }
        cout << "NULL" << endl;
    }

    void printBackward() const {
        DNode* curr = tail;
        while (curr) {
            cout << curr->data << " <-> ";
            curr = curr->prev;
        }
        cout << "NULL" << endl;
    }

    void saveToText(const string& filename) {
        ofstream file(filename);
        DNode* curr = head;
        while (curr) {
            file << curr->data << endl;
            curr = curr->next;
        }
        file.close();
    }

    void loadFromText(const string& filename) {
        clear();
        ifstream file(filename);
        if (!file.is_open()) return;
        string line;
        while (getline(file, line)) {
            pushBack(line);
        }
        file.close();
    }

    void saveToBinary(const string& filename) {
        ofstream file(filename, ios::binary);
        file.write(reinterpret_cast<char*>(&size), sizeof(size));
        DNode* curr = head;
        while (curr) {
            size_t len = curr->data.size();
            file.write(reinterpret_cast<char*>(&len), sizeof(len));
            file.write(curr->data.c_str(), len);
            curr = curr->next;
        }
        file.close();
    }

    void loadFromBinary(const string& filename) {
        clear();
        ifstream file(filename, ios::binary);
        if (!file.is_open()) return;
        int s;
        file.read(reinterpret_cast<char*>(&s), sizeof(s));
        for (int i = 0; i < s; ++i) {
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
    DNode* head;
    DNode* tail;
    int size;
};