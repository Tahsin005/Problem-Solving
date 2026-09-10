/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    int count = 0;
    
    pair<int, int> dfs(TreeNode* node) {
        if (!node) return {0, 0};
        
        auto [leftSum, leftCount] = dfs(node->left);
        auto [rightSum, rightCount] = dfs(node->right);
        
        int sum = leftSum + rightSum + node->val;
        int nodeCount = leftCount + rightCount + 1;
        
        if (sum / nodeCount == node->val) {
            count++;
        }
        
        return {sum, nodeCount};
    }
    
    int averageOfSubtree(TreeNode* root) {
        dfs(root);
        return count;
    }
};