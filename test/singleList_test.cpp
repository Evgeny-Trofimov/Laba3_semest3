#include "../singleList.h"

#include <gtest/gtest.h>
#include <cstdio>

TEST(SinglyListTest, InsertBeforeAfterHeadTail) {
    SinglyList list;
    list.pushBack("B");
    list.insertBefore("B", "A");
    list.insertAfter("B", "C");
    EXPECT_EQ(list.getSize(), 3);
    EXPECT_TRUE(list.search("A"));
    EXPECT_TRUE(list.search("B"));
    EXPECT_TRUE(list.search("C"));
    EXPECT_EQ(list.getHead(), "A");
}

TEST(SinglyListTest, PopFrontBackOnSingleAndEmpty) {
    SinglyList list;
    list.popFront();
    list.popBack();
    list.pushBack("X");
    list.popBack();
    EXPECT_EQ(list.getSize(), 0);
    EXPECT_EQ(list.getHead(), "");
}

TEST(SinglyListTest, SaveLoadText) {
    SinglyList list;
    list.pushBack("1");
    list.pushBack("2");
    list.saveToText("slist.txt");

    SinglyList list2;
    list2.loadFromText("slist.txt");
    EXPECT_EQ(list2.getSize(), 2);
    EXPECT_TRUE(list2.search("1"));
    EXPECT_TRUE(list2.search("2"));
    
    remove("slist.txt");
}

TEST(SinglyListTest, SaveLoadBinary) {
    SinglyList list;
    list.pushBack("a");
    list.pushBack("b");
    list.saveToBinary("slist.bin");

    SinglyList list2;
    list2.loadFromBinary("slist.bin");
    EXPECT_EQ(list2.getSize(), 2);
    EXPECT_TRUE(list2.search("a"));
    EXPECT_TRUE(list2.search("b"));
    
    remove("slist.bin");
}

TEST(SinglyListTest, RemoveByValueNotFoundAndClear) {
    SinglyList list;
    list.pushBack("1");
    list.pushBack("2");

    list.removeByValue("X");
    EXPECT_EQ(list.getSize(), 2);
    EXPECT_TRUE(list.search("1"));
    EXPECT_TRUE(list.search("2"));

    list.clear();
    EXPECT_EQ(list.getSize(), 0);
    EXPECT_EQ(list.getHead(), "");
    EXPECT_FALSE(list.search("1"));
}

TEST(SinglyListTest, InsertBeforeMiddleAndRemoveMiddle) {
    SinglyList list;
    list.pushBack("A");
    list.pushBack("B");
    list.pushBack("C");

    list.insertBefore("B", "X");
    EXPECT_EQ(list.getSize(), 4);
    EXPECT_TRUE(list.search("X"));

    list.removeByValue("B");
    EXPECT_FALSE(list.search("B"));
    EXPECT_EQ(list.getSize(), 3);
}

TEST(SinglyListTest, PopFrontAndPopBackOnLongList) {
    SinglyList list;
    list.pushBack("1");
    list.pushBack("2");
    list.pushBack("3");

    list.popFront();
    EXPECT_FALSE(list.search("1"));
    EXPECT_EQ(list.getSize(), 2);

    list.popBack();
    EXPECT_FALSE(list.search("3"));
    EXPECT_EQ(list.getSize(), 1);
}

TEST(SinglyListTest, PushFrontOperations) {
    SinglyList list;
    list.pushFront("c");
    list.pushFront("b");
    list.pushFront("a");
    EXPECT_EQ(list.getSize(), 3);
    EXPECT_EQ(list.getHead(), "a");
}

TEST(SinglyListTest, InsertBeforeNonExisting) {
    SinglyList list;
    list.pushBack("a");
    list.insertBefore("b", "x"); // Не должно падать
    EXPECT_EQ(list.getSize(), 1);
}

TEST(SinglyListTest, InsertAfterNonExisting) {
    SinglyList list;
    list.pushBack("a");
    list.insertAfter("b", "x"); // Не должно падать
    EXPECT_EQ(list.getSize(), 1);
}

TEST(SinglyListTest, LoadFromNonExistingFiles) {
    SinglyList list;
    list.loadFromText("non_existing.txt");
    EXPECT_EQ(list.getSize(), 0);
    
    list.loadFromBinary("non_existing.bin");
    EXPECT_EQ(list.getSize(), 0);
}

TEST(SinglyListTest, RemoveHeadTail) {
    SinglyList list;
    list.pushBack("a");
    list.pushBack("b");
    list.pushBack("c");
    
    list.removeByValue("a"); // Удаление головы
    EXPECT_EQ(list.getHead(), "b");
    
    list.removeByValue("c"); // Удаление хвоста
    EXPECT_EQ(list.getSize(), 1);
    EXPECT_TRUE(list.search("b"));
}

TEST(SinglyListTest, PrintMethod) {
    SinglyList list;
    list.pushBack("test");
    // Просто проверяем что не падает
    list.print();
}