class Solution {
public:
    int countSuffix(vector<int>& nums, vector<int>& target) {
        int n = nums.size();
        int i;
        
        for (i = 0; i < n; i++) {
            if (nums[i] != target[i]) break;
        }
        
        sort(nums.begin() + i, nums.end(), greater<int>());
        
        if (nums == target) {
            int ans = 1;
            
            for (int j = min(i - 1, n - 2); j >= 0; j--) {
                if (nums[j] >= nums[j + 1]) {
                    ans++;
                } else {
                    break;
                }
            }
            return ans;
        }
        
        return 0;
    }
};
