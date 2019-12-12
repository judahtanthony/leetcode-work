#include <vector>
#include <unordered_map>

using namespace std;


// Definition for a Node.
class Node {
public:
    int val;
    vector<Node*> neighbors;

    Node() {}

    Node(int _val, vector<Node*> _neighbors) {
        val = _val;
        neighbors = _neighbors;
    }
};

class Solution {
public:
    Node* cloneGraph(Node* node) {
        std::unordered_map<Node*, Node*> visited;
        return this->_clone(node, &visited);
    }
    
    Node* _clone(Node* node, std::unordered_map<Node*, Node*>* visited) {
        if (node == NULL) {
            return NULL;
        }
        auto cached = visited->find(node);
        if (cached != visited->end()) {
            return cached->second;
        }

        Node* cloned = new Node();
        cloned->val = node->val;
        visited->insert({node, cloned});
        
        for (Node* n : node->neighbors) {
            Node* nn = this->_clone(n, visited);
            if (nn != NULL) {
                cloned->neighbors.push_back(nn);
            }
        }
            
        return cloned;
    }
};
