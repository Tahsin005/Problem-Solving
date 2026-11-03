class Solution {
public:
    int minMergedInteger(vector<int>& nums) {
        long long best = LLONG_MAX;
        for (int i = 0; i + 1 < (int)nums.size(); i += 2) {
            long long merged = stoll(to_string(nums[i]) + to_string(nums[i + 1]));
            best = min(best, merged);
        }
        return (int)best;
    }
};
