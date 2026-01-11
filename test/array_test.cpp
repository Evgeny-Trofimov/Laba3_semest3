#include "../array.h"

#include <gtest/gtest.h>
#include <cstdio>

TEST(ArrayTest, Exceptions) {
    Array arr;
    EXPECT_THROW(arr.get(-1), std::out_of_range);
    EXPECT_THROW(arr.get(0), std::out_of_range);
    EXPECT_THROW(arr.removeAt(0), std::out_of_range);
    EXPECT_THROW(arr.insertAt(1, "x"), std::out_of_range);
}

TEST(ArrayTest, SaveLoadTextBinary) {
    Array arr;
    arr.pushBack("a");
    arr.pushBack("b");
    arr.saveToText("arr.txt");
    arr.saveToBinary("arr.bin");

    Array arr2;
    arr2.loadFromText("arr.txt");
    EXPECT_EQ(arr2.getSize(), 2);
    EXPECT_EQ(arr2.get(0), "a");

    Array arr3;
    arr3.loadFromBinary("arr.bin");
    EXPECT_EQ(arr3.getSize(), 2);
    EXPECT_EQ(arr3.get(1), "b");
    
    // Удаляем тестовые файлы
    remove("arr.txt");
    remove("arr.bin");
}

TEST(ArrayTest, BasicOperations) {
    Array arr;
    arr.pushBack("1");
    arr.pushBack("3");
    arr.pushFront("0");
    arr.insertAt(2, "2");

    EXPECT_EQ(arr.getSize(), 4);
    EXPECT_EQ(arr.get(0), "0");
    EXPECT_EQ(arr.get(1), "1");
    EXPECT_EQ(arr.get(2), "2");
    EXPECT_EQ(arr.get(3), "3");

    arr.popFront();
    arr.popBack();
    EXPECT_EQ(arr.getSize(), 2);
}

TEST(ArrayTest, FindAndSet) {
    Array arr;
    arr.pushBack("a");
    arr.pushBack("b");
    arr.pushBack("c");
    
    EXPECT_EQ(arr.find("b"), 1);
    EXPECT_EQ(arr.find("d"), -1);
    
    arr.set(1, "x");
    EXPECT_EQ(arr.get(1), "x");
}

TEST(ArrayTest, RemoveAtValid) {
    Array arr;
    arr.pushBack("a");
    arr.pushBack("b");
    arr.pushBack("c");
    
    arr.removeAt(1);
    EXPECT_EQ(arr.getSize(), 2);
    EXPECT_EQ(arr.get(0), "a");
    EXPECT_EQ(arr.get(1), "c");
}

TEST(ArrayTest, PopOperationsOnEmpty) {
    Array arr;
    arr.popBack(); // Не должно падать
    arr.popFront(); // Не должно падать
    EXPECT_EQ(arr.getSize(), 0);
}

TEST(ArrayTest, ResizeCapacity) {
    Array arr(2); // Начальная емкость 2
    arr.pushBack("1");
    arr.pushBack("2");
    // Должен произойти ресайз
    arr.pushBack("3");
    EXPECT_EQ(arr.getSize(), 3);
    EXPECT_EQ(arr.get(2), "3");
}

TEST(ArrayTest, LoadFromNonExistingFiles) {
    Array arr;
    arr.loadFromText("non_existing.txt");
    EXPECT_EQ(arr.getSize(), 0);
    
    arr.loadFromBinary("non_existing.bin");
    EXPECT_EQ(arr.getSize(), 0);
}

TEST(ArrayTest, InsertAtBoundaries) {
    Array arr;
    arr.pushBack("a");
    arr.insertAt(0, "first"); // В начало
    arr.insertAt(2, "last");  // В конец
    EXPECT_EQ(arr.getSize(), 3);
    EXPECT_EQ(arr.get(0), "first");
    EXPECT_EQ(arr.get(2), "last");
}

TEST(ArrayTest, EmptyArrayOperations) {
    Array arr;
    EXPECT_EQ(arr.getSize(), 0);
    EXPECT_THROW(arr.set(0, "x"), std::out_of_range);
    EXPECT_EQ(arr.find("x"), -1);
}

TEST(ArrayTest, PrintMethod) {
    Array arr;
    arr.pushBack("test1");
    arr.pushBack("test2");
    // Просто проверяем что не падает
    arr.print();
}