#pragma once

#include <fstream>
#include <iostream>
#include <list>
#include <string>
#include <vector>

using namespace std;

struct HashNode {
    string key;
    string value;
};

class HashTable {
   private:
    vector<list<HashNode>> table;
    int capacity;
    int size;

    size_t hashFunction(const string& key) const {
        size_t hash = 0;
        for (char c : key) {
            hash = (hash * 31 + c) % capacity;
        }
        return hash;
    }

   public:
    HashTable(int cap = 10) : capacity(cap), size(0) { table.resize(capacity); }

    void put(const string& key, const string& value) {
        size_t index = hashFunction(key);
        for (auto& node : table[index]) {
            if (node.key == key) {
                node.value = value;
                return;
            }
        }
        table[index].push_back({key, value});
        size++;
    }

    string get(const string& key) const {
        size_t index = hashFunction(key);
        for (const auto& node : table[index]) {
            if (node.key == key) {
                return node.value;
            }
        }
        return "";
    }

    void remove(const string& key) {
        size_t index = hashFunction(key);
        auto& chain = table[index];
        for (auto it = chain.begin(); it != chain.end(); ++it) {
            if (it->key == key) {
                chain.erase(it);
                size--;
                return;
            }
        }
    }

    int getSize() const { return size; }

    void saveToText(const string& filename) {
        ofstream file(filename);
        file << size << endl;
        for (const auto& chain : table) {
            for (const auto& node : chain) {
                file << node.key << " " << node.value << endl;
            }
        }
        file.close();
    }

    void loadFromText(const string& filename) {
        ifstream file(filename);
        if (!file.is_open()) return;
        int newSize;
        file >> newSize;

        for (auto& chain : table) chain.clear();
        size = 0;

        string k, v;
        for (int i = 0; i < newSize; i++) {
            file >> k >> v;
            put(k, v);
        }
        file.close();
    }
};