#include "../FBT.h"

#include <gtest/gtest.h>
#include <fstream>
#include <set>
#include <sstream>
#include <vector>
#include <cstdio>

using namespace std;

static vector<int> getElements(const string& str) {
    vector<int> result;
    istringstream iss(str);
    int value;
    while (iss >> value) {
        result.push_back(value);
    }
    return result;
}

TEST(FullBinaryTreeTest, InsertBasic) {
    FullBinaryTree tree;
    tree.TINSERT(10);
    ASSERT_TRUE(tree.ISMEMBER(10));
    ASSERT_EQ(tree.PRINT_INORDER(), "10");
}

TEST(FullBinaryTreeTest, InsertMultiple) {
    FullBinaryTree tree;
    tree.TINSERT(30);
    tree.TINSERT(10);
    tree.TINSERT(20);
    
    ASSERT_TRUE(tree.ISMEMBER(30));
    ASSERT_TRUE(tree.ISMEMBER(10));
    ASSERT_TRUE(tree.ISMEMBER(20));
    
    string inorder = tree.PRINT_INORDER();
    istringstream iss(inorder);
    int count = 0, value;
    while (iss >> value) count++;
    ASSERT_EQ(count, 3);
}

TEST(FullBinaryTreeTest, DeleteLeaf) {
    FullBinaryTree tree;
    tree.TINSERT(1);
    tree.TINSERT(2);
    tree.TINSERT(3);
    tree.TDEL(3);
    ASSERT_FALSE(tree.ISMEMBER(3));
}

TEST(FullBinaryTreeTest, DeleteRoot) {
    FullBinaryTree tree;
    tree.TINSERT(2);
    tree.TINSERT(1);
    tree.TINSERT(3);
    tree.TDEL(2);
    ASSERT_FALSE(tree.ISMEMBER(2));
    ASSERT_TRUE(tree.ISMEMBER(1));
    ASSERT_TRUE(tree.ISMEMBER(3));
}

TEST(FullBinaryTreeTest, AllTraversalsProduceSameElements) {
    FullBinaryTree tree;
    tree.TINSERT(3);
    tree.TINSERT(1);
    tree.TINSERT(2);
    
    auto inorder = getElements(tree.PRINT_INORDER());
    auto preorder = getElements(tree.PRINT_PREORDER());
    auto postorder = getElements(tree.PRINT_POSTORDER());
    auto bfs = getElements(tree.PRINT_BFS());
    
    set<int> inorderSet(inorder.begin(), inorder.end());
    set<int> preorderSet(preorder.begin(), preorder.end());
    set<int> postorderSet(postorder.begin(), postorder.end());
    set<int> bfsSet(bfs.begin(), bfs.end());
    
    ASSERT_EQ(inorderSet.size(), 3);
    ASSERT_EQ(inorderSet, preorderSet);
    ASSERT_EQ(inorderSet, postorderSet);
    ASSERT_EQ(inorderSet, bfsSet);
}

TEST(FullBinaryTreeTest, EmptyTree) {
    FullBinaryTree tree;
    ASSERT_FALSE(tree.ISMEMBER(999));
    ASSERT_EQ(tree.PRINT_INORDER(), "");
    ASSERT_EQ(tree.TGET(999), "");
}

TEST(FullBinaryTreeTest, TGETMethod) {
    FullBinaryTree tree;
    tree.TINSERT(42);
    tree.TINSERT(24);
    tree.TINSERT(100);
    
    ASSERT_EQ(tree.TGET(42), "42");
    ASSERT_EQ(tree.TGET(24), "24");
    ASSERT_EQ(tree.TGET(100), "100");
    ASSERT_EQ(tree.TGET(999), "");
}

TEST(FullBinaryTreeTest, SaveLoadBinary) {
    FullBinaryTree tree;
    tree.TINSERT(5);
    tree.TINSERT(3);
    tree.TINSERT(7);
    tree.TINSERT(2);
    tree.TINSERT(4);
    
    tree.saveToBinary("fulltree_test.bin");
    
    ifstream check("fulltree_test.bin", ios::binary);
    ASSERT_TRUE(check.is_open());
    check.close();
    
    FullBinaryTree tree2;
    tree2.loadFromBinary("fulltree_test.bin");
    
    ASSERT_TRUE(tree2.ISMEMBER(5));
    ASSERT_TRUE(tree2.ISMEMBER(3));
    ASSERT_TRUE(tree2.ISMEMBER(7));
    ASSERT_TRUE(tree2.ISMEMBER(2));
    ASSERT_TRUE(tree2.ISMEMBER(4));
    
    remove("fulltree_test.bin");
}

