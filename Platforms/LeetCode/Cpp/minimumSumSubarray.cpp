class Solution {
public:
    int minimumSumSubarray(vector<int>& nums, int l, int r) {
        int n = nums.size(), best = INT_MAX;
        for (int k = l; k <= r; ++k) {
            int s = accumulate(nums.begin(), nums.begin() + k - 1, 0);
            for (int i = k - 1; i < n; ++i) {
                s += nums[i];  // the window nums[i - k + 1 .. i]
                if (0 < s && s < best)
                    best = s;
                s -= nums[i - k + 1];
            }
        }
        return best == INT_MAX ? -1 : best;
    }
};
