#include "../hashTable.h"

#include <gtest/gtest.h>
#include <cstdio>

TEST(HashTableTest, PutGet) {
    HashTable table;
    table.put("key1", "value1");
    table.put("key2", "value2");
    ASSERT_EQ(table.get("key1"), "value1");
    ASSERT_EQ(table.get("key2"), "value2");
}

TEST(HashTableTest, Remove) {
    HashTable table;
    table.put("test", "data");
    table.remove("test");
    ASSERT_EQ(table.get("test"), "");
    ASSERT_EQ(table.getSize(), 0);
}

TEST(HashTableTest, EmptyHash) {
    HashTable table;
    ASSERT_EQ(table.get("missing"), "");
    ASSERT_EQ(table.getSize(), 0);
}

TEST(HashTableTest, UpdateValue) {
    HashTable table;
    table.put("key", "old");
    table.put("key", "new");
    ASSERT_EQ(table.get("key"), "new");
    ASSERT_EQ(table.getSize(), 1);
}

TEST(HashTableTest, MultipleCollisions) {
    HashTable table(2);
    table.put("aa", "1");
    table.put("bb", "2");
    ASSERT_EQ(table.getSize(), 2);
}

TEST(HashTableTest, RemoveNonExisting) {
    HashTable table;
    table.put("key1", "val1");
    table.remove("key2"); // Не должно падать
    ASSERT_EQ(table.getSize(), 1);
    ASSERT_EQ(table.get("key1"), "val1");
}

TEST(HashTableTest, SaveLoadText) {
    HashTable table;
    table.put("k1", "v1");
    table.put("k2", "v2");
    table.saveToText("hash.txt");
    
    HashTable table2;
    table2.loadFromText("hash.txt");
    ASSERT_EQ(table2.getSize(), 2);
    ASSERT_EQ(table2.get("k1"), "v1");
    ASSERT_EQ(table2.get("k2"), "v2");
    
    remove("hash.txt");
}

TEST(HashTableTest, LoadFromNonExistingFile) {
    HashTable table;
    table.loadFromText("non_existing.txt");
    ASSERT_EQ(table.getSize(), 0);
}

TEST(HashTableTest, ChainOperations) {
    HashTable table(1); // Все будет в одной цепи
    table.put("a", "1");
    table.put("b", "2");
    table.put("c", "3");
    
    ASSERT_EQ(table.getSize(), 3);
    ASSERT_EQ(table.get("a"), "1");
    ASSERT_EQ(table.get("b"), "2");
    ASSERT_EQ(table.get("c"), "3");
    
    table.remove("b");
    ASSERT_EQ(table.getSize(), 2);
    ASSERT_EQ(table.get("b"), "");
}

TEST(HashTableTest, ClearAndReuse) {
    HashTable table;
    table.put("x", "y");
    table.remove("x");
    
    table.put("new", "value");
    ASSERT_EQ(table.getSize(), 1);
    ASSERT_EQ(table.get("new"), "value");
}