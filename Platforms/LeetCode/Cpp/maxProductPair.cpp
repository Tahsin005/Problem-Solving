class Solution {
public:
    vector<int> maxProductPair(vector<int>& nums, int target) {
        int len = nums.size();
        vector<int> ans = {-1, -1};
        int best = INT_MIN;
        bool found = false;

        for (int i = 0; i < len; i++) {
            for (int j = 0; j < len; j++) {
                if (nums[i] > nums[j] && nums[i] + nums[j] == target) {
                    int prod = nums[i] * nums[j];
                    if (!found || prod > best) {
                        found = true;
                        best = prod;
                        ans = {i, j};
                    }
                }
            }
        }
        return ans;
    }
};
