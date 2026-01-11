#pragma once

#include <fstream>
#include <iostream>
#include <string>

using namespace std;

struct Node {
    string data;
    Node* next;
    Node(string val) : data(val), next(nullptr) {}
};

class SinglyList {
   private:
    Node* head;
    Node* tail;
    int size;

   public:
    SinglyList() : head(nullptr), tail(nullptr), size(0) {}

    ~SinglyList() { clear(); }

    void clear() {
        while (head) {
            Node* temp = head;
            head = head->next;
            delete temp;
        }
        tail = nullptr;
        size = 0;
    }

    void pushFront(const string& val) {
        Node* newNode = new Node(val);
        newNode->next = head;
        head = newNode;
        if (!tail) tail = head;
        size++;
    }

    void pushBack(const string& val) {
        Node* newNode = new Node(val);
        if (!head) {
            head = tail = newNode;
        } else {
            tail->next = newNode;
            tail = newNode;
        }
        size++;
    }

    void insertAfter(const string& target, const string& val) {
        Node* current = head;
        while (current) {
            if (current->data == target) {
                Node* newNode = new Node(val);
                newNode->next = current->next;
                current->next = newNode;
                if (current == tail) tail = newNode;
                size++;
                return;
            }
            current = current->next;
        }
    }

    void insertBefore(const string& target, const string& val) {
        if (!head) return;

        if (head->data == target) {
            pushFront(val);
            return;
        }

        Node* current = head;
        while (current->next) {
            if (current->next->data == target) {
                Node* newNode = new Node(val);
                newNode->next = current->next;
                current->next = newNode;
                size++;
                return;
            }
            current = current->next;
        }
    }

    void popFront() {
        if (!head) return;
        Node* temp = head;
        head = head->next;
        if (!head) tail = nullptr;
        delete temp;
        size--;
    }

    void popBack() {
        if (!head) return;
        if (head == tail) {
            delete head;
            head = tail = nullptr;
        } else {
            Node* current = head;
            while (current->next != tail) {
                current = current->next;
            }
            delete tail;
            tail = current;
            tail->next = nullptr;
        }
        size--;
    }

    void removeByValue(const string& val) {
        if (!head) return;

        if (head->data == val) {
            popFront();
            return;
        }

        Node* current = head;
        while (current->next) {
            if (current->next->data == val) {
                Node* temp = current->next;
                current->next = temp->next;
                if (temp == tail) tail = current;
                delete temp;
                size--;
                return;
            }
            current = current->next;
        }
    }

    bool search(const string& val) const {
        Node* current = head;
        while (current) {
            if (current->data == val) return true;
            current = current->next;
        }
        return false;
    }

    string getHead() const { return head ? head->data : ""; }
    int getSize() const { return size; }

    void print() const {
        Node* curr = head;
        while (curr) {
            cout << curr->data << " -> ";
            curr = curr->next;
        }
        cout << "NULL" << endl;
    }

    void saveToText(const string& filename) {
        ofstream file(filename);
        Node* curr = head;
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
        int s = size;
        file.write(reinterpret_cast<char*>(&s), sizeof(s));
        Node* curr = head;
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
};