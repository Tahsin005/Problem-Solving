class Solution {
public:
    bool searchMatrix(vector<vector<int>>& matrix, int target) {
        if (matrix.empty() || matrix[0].empty()) return false;
        
        int rows = matrix.size();
        int cols = matrix[0].size();
        
        int l = 0;
        int r = (rows * cols) - 1;
        
        while (l <= r) {
            int mid = l + (r - l) / 2;
            
            int row = mid / cols;
            int col = mid % cols;
            int val = matrix[row][col];
            
            if (val == target) {
                return true;
            } else if (val < target) {
                l = mid + 1;
            } else {
                r = mid - 1;
            }
        }
        
        return false;
    }
};