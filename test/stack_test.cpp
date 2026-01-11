#include "../stack.h"

#include <gtest/gtest.h>
#include <cstdio>

TEST(StackTest, EmptyStack) {
    Stack s;
    EXPECT_EQ(s.getSize(), 0);
    EXPECT_EQ(s.pop(), "");
    EXPECT_EQ(s.peek(), "");
}

TEST(StackTest, ResizeInternalBuffer) {
    Stack s(1);
    s.push("a");
    s.push("b");
    s.push("c");
    EXPECT_EQ(s.getSize(), 3);
    EXPECT_EQ(s.pop(), "c");
    EXPECT_EQ(s.pop(), "b");
    EXPECT_EQ(s.pop(), "a");
}

TEST(StackTest, SaveLoadText) {
    Stack s;
    s.push("one");
    s.push("two");
    s.saveToText("stack_txt.txt");

    Stack s2;
    s2.loadFromText("stack_txt.txt");
    EXPECT_EQ(s2.getSize(), 2);
    EXPECT_EQ(s2.pop(), "two");
    EXPECT_EQ(s2.pop(), "one");
    
    remove("stack_txt.txt");
}

TEST(StackTest, SaveLoadBinary) {
    Stack s;
    s.push("x");
    s.push("y");
    s.saveToBinary("stack_bin.dat");

    Stack s2;
    s2.loadFromBinary("stack_bin.dat");
    EXPECT_EQ(s2.getSize(), 2);
    EXPECT_EQ(s2.pop(), "y");
    EXPECT_EQ(s2.pop(), "x");
    
    remove("stack_bin.dat");
}

TEST(StackTest, LoadFromNonExistingFiles) {
    Stack s;
    s.loadFromText("no_such_file.txt");
    s.loadFromBinary("no_such_file.bin");
    EXPECT_EQ(s.getSize(), 0);
}

TEST(StackTest, ZeroInitialCapacity) {
    Stack s(0);
    s.push("a");
    EXPECT_EQ(s.getSize(), 1);
}

TEST(StackTest, PushWithoutResize) {
    Stack s(10);
    s.push("a");
    s.push("b");
    EXPECT_EQ(s.getSize(), 2);
    EXPECT_EQ(s.peek(), "b");
}

TEST(StackTest, MultiplePushPop) {
    Stack s;
    for (int i = 0; i < 100; i++) {
        s.push(to_string(i));
    }
    EXPECT_EQ(s.getSize(), 100);
    
    for (int i = 99; i >= 0; i--) {
        EXPECT_EQ(s.pop(), to_string(i));
    }
    EXPECT_EQ(s.getSize(), 0);
}

TEST(StackTest, PeekOnNonEmpty) {
    Stack s;
    s.push("first");
    s.push("second");
    EXPECT_EQ(s.peek(), "second");
    s.pop();
    EXPECT_EQ(s.peek(), "first");
}

TEST(StackTest, ConstructWithDifferentCapacities) {
    Stack s1(5);
    EXPECT_EQ(s1.getSize(), 0);
    s1.push("test");
    EXPECT_EQ(s1.pop(), "test");
    
    Stack s2(100);
    for (int i = 0; i < 50; i++) {
        s2.push(to_string(i));
    }
    EXPECT_EQ(s2.getSize(), 50);
}