TEST(FullBinaryTreeTest, ClearTree) {
    FullBinaryTree tree;
    tree.TINSERT(1);
    tree.TINSERT(2);
    tree.TINSERT(3);
    
    ASSERT_TRUE(tree.ISMEMBER(1));
    ASSERT_EQ(tree.PRINT_INORDER().find("1") != string::npos, true);
    
    tree.clear();
    ASSERT_FALSE(tree.ISMEMBER(1));
    ASSERT_EQ(tree.PRINT_INORDER(), "");
}

TEST(FullBinaryTreeTest, InsertManyElements) {
    FullBinaryTree tree;
    
    for (int i = 1; i <= 15; i++) {
        tree.TINSERT(i);
    }
    
    for (int i = 1; i <= 15; i++) {
        ASSERT_TRUE(tree.ISMEMBER(i));
        ASSERT_EQ(tree.TGET(i), to_string(i));
    }
    
    string bfs = tree.PRINT_BFS();
    istringstream iss(bfs);
    int count = 0, value;
    while (iss >> value) count++;
    ASSERT_EQ(count, 15);
}

TEST(FullBinaryTreeTest, DeleteAndReinsert) {
    FullBinaryTree tree;
    tree.TINSERT(10);
    tree.TINSERT(20);
    tree.TINSERT(30);
    tree.TINSERT(40);
    
    tree.TDEL(20);
    ASSERT_FALSE(tree.ISMEMBER(20));
    ASSERT_TRUE(tree.ISMEMBER(10));
    ASSERT_TRUE(tree.ISMEMBER(30));
    ASSERT_TRUE(tree.ISMEMBER(40));
    
    tree.TINSERT(50);
    ASSERT_TRUE(tree.ISMEMBER(50));
    
    string inorder = tree.PRINT_INORDER();
    istringstream iss(inorder);
    int count = 0, value;
    while (iss >> value) count++;
    ASSERT_EQ(count, 4);
}

TEST(FullBinaryTreeTest, DeleteNonExisting) {
    FullBinaryTree tree;
    tree.TINSERT(1);
    tree.TINSERT(2);
    tree.TDEL(999); // Не должно падать
    ASSERT_TRUE(tree.ISMEMBER(1));
    ASSERT_TRUE(tree.ISMEMBER(2));
}

TEST(FullBinaryTreeTest, DeleteNodeWithTwoChildren) {
    FullBinaryTree tree;
    tree.TINSERT(10);
    tree.TINSERT(5);
    tree.TINSERT(15);
    tree.TINSERT(3);
    tree.TINSERT(7);
    tree.TINSERT(12);
    tree.TINSERT(20);
    
    tree.TDEL(5); // Узел с двумя детьми
    ASSERT_FALSE(tree.ISMEMBER(5));
    ASSERT_TRUE(tree.ISMEMBER(10));
    ASSERT_TRUE(tree.ISMEMBER(15));
    ASSERT_TRUE(tree.ISMEMBER(3));
    ASSERT_TRUE(tree.ISMEMBER(7));
    ASSERT_TRUE(tree.ISMEMBER(12));
    ASSERT_TRUE(tree.ISMEMBER(20));
}

TEST(FullBinaryTreeTest, LoadFromNonExistingFile) {
    FullBinaryTree tree;
    tree.loadFromBinary("non_existing.bin");
    ASSERT_FALSE(tree.ISMEMBER(1));
    ASSERT_EQ(tree.PRINT_INORDER(), "");
}

TEST(FullBinaryTreeTest, SingleNodeTree) {
    FullBinaryTree tree;
    tree.TINSERT(100);
    
    ASSERT_TRUE(tree.ISMEMBER(100));
    ASSERT_EQ(tree.TGET(100), "100");
    ASSERT_EQ(tree.PRINT_INORDER(), "100");
    
    tree.TDEL(100);
    ASSERT_FALSE(tree.ISMEMBER(100));
    ASSERT_EQ(tree.PRINT_INORDER(), "");
}

TEST(FullBinaryTreeTest, MultipleDeleteOperations) {
    FullBinaryTree tree;
    for (int i = 1; i <= 7; i++) {
        tree.TINSERT(i);
    }
    
    // Удаляем в произвольном порядке
    tree.TDEL(4);
    tree.TDEL(2);
    tree.TDEL(6);
    
    ASSERT_FALSE(tree.ISMEMBER(4));
    ASSERT_FALSE(tree.ISMEMBER(2));
    ASSERT_FALSE(tree.ISMEMBER(6));
    
    for (int i : {1, 3, 5, 7}) {
        ASSERT_TRUE(tree.ISMEMBER(i));
    }
}