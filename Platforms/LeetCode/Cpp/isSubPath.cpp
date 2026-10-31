class Solution {
public:
    bool isSubPath(ListNode* head, TreeNode* root) {
        if (!root) return false;

        stack<TreeNode*> nodes;
        nodes.push(root);

        while (!nodes.empty()) {
            TreeNode* node = nodes.top();
            nodes.pop();

            if (isMatch(node, head)) {
                return true;
            }

            if (node->left) {
                nodes.push(node->left);
            }
            if (node->right) {
                nodes.push(node->right);
            }
        }

        return false;
    }

private:
    bool isMatch(TreeNode* node, ListNode* lst) {
        stack<pair<TreeNode*, ListNode*>> s;
        s.push({node, lst});

        while (!s.empty()) {
            auto [currentNode, currentList] = s.top();
            s.pop();

            while (currentNode && currentList) {
                if (currentNode->val != currentList->val) {
                    break;
                }
                currentList = currentList->next;

                if (currentList) {
                    if (currentNode->left) {
                        s.push({currentNode->left, currentList});
                    }
                    if (currentNode->right) {
                        s.push({currentNode->right, currentList});
                    }
                    break;
                }
            }

            if (!currentList) {
                return true;
            }
        }

        return false;
    }
};
