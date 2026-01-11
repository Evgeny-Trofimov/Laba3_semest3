#include "../queue.h"

#include <gtest/gtest.h>
#include <cstdio>

TEST(QueueTest, PushPop) {
    Queue q;
    q.push("1");
    q.push("2");
    q.push("3");
    ASSERT_EQ(q.pop(), "1");
    ASSERT_EQ(q.pop(), "2");
    ASSERT_EQ(q.getSize(), 1);
}

TEST(QueueTest, EmptyQueue) {
    Queue q;
    ASSERT_EQ(q.pop(), "");
    ASSERT_EQ(q.peek(), "");
    ASSERT_EQ(q.getSize(), 0);
}

TEST(QueueTest, MultipleOperations) {
    Queue q;
    q.push("a");
    q.push("b");
    q.pop();
    q.push("c");
    ASSERT_EQ(q.pop(), "b");
    ASSERT_EQ(q.peek(), "c");
}

TEST(QueueTest, Resize) {
    Queue q(1);
    q.push("1");
    q.push("2");
    ASSERT_EQ(q.getSize(), 2);
}

TEST(QueueTest, Peek) {
    Queue q;
    q.push("test");
    ASSERT_EQ(q.peek(), "test");
    ASSERT_EQ(q.getSize(), 1);
}

TEST(QueueTest, SaveLoadText) {
    Queue q;
    q.push("a");
    q.push("b");
    q.push("c");
    q.saveToText("queue.txt");
    
    Queue q2;
    q2.loadFromText("queue.txt");
    ASSERT_EQ(q2.getSize(), 3);
    ASSERT_EQ(q2.pop(), "a");
    ASSERT_EQ(q2.pop(), "b");
    ASSERT_EQ(q2.pop(), "c");
    
    remove("queue.txt");
}

TEST(QueueTest, SaveLoadBinary) {
    Queue q;
    q.push("x");
    q.push("y");
    q.saveToBinary("queue.bin");
    
    Queue q2;
    q2.loadFromBinary("queue.bin");
    ASSERT_EQ(q2.getSize(), 2);
    ASSERT_EQ(q2.pop(), "x");
    ASSERT_EQ(q2.pop(), "y");
    
    remove("queue.bin");
}

TEST(QueueTest, LoadFromNonExistingFiles) {
    Queue q;
    q.loadFromText("non_existing.txt");
    ASSERT_EQ(q.getSize(), 0);
    
    q.loadFromBinary("non_existing.bin");
    ASSERT_EQ(q.getSize(), 0);
}

TEST(QueueTest, CircularBufferWrap) {
    Queue q(3);
    q.push("1");
    q.push("2");
    q.push("3");
    q.pop(); // Освобождает место
    q.push("4"); // Должен перезаписать в начало
    
    ASSERT_EQ(q.getSize(), 3);
    ASSERT_EQ(q.pop(), "2");
    ASSERT_EQ(q.pop(), "3");
    ASSERT_EQ(q.pop(), "4");
}

TEST(QueueTest, MultipleResizeOperations) {
    Queue q(2);
    for (int i = 0; i < 10; i++) {
        q.push(to_string(i));
    }
    
    ASSERT_EQ(q.getSize(), 10);
    for (int i = 0; i < 10; i++) {
        ASSERT_EQ(q.pop(), to_string(i));
    }
}