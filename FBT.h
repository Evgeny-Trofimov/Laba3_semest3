#pragma once
#include <iostream>
#include <queue>
#include <sstream>
#include <string>
#include <vector>
#include <fstream>

using namespace std;

struct FBNode {
    int key;
    FBNode* left;
    FBNode* right;
    
    FBNode(int k) : key(k), left(nullptr), right(nullptr) {}
};

class FullBinaryTree {
public:
    FullBinaryTree() : root(nullptr) {}
    ~FullBinaryTree() { clear(); }

    void TINSERT(int key) { insert(key); }
    void TDEL(int key) { remove(key); }
    bool ISMEMBER(int key) const { return search(root, key); }
    string TGET(int key) const { 
        return search(root, key) ? to_string(key) : ""; 
    }
    
    string PRINT_PREORDER() const {
        vector<int> res;
        preorder(root, res);
        return vecToString(res);
    }
    
    string PRINT_INORDER() const {
        vector<int> res;
        inorder(root, res);
        return vecToString(res);
    }
    
    string PRINT_POSTORDER() const {
        vector<int> res;
        postorder(root, res);
        return vecToString(res);
    }
    
    string PRINT_BFS() const {
        vector<int> res;
        bfs(root, res);
        return vecToString(res);
    }
    
    void saveToBinary(const string& filename) const {
        ofstream file(filename, ios::binary);
        if (!file.is_open()) return;
        
        vector<int> keys;
        bfsForSerialization(root, keys);
        
        int size = keys.size();
        file.write(reinterpret_cast<const char*>(&size), sizeof(size));
        
        for (int key : keys) {
            file.write(reinterpret_cast<const char*>(&key), sizeof(key));
        }
        file.close();
    }
    
    void loadFromBinary(const string& filename) {
        clear();
        ifstream file(filename, ios::binary);
        if (!file.is_open()) return;
        
        int size;
        file.read(reinterpret_cast<char*>(&size), sizeof(size));
        
        if (size > 0) {
            vector<int> keys(size);
            for (int i = 0; i < size; ++i) {
                file.read(reinterpret_cast<char*>(&keys[i]), sizeof(int));
            }
            root = buildCompleteTree(keys, 0);
        }
        file.close();
    }
    
    void clear() {
        clearTree(root);
        root = nullptr;
    }
    
private:
    FBNode* root;
    
    bool isFull(FBNode* node) const {
        if (!node) return true;
        return (node->left && node->right) || (!node->left && !node->right);
    }
    
    void insert(int key) {
        if (!root) {
            root = new FBNode(key);
            return;
        }
        
        queue<FBNode*> q;
        q.push(root);
        
        while (!q.empty()) {
            FBNode* current = q.front();
            q.pop();
            
            if (!current->left) {
                current->left = new FBNode(key);
                return;
            }
            else if (!current->right) {
                current->right = new FBNode(key);
                return;
            }
            else {
                q.push(current->left);
                q.push(current->right);
            }
        }
    }
    
    bool search(FBNode* node, int key) const {
        if (!node) return false;
        if (node->key == key) return true;
        return search(node->left, key) || search(node->right, key);
    }
    
    void remove(int key) {
        if (!root) return;
        
        if (root->key == key && !root->left && !root->right) {
            delete root;
            root = nullptr;
            return;
        }
        
        FBNode* keyNode = nullptr;
        FBNode* deepest = nullptr;
        FBNode* parentOfDeepest = nullptr;
        
        queue<pair<FBNode*, FBNode*>> q;
        q.push({nullptr, root});
        
        while (!q.empty()) {
            auto [parent, current] = q.front();
            q.pop();
            
            if (current->key == key) {
                keyNode = current;
            }
            
            deepest = current;
            parentOfDeepest = parent;
            
            if (current->left) {
                q.push({current, current->left});
            }
            if (current->right) {
                q.push({current, current->right});
            }
        }
        
        if (!keyNode) return;
        
        keyNode->key = deepest->key;
        
        if (parentOfDeepest) {
            if (parentOfDeepest->left == deepest) {
                delete parentOfDeepest->left;
                parentOfDeepest->left = nullptr;
            } else {
                delete parentOfDeepest->right;
                parentOfDeepest->right = nullptr;
            }
        }
    }
    void preorder(FBNode* node, vector<int>& result) const {
        if (!node) return;
        result.push_back(node->key);
        preorder(node->left, result);
        preorder(node->right, result);
    }
    
    void inorder(FBNode* node, vector<int>& result) const {
        if (!node) return;
        inorder(node->left, result);
        result.push_back(node->key);
        inorder(node->right, result);
    }
    
    void postorder(FBNode* node, vector<int>& result) const {
        if (!node) return;
        postorder(node->left, result);
        postorder(node->right, result);
        result.push_back(node->key);
    }
    
    void bfs(FBNode* node, vector<int>& result) const {
        if (!node) return;
        queue<FBNode*> q;
        q.push(node);
        
        while (!q.empty()) {
            FBNode* current = q.front();
            q.pop();
            result.push_back(current->key);
            
            if (current->left) q.push(current->left);
            if (current->right) q.push(current->right);
        }
    }
    
    void bfsForSerialization(FBNode* node, vector<int>& result) const {
        if (!node) return;
        queue<FBNode*> q;
        q.push(node);
        
        while (!q.empty()) {
            FBNode* current = q.front();
            q.pop();
            result.push_back(current->key);
            
            if (current->left) q.push(current->left);
            if (current->right) q.push(current->right);
        }
    }
    
    FBNode* buildCompleteTree(const vector<int>& keys, int index) {
        if (index >= keys.size()) return nullptr;
        
        FBNode* node = new FBNode(keys[index]);
        node->left = buildCompleteTree(keys, 2 * index + 1);
        node->right = buildCompleteTree(keys, 2 * index + 2);
        
        return node;
    }
    
    void clearTree(FBNode* node) {
        if (!node) return;
        clearTree(node->left);
        clearTree(node->right);
        delete node;
    }
    
    string vecToString(const vector<int>& vec) const {
        if (vec.empty()) return "";
        ostringstream oss;
        for (size_t i = 0; i < vec.size(); ++i) {
            oss << vec[i];
            if (i + 1 < vec.size()) oss << " ";
        }
        return oss.str();
    }
};