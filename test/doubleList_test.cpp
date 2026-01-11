#include "../doubleList.h"

#include <gtest/gtest.h>
#include <cstdio>

TEST(DoubleListTest, InsertBeforeHeadAndAfterTail) {
    DoublyList list;
    list.pushBack("B");
    list.insertBefore("B", "A");
    list.insertAfter("B", "C");
    EXPECT_EQ(list.getSize(), 3);
    EXPECT_TRUE(list.search("A"));
    EXPECT_TRUE(list.search("B"));
    EXPECT_TRUE(list.search("C"));
}

TEST(DoubleListTest, SaveLoadBinary) {
    DoublyList list;
    list.pushBack("x");
    list.pushBack("y");
    list.saveToBinary("dlist.bin");

    DoublyList list2;
    list2.loadFromBinary("dlist.bin");
    EXPECT_EQ(list2.getSize(), 2);
    EXPECT_TRUE(list2.search("x"));
    EXPECT_TRUE(list2.search("y"));
    
    remove("dlist.bin");
}

TEST(DoubleListTest, PopFrontBackAndRemove) {
    DoublyList list;
    list.pushBack("1");
    list.pushBack("2");
    list.pushBack("3");

    list.popFront();
    EXPECT_FALSE(list.search("1"));
    EXPECT_EQ(list.getSize(), 2);

    list.popBack();
    EXPECT_FALSE(list.search("3"));
    EXPECT_EQ(list.getSize(), 1);

    list.removeByValue("X");
    EXPECT_EQ(list.getSize(), 1);

    list.removeByValue("2");
    EXPECT_EQ(list.getSize(), 0);
}

TEST(DoubleListTest, SaveLoadText) {
    DoublyList list;
    list.pushBack("a");
    list.pushBack("b");
    list.saveToText("dlist.txt");

    DoublyList list2;
    list2.loadFromText("dlist.txt");
    EXPECT_EQ(list2.getSize(), 2);
    EXPECT_TRUE(list2.search("a"));
    EXPECT_TRUE(list2.search("b"));
    
    remove("dlist.txt");
}

TEST(DoubleListTest, PushFrontOperations) {
    DoublyList list;
    list.pushFront("c");
    list.pushFront("b");
    list.pushFront("a");
    EXPECT_EQ(list.getSize(), 3);
    EXPECT_TRUE(list.search("a"));
}

TEST(DoubleListTest, InsertInMiddle) {
    DoublyList list;
    list.pushBack("a");
    list.pushBack("c");
    list.insertBefore("c", "b");
    list.insertAfter("b", "b2");
    EXPECT_EQ(list.getSize(), 4);
    EXPECT_TRUE(list.search("b"));
    EXPECT_TRUE(list.search("b2"));
}

TEST(DoubleListTest, PopOperationsOnEmpty) {
    DoublyList list;
    list.popFront(); // Не должно падать
    list.popBack();  // Не должно падать
    EXPECT_EQ(list.getSize(), 0);
}

TEST(DoubleListTest, RemoveNonExistingValue) {
    DoublyList list;
    list.pushBack("a");
    list.removeByValue("b"); // Не должно падать
    EXPECT_EQ(list.getSize(), 1);
}

TEST(DoubleListTest, LoadFromNonExistingFiles) {
    DoublyList list;
    list.loadFromText("non_existing.txt");
    EXPECT_EQ(list.getSize(), 0);
    
    list.loadFromBinary("non_existing.bin");
    EXPECT_EQ(list.getSize(), 0);
}

TEST(DoubleListTest, ClearList) {
    DoublyList list;
    list.pushBack("a");
    list.pushBack("b");
    list.clear();
    EXPECT_EQ(list.getSize(), 0);
    EXPECT_FALSE(list.search("a"));
}

TEST(DoubleListTest, GetTailOnEmpty) {
    DoublyList list;
    EXPECT_EQ(list.getTail(), "");
}

TEST(DoubleListTest, PrintMethods) {
    DoublyList list;
    list.pushBack("a");
    list.pushBack("b");
    // Просто проверяем что не падают
    list.printForward();
    list.printBackward();
}